package node

import (
	"github.com/appcelerator/amp/pkg/docker"
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// Server is used to implement log.LogServer
type Server struct {
	Docker *docker.Docker
}

// GetNodes implements Node.GetNodes
func (s *Server) GetNodes(ctx context.Context, in *GetNodesRequest) (*GetNodesReply, error) {
	list, err := s.Docker.GetClient().NodeList(ctx, types.NodeListOptions{})
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, "%v", err)
	}
	nodeList := &GetNodesReply{}
	for _, item := range list {
		node := &NodeEntry{
			Id:           item.ID,
			Name:         item.Spec.Name,
			Hostname:     item.Description.Hostname,
			Role:         string(item.Spec.Role),
			Architecture: item.Description.Platform.Architecture,
			Os:           item.Description.Platform.OS,
			Engine:       item.Description.Engine.EngineVersion,
			Addr:         item.Status.Addr,
			Status:       string(item.Status.State),
			Availability: string(item.Spec.Availability),
			Leader:       item.ManagerStatus.Leader,
			Reachability: string(item.ManagerStatus.Reachability),
			Labels:       item.Description.Engine.Labels,
		}
		nodeList.Entries = append(nodeList.Entries, node)
	}
	return nodeList, nil
}
