package main

import (
	"context"
	"fmt"
	"os"

	sk "github.com/appcelerator/amp/cluster/swarmkit"
	"github.com/docker/swarmkit/api"
	"google.golang.org/grpc/status"
)

func main() {
	c, conn, err := sk.Dial(sk.DefaultSocket())
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			fmt.Println("Error: ", s)
		} else {
			fmt.Println("Error:", err)
		}
		os.Exit(-1)
	}

	// this is just to prove things are working...
	nodes, err := sk.ListNodes(sk.Context(), c, sk.LiveNodesFilter)
	for _, n := range nodes {
		fmt.Println("Node ID: ", n.GetID())
	}

	watcher := api.NewWatchClient(conn)
	watchEntry := sk.NewWatchRequestEntry("service", sk.WatchActionKindAll, nil)
	watchEntries := []*api.WatchRequest_WatchEntry{
		watchEntry,
	}

	// will probably need a cancelable context
	ctx := context.TODO()
	in := sk.NewWatchRequest(watchEntries, nil, true)
	w, err := watcher.Watch(ctx, in)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	for {
		msg, err := w.Recv()
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		fmt.Println(msg.String())
	}
}
