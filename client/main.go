package main

import (
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalln("usage: client [IP_ADDR]")
	}

	addr := args[0]
	// we are goint ot creat an instance of DialOption, and to keep this boilerplate
	// generic, we are going to make an insecure connection to the server with the
	// insecure.NewCredentials() fuction.
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// finally, we can just call the grpc.Dial functio to create a grpc.ClientConn
	// object. This is the object that we are going to need later to call the API
	// endpoints. Lastly, this is a conneciton object, so at the end of our client's
	// lifetime, we are goign to close it:
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		if err := conn.Close(); err != nil {
			log.Fatalf("unexpected error: %v", err)
		}
	}(conn)
}
