package main

import (
	"fmt"
	"github.com/appcelerator/amp/cmd/swarm-server/servercore"
	"google.golang.org/grpc"
	"time"
)

// build vars
var (
	Version string
	Build   string
	client  *swarmClient  = &swarmClient{}
	conf    *ClientConfig = &ClientConfig{}
)

func main() {
	//conf.init(Version, Build)
	//client.cli()
	connectServer()
	for {
		time.Sleep(1 * time.Second)
	}
}

func connectServer() {
	conn, err := grpc.Dial("127.0.0.1:31315",
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(time.Second*20))
	if err != nil {
		fmt.Println("Error: %v\n", err)
	}
	fmt.Printf("pass\n")
	cl := servercore.NewSwarmServerServiceClient(conn)
	fmt.Printf("Connected :%v\n", cl)
	conn.Close()
}
