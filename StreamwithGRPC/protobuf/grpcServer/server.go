package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/atanda0x/protobuf-go/StreamwithGRPC/datafiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	port      = flag.Int("port", 50051, "Server port")
	noOfSteps = 3
)

type server struct {
	pb.UnimplementedMoneyTransactionServer
}

func (s *server) makeTransaction(in *pb.TransactionRequest, stream pb.MoneyTransaction_MoneyTransactionServer) error {
	log.Printf("Got request for money.....")
	log.Printf("Amount: $ %f, From A/c:%s, To A/c: %s", in.Amount, in.From, in.To)

	for i := 0; i < noOfSteps; i++ {
		time.Sleep(time.Second * 2)

		if err := stream.Send(&pb.TransactionResponse{Status: "good", Step: int32(i), Description: fmt.Sprintln("Description of step &d", int(i))}); err != nil {
			log.Fatalf("%v.Send(%v)  = %v", stream, "status", err)
		}
	}
	log.Printf("Successfully transfered amount $%v from %v to %v", in.Amount, in.From, in.To)
	return nil
}

func main() {
	flag.Parse()
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	pb.RegisterMoneyTransactionServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
