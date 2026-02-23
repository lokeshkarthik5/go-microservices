package main

import (
	"context"
	"log"
	"net"
	pb "github.com/lokeshkarthik5/go-services/proto"
	"google.golang.org/grpc"
)

	

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUser( ctx context.Context, req *pb.UserRequest) (*pb.UserResponse,error) {
	return &pb.UserResponse{
		Id: req.Id,
		Name: "Lok",
	},nil
}

func main() {
	lis, _ := net.Listen("tcp",":50051")
	grpcServer := grpc.NewServer()


	pb.RegisterUserServiceServer(grpcServer,&server{})

	log.Println("User service running on :50051")
	grpcServer.Serve(lis)

}