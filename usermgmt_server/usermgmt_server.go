package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "example.com/go-usermgmt-grpc/usermgmt"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// UserManagementServer is the implementation of gRPC service. And we have to embed this type to gRPC register
type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
}

// CreateNewUser is the implementation of our proto file method as a receiver function
func (s *UserManagementServer) CreateNewUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	// Log received username
	log.Printf("Received: %v", in.GetName())
	// Generate user ID with random int32
	var userId = int32(rand.Intn(1000))
	// Return the address of protobuf user with input fields
	return &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: userId}, nil
}

func main() {
	// Create listener with specified port number using net package
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Define new variable as gRPC server by invoking NewServer() function
	s := grpc.NewServer()
	// Register the server as new gRPC service by passing our new variable and address of UserManagementServer{} type
	pb.RegisterUserManagementServer(s, &UserManagementServer{})
	log.Printf("server listening at %v", lis.Addr())
	// Start the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
