package servercore

import (
	"fmt"
	"github.com/appcelerator/amp/cmd/cluster-server/servergrpc"
	"golang.org/x/net/context"
	"io"
)

func (s *ClusterServer) GetClientStream(stream servergrpc.ClusterServerService_GetClientStreamServer) error {
	for {
		mes, err := stream.Recv()
		if err == io.EOF {
			logf.error("Stream Server-client EOF\n")
			return nil
		}
		if err != nil {
			return fmt.Errorf("Stream Server-client error: %v\n", err)
		}
		logf.info("received client message: %v\n", mes)
	}
}

func (s *ClusterServer) DeclareAgent(ctx context.Context, req *servergrpc.DeclareRequest) (*servergrpc.ServerRet, error) {
	s.agentMap[req.Name] = &Agent{
		name: req.Name,
	}
	logf.info("Received agent declaration %s\n", req.Name)
	return &servergrpc.ServerRet{Ack: true}, nil
}
