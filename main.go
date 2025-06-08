package main

import (
	"fmt"
	"os"
	"time"
)

type Args struct {
	A int
	B int
}

func Add(args *Args) int {
	return args.A + args.B
}

func main() {
	fmt.Println(os.Getpid())
	fmt.Println(Add(&Args{10, 20}))
	time.Sleep(100 * time.Second)
}
