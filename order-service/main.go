package main

import (
    "context"
    "fmt"
    "net"
    "os"

    pb "go-istio-microservices/proto"

    "google.golang.org/grpc"
)

type orderServiceServer struct {
    pb.UnimplementedOrderServiceServer
}

func (s *orderServiceServer) GetOrdersByCustomerId(ctx context.Context, req *pb.CustomerRequest) (*pb.OrderList, error) {
    orders := []*pb.Order{
        {Id: "1", Item: "Book", Price: 10.5},
        {Id: "2", Item: "Pen", Price: 1.2},
    }
    fmt.Printf("GetOrdersByCustomerId: %v", req.CustomerId)
    var result []*pb.Order
    for _, order := range orders {
        if order.Id == req.CustomerId {
            result = append(result, order)
        }
    }
    return &pb.OrderList{Orders: result}, nil
}

func main() {
    lis, err := net.Listen("tcp", "127.0.0.1:50051")
    if err != nil {
        fmt.Printf("failed to listen: %v\n", err)
        os.Exit(1)
    }
    s := grpc.NewServer()
    pb.RegisterOrderServiceServer(s, &orderServiceServer{})
    fmt.Println("Order service running on :50051")
    if err := s.Serve(lis); err != nil {
        fmt.Printf("failed to serve: %v\n", err)
        os.Exit(1)
    }
}
