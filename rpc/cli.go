package main

import (
	"fmt"

	"github.com/genekuo/microservices/rpc/client"
	"github.com/genekuo/microservices/rpc/server"
)

func main() {
	go server.StartServer()

	c := client.CreateCient()
	defer c.Close()

	reply := client.PerformRequest(c)
	fmt.Println(reply.Message)
}
