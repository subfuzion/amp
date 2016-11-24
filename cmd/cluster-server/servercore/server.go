package servercore

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	//"github.com/docker/docker/api/types"
	"github.com/appcelerator/amp/cmd/cluster-server/servergrpc"
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
	s.agentMap = make(map[string]*Agent)
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
	logf.info("GRPC server started\n")
	for {
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (s *SwarmServer) startGRPCServer() {
	serv := grpc.NewServer()
	servergrpc.RegisterSwarmServerServiceServer(serv, s)
	go func() {
		logf.info("Starting GRPC server\n")
		lis, err := net.Listen("tcp", ":"+conf.grpcPort)
		if err != nil {
			logf.error("swarm-server is unable to listen on: %s\n%v", ":"+conf.grpcPort, err)
		}
		logf.info("swarm-server is listening on port %s\n", conf.grpcPort)
		if err := serv.Serve(lis); err != nil {
			logf.error("Problem in swarm-server: %s\n", err)
		}
	}()
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
