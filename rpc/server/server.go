package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/genekuo/microservices/rpc/contract"
)

const port = 1234

type HelloWorldHandler struct{}

func (h *HelloWorldHandler) HelloWorld(request *contract.HelloWordRequest, reply *contract.HelloWorldResponse) error {
	reply.Message = "Hello " + request.Name
	return nil
}

func StartServer() {
	helloWolrd := &HelloWorldHandler{}
	rpc.Register(helloWolrd)

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to listen on the given port: %s", err))
	}
	defer l.Close()

	for {
		conn, _ := l.Accept()
		go rpc.ServeConn(conn)
	}
}
