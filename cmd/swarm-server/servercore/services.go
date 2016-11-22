package servercore

import (
	"golang.org/x/net/context"
)

func (s *SwarmServer) GetStream(stream SwarmServerService_GetStreamServer) error {
	s.startServerReader(stream)
	return nil
}

func (s *SwarmServer) DeclareAgent(ctx context.Context, req *DeclareRequest) (*ServerRet, error) {
	s.agentMap[req.Name] = &Agent{
		name: req.Name,
	}
	return &ServerRet{Ack: true}, nil
}
