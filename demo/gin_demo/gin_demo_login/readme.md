
新建数据库`demo_login`
```bash
create database demo_login character set utf8mb4;
```

执行`main.go`
```bash
go run main.go
```

测试
```
http://localhost:8080/register
http://localhost:8080/login
```

本demo使用了gin框架和gorm