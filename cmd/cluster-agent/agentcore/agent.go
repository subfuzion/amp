package agentcore

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/appcelerator/amp/cmd/cluster-agent/agentgrpc"
	"github.com/appcelerator/amp/cmd/cluster-server/servergrpc"
	//"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//Agent data
type ClusterAgent struct {
	name         string
	dockerClient *client.Client
	client       servergrpc.ClusterServerServiceClient
	conn         *grpc.ClientConn
}

//Init Connect to docker engine, get initial containers list and start the agent
func (g *ClusterAgent) Init(version string, build string) error {
	g.trapSignal()
	conf.init(version, build)
	g.name = os.Getenv("HOSTNAME")

	// Connection to Docker
	defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
	cli, err := client.NewClient(conf.dockerEngine, "v1.24", nil, defaultHeaders)
	if err != nil {
		return err
	}
	g.dockerClient = cli
	fmt.Println("Connected to Docker-engine")

	// Connection to server
	if err := g.connectServer(); err != nil {
		return err
	}
	fmt.Println("Connected to cluster-server")
	g.startGRPCServer()
	return nil
}

func (g *ClusterAgent) connectServer() error {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", conf.serverAddr, conf.serverPort),
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(time.Second*20))
	if err != nil {
		return err
	}
	g.conn = conn
	g.client = servergrpc.NewClusterServerServiceClient(conn)
	g.client.DeclareAgent(context.Background(), &servergrpc.DeclareRequest{
		Name: g.name,
	})
	return nil
}

func (g *ClusterAgent) startGRPCServer() {
	serv := grpc.NewServer()
	agentgrpc.RegisterClusterAgentServiceServer(serv, g)
	logf.info("Starting GRPC server\n")
	lis, err := net.Listen("tcp", ":"+conf.grpcPort)
	if err != nil {
		logf.error("cluster-agent is unable to listen on: %s\n%v", ":"+conf.grpcPort, err)
		return
	}
	logf.info("cluster-agent is listening on port %s\n", conf.grpcPort)
	if err := serv.Serve(lis); err != nil {
		logf.error("Problem in cluster-agent: %s\n", err)
		return
	}
}

// Launch a routine to catch SIGTERM Signal
func (g *ClusterAgent) trapSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)
	go func() {
		<-ch
		fmt.Println("cluster-agent received SIGTERM signal")
		g.conn.Close()
		os.Exit(1)
	}()
}
