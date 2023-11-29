// Package main implements a client for Prober service.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/CodeYourFuture/immersive-go-course/grpc-client-server/prober"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	nreq = flag.Int("nreq", 5, "Number of GET requests to send to addr")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProberClient(conn)

	// Contact the server and print out its response.
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*1)
	defer cancelFunc()

	r, err := c.DoProbes(ctx, &pb.ProbeRequest{Endpoint: "https://google.com", Nreq: int32(*nreq)})
	if err != nil {
		log.Fatalf("could not probe: %v", err)
	}
	log.Printf("Average Response Time: %f", r.GetAvgLatencyMsecs())
}
