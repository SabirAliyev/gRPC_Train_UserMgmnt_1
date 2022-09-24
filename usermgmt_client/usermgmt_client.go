package main

import (
	"context"
	"log"
	"time"

	pb "example.com/go-usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Create Dial connection to gRPC server. The WithBlock() option means that this function will not return until the connection is made.
	// So it`s blocking the dial.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatal("could not close client connection")
		}
	}(conn)

	// Using this connection create the client
	c := pb.NewUserManagementClient(conn)

	// Define the context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Providing a hard code of creation new users map
	var newUsers = make(map[string]int32)
	newUsers["Alice"] = 43
	newUsers["Bob"] = 30

	for name, age := range newUsers {
		// In the loop create a response with new users using the Client API
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})
		if err != nil {
			log.Fatalf("could not create user: %v", err)
		}
		// If OK it prints the log of created users data
		log.Printf(`User Details:
NAME: %s
AGE: %d
ID: %d`, r.GetName(), r.GetAge(), r.GetId())
	}
}
