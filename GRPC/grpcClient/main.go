package main

import (
	"flag"
	"log"

	pb "github.com/atanda0x/protobuf-go/GRPC/protofile"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	address = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	conn, err := grpc.Dial(*address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMoneyTransferedClient(conn)

	from := "1234"
	to := "5678"
	amount := float32(1250.75)

	r, err := c.MoneyTransfered(context.Background(), &pb.TransferRequest{
		From:   from,
		To:     to,
		Amount: amount,
	})
	if err != nil {
		log.Fatalf("Could not transfer %v", err)
	}
	log.Printf("Transfer confirmed: %t", r.Confirmation)
}
