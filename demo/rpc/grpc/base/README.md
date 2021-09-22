生成代码
```go
protoc --go_out=plugins=grpc:. proto/hello/*.proto
```

run
```
.\server.exe
.\client.exe
```