package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"golang.org/x/net/context"

	pb "github.com/moveio/server/moveio/protobuf"
)

const (
	port = ":50051"
)

type Server struct {

}

func (s Server) Move(ctx context.Context, req *pb.MoveIORequest) (*pb.MoveIOResponse, error) {
	return &pb.MoveIOResponse{Response: "Hello " + req.Name + "!"}, nil
}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
