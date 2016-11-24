package servercore

import (
	"fmt"
	"github.com/appcelerator/amp/cmd/cluster-server/servergrpc"
	"golang.org/x/net/context"
	"io"
)

func (s *SwarmServer) GetAgentStream(stream servergrpc.SwarmServerService_GetAgentStreamServer) error {
	for {
		mes, err := stream.Recv()
		if err == io.EOF {
			logf.error("Stream Server-agent EOF\n")
			return nil
		}
		if err != nil {
			return fmt.Errorf("Stream Server-agent error: %v\n", err)
		}
		logf.info("received agent message: %v\n", mes)
	}
}

func (s *SwarmServer) GetClientStream(stream servergrpc.SwarmServerService_GetClientStreamServer) error {
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

func (s *SwarmServer) DeclareAgent(ctx context.Context, req *servergrpc.DeclareRequest) (*servergrpc.ServerRet, error) {
	s.agentMap[req.Name] = &Agent{
		name: req.Name,
	}
	logf.info("Received agent declaration %s\n", req.Name)
	return &servergrpc.ServerRet{Ack: true}, nil
}
