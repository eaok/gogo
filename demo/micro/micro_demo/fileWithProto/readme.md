
# 运行
```bash
go run server.go

go run client.go
```

测试：
```bash
curl localhost:8080/upload -F "file=@readme.md?"
```