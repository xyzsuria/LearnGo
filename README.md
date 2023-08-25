新建 go 项目

cd 到该文件夹下 `go mod init 模板名`

rpc文件夹中是简单的客户端和服务端的实现代码

varintDemo.go 演示 varint 编码和解码
获得protobuf 插件 `go get google.golang.org/protobuf`

熟悉 protocol buffer
1. 在 .protocal 文件中定义消息格式
2. 使用 protocal 编译器生成 Go 代码
3. 使用 Go 的 protocal buffer API 读写消息

protocalRpc 文件夹中
安装包 `go get "google.golang.org/protobuf"`
生成 .pb.go 文件 `protoc --go_out=./ ./addressbook.proto`