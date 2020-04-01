启动服务
```bash
nsqlookupd
nsqd --lookupd-tcp-address=127.0.0.1:4160
nsqadmin --lookupd-http-address=127.0.0.1:4161
nsq_to_file --topic=topic_demo --output-dir=/tmp --lookupd-http-address=127.0.0.1:4161
```

运行程序
```go
go run consumer.go
go run producer.go
```

网页端查看
```
localhost:4171
```