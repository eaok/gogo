[toc]

项目常见的代码组织形式

- controller  --> 处理请求的函数
- logic  --> 逻辑处理层
- models  --> 数据的定义及增删改查





# 日志

`zap`  --> `gin-zap`（Logger()/Recovery()）-->日志切割(`"github.com/natefinch/lumberjack`)

```go
"go.uber.org/zap"
```

配置部分见`logger/logger.go`

```go
// Logger 全局的日志对象
var Logger *zap.Logger

// InitLogger 初始化Logger
func InitLogger(cfg *config.LogConfig) (err error) {
	ws := getLogWriter(cfg.Filename, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge) // 做日志切割第三方包
	encoder := getEncoder()                                                   // 日志输出的格式
	var level = new(zapcore.Level)
	err = level.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return
	}
	//var writeSyncer zapcore.WriteSyncer
	//if *l == zapcore.DebugLevel {
	//	zapcore.AddSync(os.Stdout)
	//	writeSyncer = zapcore.NewMultiWriteSyncer(ws,os.Stdout)
	//}else {
	//	writeSyncer = ws
	//}
	core := zapcore.NewCore(encoder, ws, level)

	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 时间字符串
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder // 函数调用
	return zapcore.NewJSONEncoder(encoderConfig)            // JSON格式
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

//进行封装，方便外部调用
func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

func With(fields ...zap.Field) *zap.Logger {
	return Logger.With(fields...)
}
```

载入配置文件中的log配置：

```go
type LogConfig struct {
	Level      string `json:"level" ini:"level"`
	Filename   string `json:"filename" ini:"filename"`
	MaxSize    int    `json:"maxsize" ini:"maxsize"`
	MaxAge     int    `json:"max_age" ini:"max_age"`
	MaxBackups int    `json:"max_backups" ini:"max_backups"`
}

// 定义了全局的配置文件实例
var Conf = new(AppConfig)

func InitFromJson(file string) error {
	jsonData, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonData, Conf); err != nil {
		return err
	}

	return nil
}
```



使用：

```go
logger.Info("LoginPost", zap.Any("username", userName), zap.Any("password", password))
logger.Debug("UploadPost", zap.String("filename", fh.Filename))
```

可以对Logger()/Recovery()进行重写，让zap支持gin框架的日志，例如`middlewares/log.go`
```go
router.Use(middlewares.GinLogger(logger.Logger), middlewares.GinRecovery(logger.Logger, true))
```



# 配置

json格式的配置文件

```json
{
  "server": {
    "port": 8080
  },
}
```

解析json格式的配置文件：

```go
// AppConfig 应用的配置结构体
type AppConfig struct {
	*ServerConfig `json:"server"`
}

// ServerConfig web server配置
type ServerConfig struct {
	Port int `json:"port"`
}

// 定义了全局的配置文件实例
var Conf = new(AppConfig)

func InitFromJson() error {
	jsonData, err := ioutil.ReadFile("conf/conf.json")
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonData, Conf); err != nil {
		return err
	}

	return nil
}
```



ini格式的配置文件

```ini
[server]
port=8080

[redis]
host = "127.0.0.1"
port = 6379
db = 0
password = ""
```

解析ini格式的配置文件，用到了`https://github.com/go-ini/ini`：

```go
// AppConfig 应用的配置结构体
type AppConfig struct {
	*ServerConfig `json:"server" ini:"server"`
	*RedisConfig  `json:"redis" ini:"redis"`
}

// ServerConfig web server配置
type ServerConfig struct {
	Port int `json:"port" ini:"port"`
}

// RedisConfig redis配置
type RedisConfig struct {
	Host     string `json:"host" ini:"host"`
	Password string `json:"password" ini:"password"`
	Port     int    `json:"port" ini:"port"`
	DB       int    `json:"db" ini:"db"`
}

// 定义了全局的配置文件实例
var Conf = new(AppConfig)

func InitFromIni(filename string) error {
	err := ini.MapTo(Conf, filename)
	if err != nil {
		panic(err)
	}

	return err
}
```





# session部分

设置session中间件，这里用redis保存：

```go
import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// Session 设置session中间件
func Session() gin.HandlerFunc {
	address := fmt.Sprintf("%s:%d", config.Conf.RedisConfig.Host, 
                           config.Conf.RedisConfig.Port)
	store, err := redis.NewStore(10, "tcp", address, "", []byte("secret"))
	if err != nil {
		panic(err)
	}

	return sessions.Sessions("mySession", store)
}
```

登录成功时要保存session

```go
session := sessions.Default(c)
session.Set("login_user", userName)
session.Save()
```

认证中间件：

```go
// BasicAuth 最基础的认证校验 只要cookie中带了login_user标识就认为是登录用户
func BasicAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		loginUser := session.Get("login_user")
		if loginUser == nil { // 请求对应的session中找不到我想要的数据，说明不是登录的用户
			c.Redirect(http.StatusFound, "/login")
			c.Abort() // 终止当前请求的处理函数调用链
			return    // 终止当前处理函数
		}
		// 根据loginUser 去数据库里用户对象取出来  gob是go语言里面二进制的数据格式
		//userObj := queryFromMySQL(loginUser)
		//c.Set("user", userObj)
        
		// 如果是一个登录的用户，我就在c上设置两个自定义的键值对！！！
		c.Set("is_login", true)
		c.Set("login_user", loginUser)
		c.Next()
	}
}
```





数据库

- MySQL --> sqlx
- Redis  --> `redigo`/`go-redis`支持哨兵集群模式  --> `gin-session`是基于redigo的



项目的返回值设计

```go
 gin.H{
     "code": 前后端一起约定的状态码,
     "msg": 错误提示信息,
     "data": 数据
 }
```

其中错误提示信息要特别注意，不要把系统内部的错误给暴露出去！



# 表单的跳转通过js实现

例如登录表单的js控制部分，`js/blog.js`

```js
    //登录
    $("#login-form").validate({
        rules: {
            username: {
                required: true,
                rangelength: [5, 10]
            },
            password: {
                required: true,
                rangelength: [5, 10]
            }
        },
        messages: {
            username: {
                required: "请输入用户名",
                rangelength: "用户名必须是5-10位"
            },
            password: {
                required: "请输入密码",
                rangelength: "密码必须是5-10位"
            }
        },
        submitHandler: function (form) {
            var urlStr = "/login";
            // alert("urlStr:" + urlStr);
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    if (data.code === 200) {
                        window.location.href = "/";
                    }else{
                        alert(data.message);
                    }
                },
                error: function (data, status) {
                    alert("err:" + data.message + ":" + status)
                }
            });
        }
    });
```



# 需求实现

## 文章点击排行实现

需求：24小时文章的点击数排行榜

分析：

> **数据结构**
>
> 用redis的zset (sorted set:有序的集合)
>
> - key：当天的文章点击数标识  --> 分段式的key   `blog:article:count:20200315`
> - 集合的元素：每个文章的id
> - 元素的分数:每个文章的点击数
>
> **业务**
>
> 1。点击文章，阅读数加 1
>
> 每次查看文章的时候（每次请求`/article/show/:id`URL的时候）
>
> ```redis
> zincrby blog:article:count:20200315 1 id
> ```
>
> 2。文章排行
>
> 请求一个`top/:n`的URL，返回按阅读数排行前n个文章id及文章标题
>
> ```redis
> zrevrange blog:article:count:20200315 0 5 withscores
> ```
>



实现

```go
// 创建key并记录数据
// IncArticleReadCount 给指定文章的阅读数+1
func IncArticleReadCount(articleId string) (err error) {
	todayStr := time.Now().Format("20060102")
	key := fmt.Sprintf(dao.KeyArticleCount, todayStr)

	if dao.Client.Exists(key).Val() == int64(0) {
		err = dao.Client.ZIncrBy(key, 1, articleId).Err()
		dao.Client.Expire(key, time.Hour*24)
	} else {
		err = dao.Client.ZIncrBy(key, 1, articleId).Err()
	}

	return err
}

// ArticleTopN 按照阅读数排行返回前n篇文章的id和title
func ArticleTopN(c *gin.Context) {
	nStr := c.Param("n")
	n, err := strconv.ParseInt(nStr, 0, 16)
	if err != nil {
		logger.Error("ArticleTopN", zap.Any("error", err))
		c.JSON(http.StatusOK, gin.H{"code": 2001, "msg": "无效的参数"})
		return
	}
	// 调用业务逻辑层 获取返回数据结果
	articleList := logic.GetArticleReadCountTopN(n)
	logger.Error("ArticleTopN", zap.Any("articleList", articleList))
	// 3. 返回
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": articleList,
	})
	return
}

// GetArticleReadCountTopN 逻辑处理
func GetArticleReadCountTopN(n int64) []*models.Article {
	// 1. zrevrange Key 0 n-1 从redis取出前n位的文章id
	todayStr := time.Now().Format("20060102")
	key := fmt.Sprintf(dao.KeyArticleCount, todayStr)
	idStrs, err := dao.Client.ZRevRange(key, 0, n-1).Result()
	if err != nil {
		logger.Error("ZRevRange", zap.Any("error", err))
	}

	// 2. 根据上一步获取的文章id查询数据库取文章标题  ["3" "1" "5"]
	// select id, title from article where id in (3, 1, 5);  // 文章的顺序对吗？
	// 1. 让MySQL排序
	// select id, title from article where id in (3, 1, 5) order by FIND_IN_SET(id, (3, 1, 5));
	// 2. 查询出来自己排序
	var ids = make([]int64, len(idStrs))
	for _, idStr := range idStrs {
		id, err := strconv.ParseInt(idStr, 0, 16)
		if err != nil {
			logger.Warn("ArticleTopN:strconv.ParseInt failed", zap.Any("error", err))
			continue
		}
		ids = append(ids, id)
	}
	articleList, err := models.QueryArticlesByIds(ids, idStrs)
	if err != nil {
		logger.Error("queryArticlesByIds", zap.Any("error", err))
	}
	return articleList
}

// QueryArticlesByIds 在数据库中根据id查文章 按指定顺序
func QueryArticlesByIds(ids []int64, idStrs []string) ([]*Article, error) {
	sqlStr := "select id, title from article where id in (?) order by FIND_IN_SET(id, ?)"
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(idStrs, ","))
	if err != nil {
		logger.Error("QueryArticlesByIds", zap.Any("error", err))
		return nil, err
	}
	var dest []*Article
	err = dao.QueryRows(&dest, query, args...)
	return dest, err
}
```



html文件中同构ajax发送请求：

```html
        <div class="whitebg paihang">
            <h2 class="htitle">点击排行</h2>
            <ul id="today-top">

            </ul>
        </div>

<script>
    $(document).ready(function(){
        // 请求后端 /article/top/:n 接口
        $.ajax({
            url: "/article/top/5",
            type: "GET",
            success: function (data) {
                console.log(data.data);
                let s = "";
                $.each(data.data, function (idx, value) {
                    s+=`<li><i></i><a href="/article/show/${value.Id}">${value.Title}</a></li>`
                });
                $("#today-top").append(s)
            }
        })
        // 将数据渲染到 #today-top 标签中
    })
</script>
```

## 上传文件

文件是存到文件系统，数据库里存的都是路径



# TODO

排行榜防刷

分析：

> 如何防止某些人频繁的访问某篇文章刷点击？
>
> 关键点在于：如何区分正常的阅读数和不正常的阅读数。
>
> - 根据ip来
> - 根据访问的用户来
> - 24小时时间内 同一个用户对某一篇文章的点击只记录一次！
>
> ##### 实现
>
> redis set实现 记录 阅读某篇文章的用户都有哪一些
>
> - Key: blog:article:read:username:20200315 
> - 元素：用户名/用户id
>
> 用户请求`/article/show/:id`这个URL之后，我们需要：
>
> 1. 先判断 当前用户是否存在于 当天这篇文章阅读的用户set里
> 2. 如果不在，把这个用户加到阅读的用户set里，同时去给这篇文章的阅读数+1 （事务操作）
>
> #### 其他需求
>
> 一年之前的文章不参与阅读数排行榜的评选  --> 判断文章的创建时间



