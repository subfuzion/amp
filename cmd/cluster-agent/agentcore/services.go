package agentcore

import (
	"github.com/appcelerator/amp/cmd/cluster-agent/agentgrpc"
	"golang.org/x/net/context"
)

func (g *ClusterAgent) GetNodeInfo(ctx context.Context, req *agentgrpc.GetNodeInfoRequest) (*agentgrpc.NodeInfo, error) {
	return &agentgrpc.NodeInfo{}, nil
}
