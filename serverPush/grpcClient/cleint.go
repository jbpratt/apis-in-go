package main

import (
	"context"
	"io"
	"log"

	pb "github.com/jbpratt78/apis/serverPush/datafiles"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

// ReceiveStream listens to the stream contents and uses them
func ReceiveStream(client pb.MoneyTransactionClient, request *pb.TransactionRequest) {
	log.Println("Started listening to the server stream...")
	stream, err := client.MakeTransaction(context.Background(), request)
	if err != nil {
		log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
	}
	// listen to the stream of messages
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			// if there are no more messages, get out of loop
			break
		}
		if err != nil {
			log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
		}
		log.Printf("Status: %v, Operation: %v", response.Status, response.Description)
	}
}

func main() {
	// set up conn to server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewMoneyTransactionClient(conn)

	// prep data, received from front end app
	from := "1234"
	to := "5678"
	amount := float32(1250.75)

	// contact the server and print its response
	ReceiveStream(client, &pb.TransactionRequest{From: from, To: to, Amount: amount})
}
