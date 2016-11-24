package main

import (
	"github.com/nats-io/go-nats-streaming"
	"log"
	"os"
	"os/signal"
)

// build vars
var (
	// Version is set with a linker flag (see Makefile)
	Version string

	// Build is set with a linker flag (see Makefile)
	Build string
)

const (
	clusterID = "test-cluster"
	clientID  = "amp-event-worker"
	natsURL   = "nats://nats:4222"
	natsTopic = "amp-events"
)

func main() {
	log.Printf("amp-event-worker (version: %s, build: %s)\n", Version, Build)

	log.Printf("Connecting to nats on %s\n", natsURL)
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	if err != nil {
		log.Fatalf("Unable to connect to nats on %s: %s", natsURL, err)
	}
	log.Printf("Connected to nats on %s\n", natsURL)

	_, err = sc.Subscribe(natsTopic, messageHandler, stan.DeliverAllAvailable(), stan.DurableName(natsTopic+"-durable"))
	if err != nil {
		sc.Close()
		log.Fatalf("Unable to subscribe to %s topic: %s", natsTopic, err)
	}
	log.Printf("Listening on", natsTopic)

	// Wait for a SIGINT (perhaps triggered by user with CTRL-C)
	// Run cleanup when signal is received
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			log.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
			sc.Close()
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}

func messageHandler(msg *stan.Msg) {
	//logEntry := logs.LogEntry{}
	//err := proto.Unmarshal(msg.Data, &logEntry)
	//if err != nil {
	//	log.Printf("error unmarshalling log entry: %v", err)
	//}
	//timestamp, err := time.Parse(time.RFC3339Nano, logEntry.Timestamp)
	//if err != nil {
	//	log.Printf("error parsing timestamp: %v", err)
	//}
	//logEntry.Timestamp = timestamp.Format("2006-01-02T15:04:05.999")
	//err = es.Index(esIndex, esType, logEntry)
	//if err != nil {
	//	log.Printf("error indexing log entry: %v", err)
	//}
}
