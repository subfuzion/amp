package servercore

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"

	//"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"google.golang.org/grpc"
)

//Server data
type SwarmServer struct {
	dockerClient *client.Client
	agentMap     map[string]*Agent
}

//Agent data
type Agent struct {
	name string
}

//AgentInit Connect to docker engine, get initial containers list and start the agent
func (s *SwarmServer) Init(version string, build string) error {
	s.trapSignal()
	conf.init(version, build)

	// Connection to Docker
	defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
	cli, err := client.NewClient(conf.dockerEngine, "v1.24", nil, defaultHeaders)
	if err != nil {
		return err
	}
	s.dockerClient = cli
	logf.info("Connected to Docker-engine\n")

	// Start server
	s.startGRPCServer()
	logf.info("Start GRPC server\n")
	return nil
}

func (s *SwarmServer) startGRPCServer() {
	serv := grpc.NewServer()
	RegisterSwarmServerServiceServer(serv, s)
	lis, err := net.Listen("tcp", ":"+conf.grpcPort)
	if err != nil {
		logf.error("gnode is unable to listen on: %s\n%v", ":"+conf.grpcPort, err)
	}
	logf.info("gnode is listening on port %s\n", ":"+conf.grpcPort)
	if err := serv.Serve(lis); err != nil {
		logf.error("Problem in gnode server: %s\n", err)
	}
}

func (s *SwarmServer) startServerReader(stream SwarmServerService_GetStreamServer) {
	for {
		mes, err := stream.Recv()
		if err == io.EOF {
			logf.error("Stream Server EOF\n")
			return
		}
		if err != nil {
			logf.error("Stream Server error: %v\n", err)
			return
		}
		logf.info("received: %v\n", mes)
	}
}

// Launch a routine to catch SIGTERM Signal
func (s *SwarmServer) trapSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)
	go func() {
		<-ch
		fmt.Println("\nswarm-server received SIGTERM signal")
		os.Exit(1)
	}()
}
