package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A int
	B int
}

type ServiceA struct{}

func (s *ServiceA) Add(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

type ServiceB struct{}

func (s *ServiceB) Sub(args *Args, reply *int) error {
	*reply = args.A - args.B
	return nil
}

func main() {
	serviceA := new(ServiceA)
	serviceB := new(ServiceB)
	rpc.Register(serviceA)
	rpc.Register(serviceB)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	http.Serve(listener, nil)
}
