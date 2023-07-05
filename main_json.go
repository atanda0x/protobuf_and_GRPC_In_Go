package main

import (
	"encoding/json"
	"fmt"

	pb "github.com/atanda0x/protobuf-go/protofiles"
)

func main() {
	u := &pb.User{
		Id:    1234,
		Name:  "Atanda0x",
		Email: "atandanafiu@gmail.com",
		Phones: []*pb.User_PhoneNumber{
			{
				Number: "+23490", Type: pb.User_HOME,
			},
		},
	}

	body, _ := json.Marshal(u)
	fmt.Println(string(body))
}
