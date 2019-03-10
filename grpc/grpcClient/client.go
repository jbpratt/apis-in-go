package main

import (
	"context"
	"log"

	pb "github.com/jbpratt78/apis/grpc/datafiles"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// set up conn to server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMoneyTransactionClient(conn)

	// prep data, received from clients
	from := "1234"
	to := "5678"
	amount := float32(1250.75)

	// contact server, print response
	r, err := c.MakeTransaction(context.Background(), &pb.TransactionRequest{From: from, To: to, Amount: amount})
	if err != nil {
		log.Fatalf("Could not transact: %v", err)
	}
	log.Printf("Transaction confirmed: %t", r.Confirmation)
}
