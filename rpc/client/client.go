package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	//	连接服务器
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("err:", err)
	}
	defer client.Close()

	args := [2]int{5, 8}
	var result int
	// 使用 rpc 方法
	err = client.Call("MathService.Add", args, &result)
	if err != nil {
		fmt.Println("client Call err:", err)
	}
	fmt.Println("result:", result)
}
