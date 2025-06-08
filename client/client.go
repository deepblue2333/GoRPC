package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// 定义参数结构体
type Args struct {
	A int
	B int
}

func main() {
	// DialHTTP连接到RPC服务器
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	defer client.Close()

	args := &Args{10, 20}
	var reply int
	err = client.Call("ServiceA.Add", args, &reply)
	if err != nil {
		log.Fatal("ServiceA.Add error:", err)
	}
	fmt.Printf("ServiceA.Add: %d+%d=%d\n", args.A, args.B, reply)

	var reply2 int
	divCall := client.Go("ServiceB.Sub", args, &reply2, nil)
	replyCall := <-divCall.Done
	fmt.Printf("ServiceB.Sub: %d-%d=%d\n", args.A, args.B, reply2)
	fmt.Printf("ServiceB.Sub: %d-%d=%d\n", args.A, args.B, replyCall.Reply)
}
