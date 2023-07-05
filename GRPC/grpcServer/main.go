package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/atanda0x/protobuf-go/GRPC/protofile"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedMoneyTransferedServer
}

func main() {

	s := grpc.NewServer()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	pb.RegisterMoneyTransferedServer(s, &server{})
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func (s *server) MakeTransaction(ctx context.Context, in *pb.TransferRequest) (*pb.TransferResponse, error) {
	// Business logic will come here
	fmt.Println("Got amount ", in.Amount)
	fmt.Println("Got from ", in.From)
	fmt.Println("For ", in.To)
	// Returning a response of type Transaction Response
	return &pb.TransferResponse{Confirmation: true}, nil
}
