package main

import (
	"fmt"

	pb "github.com/atanda0x/protobuf-go/protofiles"
	"google.golang.org/protobuf/proto"
)

func main() {
	u := &pb.User{
		Id:    1234,
		Name:  "Atanda N",
		Email: "atandanafiu@gmail.com",
		Phones: []*pb.User_PhoneNumber{
			{
				Number: "+234", Type: pb.User_HOME,
			},
		},
	}

	u1 := &pb.User{}
	body, _ := proto.Marshal(u)
	_ = proto.Unmarshal(body, u1)
	fmt.Println("Original struct loaded from proto file:", u)
	fmt.Println("Marshaled proto data: ", body)
	fmt.Println("Unmarshaled struct: ", u1)
}
