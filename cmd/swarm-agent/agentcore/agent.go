package agentcore

import (
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/appcelerator/amp/cmd/swarm-server/servercore"
	//"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//Agent data
type SwarmAgent struct {
	name         string
	dockerClient *client.Client
	client       servercore.SwarmServerServiceClient
	stream       servercore.SwarmServerService_GetAgentStreamClient
	conn         *grpc.ClientConn
}

//Init Connect to docker engine, get initial containers list and start the agent
func (g *SwarmAgent) Init(version string, build string) error {
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
	fmt.Println("Connected to swarm-server")
	return g.startAgentReader(context.Background())
}

func (g *SwarmAgent) connectServer() error {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", conf.serverAddr, conf.serverPort),
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithTimeout(time.Second*20))
	if err != nil {
		return err
	}
	g.conn = conn
	g.client = servercore.NewSwarmServerServiceClient(conn)
	g.client.DeclareAgent(context.Background(), &servercore.DeclareRequest{
		Name: g.name,
	})
	return nil
}

func (g *SwarmAgent) startAgentReader(ctx context.Context) error {
	stream, err := g.client.GetAgentStream(ctx)
	if err != nil {
		return err
	}
	g.stream = stream
	for {
		mes, err := g.stream.Recv()
		if err == io.EOF {
			logf.info("Server stream EOF\n")
			//close(g.recvChan)
			return nil
		}
		if err != nil {
			return fmt.Errorf("Server stream error: %v\n", err)
		}
		//g.recvChan <- mes
		logf.info("Receive answer: %v\n", mes)
	}
}

// Launch a routine to catch SIGTERM Signal
func (g *SwarmAgent) trapSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)
	go func() {
		<-ch
		fmt.Println("\nswarm-agent received SIGTERM signal")
		g.conn.Close()
		os.Exit(1)
	}()
}
