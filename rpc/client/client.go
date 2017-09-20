package client

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/genekuo/microservices/rpc/contract"
)

const port = 1234

func CreateCient() *rpc.Client {
	client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal("dailing", err)
	}
	return client
}

func PerformRequest(client *rpc.Client) contract.HelloWorldResponse {
	request := &contract.HelloWordRequest{Name: "World"}
	var reply contract.HelloWorldResponse

	err := client.Call("HelloWorldHandler.HelloWorld", request, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}

	return reply
}
