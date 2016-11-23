package servercore

import (
	"fmt"
	"golang.org/x/net/context"
	"io"
)

func (s *SwarmServer) GetAgentStream(stream SwarmServerService_GetAgentStreamServer) error {
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

func (s *SwarmServer) GetClientStream(stream SwarmServerService_GetClientStreamServer) error {
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

func (s *SwarmServer) DeclareAgent(ctx context.Context, req *DeclareRequest) (*ServerRet, error) {
	s.agentMap[req.Name] = &Agent{
		name: req.Name,
	}
	logf.info("Received agent declaration %s\n", req.Name)
	return &ServerRet{Ack: true}, nil
}
