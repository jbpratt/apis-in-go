package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

// holds info about args passed from the rpc client
type Args struct{}

// number to register with the rpc.Register
type TimeServer int64

func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	*reply = time.Now().Unix()
	return nil
}

func main() {
	// create a new rpc server
	timeserver := new(TimeServer)
	// register rpc server
	rpc.Register(timeserver)
	rpc.HandleHTTP()
	// Listen for requests on port 1234
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("Listen error:", e)
	}
	http.Serve(l, nil)
}
