package main

import (
	"context"
	"io"
	"log"
	proto "micro_demo/fileWithProto/proto"
	"net/http"

	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
)

var c client.Client
var fileService proto.FileService

func UploadFile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	// 取到文件对象
	files, ok := r.MultipartForm.File["file"]
	if !ok {
		w.WriteHeader(400)
		_, _ = w.Write([]byte("请选择文件上传"))
		return
	}

	// 将文件通过流式传输到srv
	file, err := files[0].Open()
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	// 建立链接
	// 因为这里是用的临时文件储存的方式,如果因为负载均衡算法导致下一次节点切换,另外一个节点是无法通过,文件名来获取到文件数据的
	// 使用这种方法来固定一个节点
	next, _ := c.Options().Selector.Select("file.service")
	node, _ := next()
	stream, err := fileService.File(context.Background(), func(options *client.CallOptions) {
		// 指定节点
		options.Address = []string{node.Address}
	})
	if err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	for {
		// 缓冲1MB,每次发送1MB的内容,注意不能超过rpc的限制(grpc默认为4MB)
		buff := make([]byte, 1024*1024)
		sendLen, err := file.Read(buff)
		if err != nil {
			if err == io.EOF {
				//全部读取完成,发送一个完成标识,跳出
				err = stream.Send(&proto.FileSlice{
					Byte: nil,
					Len:  -1,
				})
				if err != nil {
					w.WriteHeader(500)
					_, _ = w.Write([]byte(err.Error()))
					return
				}
				break
			}
			w.WriteHeader(500)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		err = stream.Send(&proto.FileSlice{
			Byte: buff[:sendLen],
			Len:  int64(sendLen),
		})
		if err != nil {
			w.WriteHeader(500)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
	}

	// 等待接收，当收到的消息之后就可以关闭了
	fileMsg := &proto.FileSliceMsg{}
	if err := stream.RecvMsg(fileMsg); err != nil {
		w.WriteHeader(500)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	_ = stream.Close()
	println(fileMsg.FileName)
}

func main() {
	service := micro.NewService(
		micro.Name("file.client"),
	)
	service.Init()

	// 创建客户端
	c = service.Client()
	fileService = proto.NewFileService("file.service", c)

	http.HandleFunc("/upload", UploadFile)

	log.Println("Listening at 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
