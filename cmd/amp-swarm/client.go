package main

import (
	"fmt"
	"github.com/appcelerator/amp/cmd/swarm-server/servercore"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"time"
)

type swarmClient struct {
	name     string
	client   servercore.SwarmServerServiceClient
	stream   servercore.SwarmServerService_GetClientStreamClient
	conn     *grpc.ClientConn
	ctx      context.Context
	verbose  bool
	id       int
	nodeName string
	nodeHost string
	recvChan chan *servercore.ClientMes
}

var (
	RootCmd = &cobra.Command{
		Use:   `amp-swarm`,
		Short: "amp-swarm ",
		Run: func(cmd *cobra.Command, args []string) {
			//fmt.Println(cmd.UsageString())
		},
	}
)

func (g *swarmClient) init() error {
	g.ctx = context.Background()
	g.recvChan = make(chan *servercore.ClientMes)
	if err := g.connectServer(); err != nil {
		return err
	}
	if err := g.startServerReader(); err != nil {
		return err
	}
	g.printf("Client connected to swarm-server\n")
	return nil
}

func (g *swarmClient) connectServer() error {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", conf.serverAddr, conf.serverPort),
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(time.Second*20))
	if err != nil {
		return err
	}
	g.conn = conn
	g.client = servercore.NewSwarmServerServiceClient(conn)
	return nil
}

func (g *swarmClient) createSendMessageNoAnswer(target string, functionName string, args ...string) error {
	mes := &servercore.ClientMes{} //TODO
	_, err := g.sendMessage(mes, true)
	return err
}

func (g *swarmClient) createSendMessage(target string, waitForAnswer bool, functionName string, args ...string) (*servercore.ClientMes, error) {
	mes := &servercore.ClientMes{} //TODO
	return g.sendMessage(mes, waitForAnswer)
}

func (g *swarmClient) sendMessage(mes *servercore.ClientMes, wait bool) (*servercore.ClientMes, error) {
	err := g.stream.Send(mes)
	if err != nil {
		return nil, err
	}
	g.printf("Message sent: %v\n", mes)
	if wait {
		ret := <-client.recvChan
		return ret, nil
	}
	return nil, nil
}

func (g *swarmClient) getNextAnswer() *servercore.ClientMes {
	mes := <-g.recvChan
	return mes
}

func (g *swarmClient) startServerReader() error {
	stream, err := g.client.GetClientStream(g.ctx)
	if err != nil {
		return err
	}
	g.stream = stream
	go func() {
		for {
			mes, err := g.stream.Recv()
			if err == io.EOF {
				g.printf("Server stream EOF\n")
				close(g.recvChan)
				return
			}
			if err != nil {
				g.printf("Server stream error: %v\n", err)
				return
			}
			g.recvChan <- mes
			g.printf("Receive answer: %v\n", mes)
		}
	}()
	return nil
}

func (g *swarmClient) printf(format string, args ...interface{}) {
	//TODO color/mode
	fmt.Printf(format, args...)
}
