package main

import (
	"encoding/json"
	"fmt"

	pb "github.com/jbpratt78/apis/protofiles"
)

func main() {
	p := &pb.Person{
		Id:    1243,
		Name:  "Majora P",
		Email: "mp@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4312", Type: pb.Person_HOME},
		},
	}
	body, _ := json.Marshal(p)
	fmt.Println(string(body))
}
