package main

import (
	"context"
	"log"
	"net"

	userpb "github.com/lokeshkarthik5/go-services/proto"
	inventorypb "github.com/lokeshkarthik5/go-services/proto"
	orderpb "github.com/lokeshkarthik5/go-services/proto"
	"google.golang.org/grpc"
)

type server struct{
	orderpb.UnimplementedOrderServiceServer
	userClient userpb.UserServiceClient
	inventoryClient inventorypb.InventoryServiceClient
}

func (s *server) CreateOrder(ctx context.Context, req *orderpb.OrderRequest) (*orderpb.OrderResponse,error) {
	//Call User Service

	userResp,err := s.userClient.GetUser(ctx,&userpb.UserRequest{
		Id:req.UserId,
	})

	if err != nil {
		return nil,err
	}

	//Call Inventory service

	stockResp,err := s.inventoryClient.CheckStock(ctx, &inventorypb.StockRequest{
		ProductId: req.ProductId,
	})

	if err != nil {
		return nil,err
	}

	if !stockResp.InStock {
		return &orderpb.OrderResponse{
			Message: "Product out of stock",
		},nil
	}

	return &orderpb.OrderResponse{
		Message: "Order created for" + userResp.Name,
	},nil
}

func main() {
	//Connect to other services
	userConn,_ := grpc.Dial("localhost:50051",grpc.WithInsecure())
	inventoryConn,_ := grpc.Dial("localhost:50052",grpc.WithInsecure())

	userClient := userpb.NewUserServiceClient(userConn)
	inventoryClient := inventorypb.NewInventoryServiceClient(inventoryConn)

	lis,_ := net.Listen("tcp",":50053")
	grpcServer := grpc.NewServer()

	orderpb.RegisterOrderServiceServer(grpcServer,&server{
		userClient: userClient,
		inventoryClient: inventoryClient,
	})

	log.Println("Order service running on :50053")
	grpcServer.Serve(lis)
}



