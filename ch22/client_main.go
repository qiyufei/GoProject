package main

import (
	"fmt"
	"hello/ch22/server"
	"log"
	"net/rpc"
)

func main() {
	//goRPC()
	goHttpRpc()
}

func goHttpRpc() {
	client, e := rpc.DialHTTP("tcp", ":1234")
	if e != nil {
		log.Fatalln("dialing", e)
	}
	args := server.Args{A: 4, B: 5}
	var reply int
	e = client.Call("MainService.Add", args, &reply)
	if e != nil {
		log.Fatalln("MainService error", e)
	}
	fmt.Printf("%d + %d = %d", args.A, args.B, reply)
}

func goRPC() {
	client, err := rpc.Dial("tcp", ":1234")
	if err != nil {
		log.Fatalln("dialing", err)
	}
	args := server.Args{A: 2, B: 3}
	var reply int
	err = client.Call("MainService.Add", args, &reply)
	if err != nil {
		log.Fatalln("MainService.Add error", err)
	}
	fmt.Printf("MainService.Add: %d+%d=%d", args.A, args.B, reply)
}
