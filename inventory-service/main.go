package main

import (
	"context"
	"log"
	"net"

	pb "github.com/lokeshkarthik5/go-services/proto"
	"google.golang.org/grpc"
)

type server struct{
	pb.UnimplementedInventoryServiceServer
}

func (s *server) CheckStock(ctx context.Context,req *pb.StockRequest) (*pb.StockResponse,error){
	return &pb.StockResponse{
		InStock: true,
	},nil
}

func main(){
	lis,_ := net.Listen("tcp",":50052")
	grpcServer := grpc.NewServer()

	pb.RegisterInventoryServiceServer(grpcServer,&server{})

	log.Println("Inventory service running on :50052")
	grpcServer.Serve(lis)
}