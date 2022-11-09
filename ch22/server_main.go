package main

import (
	"hello/ch22/server"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	//goServerRPC()
	goServerHttpRpc()
}
func goServerHttpRpc() {
	rpc.RegisterName("MainService", new(server.MathService))
	rpc.HandleHTTP() //交给
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error", e)
	}
	http.Serve(l, nil)
}

func goServerRPC() {
	rpc.RegisterName("MainService", new(server.MathService))
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error", err)
	}
	rpc.Accept(l)
}
