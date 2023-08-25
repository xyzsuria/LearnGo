package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type MathService struct {
}

// Add 定义一个方法，用于加法运算
func (m *MathService) Add(args [2]int, reply *int) error {
	*reply = args[0] + args[1]
	return nil
}

func main() {
	//	创建 MathService 实例
	mathService := new(MathService)
	//	注册服务
	err := rpc.Register(mathService)
	if err != nil {
		fmt.Println(err)
	}
	//	监听
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("监听出错")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
