长连接

生成代码
```
protoc --go_out=plugins=grpc:. proto/hello/hello.proto
```

run
```
.\server.exe
.\client.exe
```