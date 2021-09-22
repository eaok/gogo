
```bash
go run greeter/srv.go
go run greeter/web.go
micro web

#浏览器中访问
127.0.0.1:8082
```


轮询负载测试
```bash
go run greeter/srv.go
go run greeter/srv.go


#启动api服务
micro api --handler=api
go run api.go


#测试
curl localhost:8080/greeter/say/hello
```

查看服务
```
micro list services
```