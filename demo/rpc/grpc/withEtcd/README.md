生成proto代码文件
```bash
protoc --go_out=plugins=grpc:. protoes/hello.proto
```

解决错误1
```
module declares its path as: go.etcd.io/bbolt
                but was required as: github.com/coreos/bbolt


replace github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4
```

etcd3.3.20 的 release 版本要求 grpc 的版本是 v1.26.0 之前的
```
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
```

把 proto-gen-go 降级到能够匹配 grpc v.1.26.0 的版本，重新生成proto代码文件
```bash
go get -u -v github.com/golang/protobuf/protoc-gen-go@v1.2.0
protoc --go_out=plugins=grpc:. protoes/hello.proto
```

测试
```bash
etcd
go run server.go -Port 3000
go run server.go -Port 3001
go run server.go -Port 3002

go run client.go
```

etcdlock etcd分布式锁的实现
```
go run etcdlock\main.go
```