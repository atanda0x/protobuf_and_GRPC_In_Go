package main

import (
	"flag"
	"log"
	"time"

	pb "github.com/atanda0x/protobuf-go/GRPC/protofile"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	address = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMoneyTransferedClient(conn)

	from := "atanda0x"
	to := "atanda nafiu"
	amount := float32(1250.75)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.MoneyTransfered(ctx, &pb.TransferRequest{
		From:   from,
		To:     to,
		Amount: amount,
	})
	if err != nil {
		log.Fatalf("Could not transfer %v", err)
	}
	log.Printf("Transfer confirmed: %t", r.Confirmation)
}
