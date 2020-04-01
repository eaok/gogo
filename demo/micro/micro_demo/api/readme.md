
Run the micro API
```
micro api --handler=api
```

Run this example
```
go run api.go
```

Make a request
```
curl http://localhost:8080/example/foo/bar?name=john            //get
curl -d '{"name":"john"}' http://localhost:8080/example/foo/bar //post

curl http://localhost:8080/example/example/call?name=john
curl -d '{"name":"john"}' http://localhost:8080/example/example/call
```

查看服务
```
micro list services
```