package main

import (
	"flag"

	pb "github.com/atanda0x/protobuf-go/StreamwithGRPC/datafiles"
)

var (
	port = flag.Int("port", 50051, "Server port")
)

type server struct {
	pb.UnimplementedMoneyTransferServer
}
