package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	pb "github.com/jbpratt78/apis/protofiles"
)

func main() {
	p := &pb.Person{
		Id:    1234,
		Name:  "Roger B",
		Email: "bl@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-5421", Type: pb.Person_HOME},
		},
	}

	p1 := &pb.Person{}
	body, _ := proto.Marshal(p)
	_ = proto.Unmarshal(body, p1)
	fmt.Println("Original struct loaded from proto file:", p, "\n")
	fmt.Println("Marshaled proto data: ", body, "\n")
	fmt.Println("Unmarshaled struct: ", p1)
}
