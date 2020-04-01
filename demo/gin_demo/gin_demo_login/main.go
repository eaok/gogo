package main

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID        int64     `gorm:"column:id;primary_key" json:"id"`
	Email     string    `gorm:"column:email" json:"email"`
	Password  string    `gorm:"column:password" json:"password"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:root@/demo_login?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	// 如果不设置则gorm默认表名会在后面加s（注意struct的名字要和表名一致，否则会找不到）
	db.SingularTable(true)

	//如果表不存在就创建表
	if db.HasTable(User{}) == false {
		db.CreateTable(&User{})
	}
}

func SHA(str string) string {
	shaStr := fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
	return shaStr
}

func isRegister(email string) (flag bool) {
	var user User
	db.Where("email = ?", email).First(&user)

	if user.Email == email {
		return true
	} else {
		return false
	}

}

func indexFunc(c *gin.Context) {
	c.Redirect(http.StatusFound, "/login")
}

func loginFunc(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "login.html", nil)
	} else {
		var user User

		email := c.Request.FormValue("email")
		password := c.Request.FormValue("password")
		db.Where("email = ?", email).First(&user)

		if SHA(password) != user.Password {
			c.String(http.StatusOK, "Login failed!")
		} else {
			c.String(http.StatusOK, "Login success!")
		}
	}
}

func registerFunc(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func registerPostFunc(c *gin.Context) {
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")

	//邮箱已经注册
	if isRegister(email) {
		c.Redirect(http.StatusFound, "/register")
		return
	}

	//账户信息写入数据库
	userRole := &User{
		Email:    email,
		Password: SHA(password),
	}
	db.Create(userRole)
	c.Redirect(http.StatusFound, "/login")
}

func main() {
	defer db.Close()
	router := gin.Default()

	router.Static("static", "static")
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", indexFunc)
	router.Any("/login", loginFunc)
	router.GET("/register", registerFunc)
	router.POST("/register", registerPostFunc)

	router.Run(":8080")
}
