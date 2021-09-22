protoc编译工具安装
```bash
#https://github.com/protocolbuffers/protobuf/releases
wget ...zip
sudo unzip protoc-3.10.0-rc-1-linux-x86_64.zip -d /usr/local/
```

安装go插件 protoc-gen-go
```bash
go get -u github.com/golang/protobuf/protoc-gen-go
```

编译生成代码文件
```bash
protoc --proto_path=. --go_out=. proto/person/person.proto
protoc --proto_path=. --go_out=. proto/message/message.proto
```
