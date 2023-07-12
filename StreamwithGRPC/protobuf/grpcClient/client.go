package main

import (
	"context"
	"flag"
	"io"
	"log"

	pb "github.com/atanda0x/protobuf-go/StreamwithGRPC/datafiles"
	"google.golang.org/grpc"
)

var (
	address = flag.String("add", "localhost:500051", "The address to connect")
)

func ReceiveStream(client pb.MoneyTransactionClient, request *pb.TransactionRequest) {
	log.Println("Started listening to the server stream!!!!")
	stream, err := client.MoneyTransaction(context.Background(), request)
	if err != nil {
		log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
	}

	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
		}

		log.Printf("Status: %v, Operation: %v", response.Status, response.Description)
	}
}

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewMoneyTransactionClient(conn)

	from := "1234"
	to := "5678"
	amount := float32(1250.75)

	ReceiveStream(client, &pb.TransactionRequest{From: from, To: to, Amount: amount})
}
