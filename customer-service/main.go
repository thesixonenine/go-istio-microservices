package main

import (
    "context"
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
    pb "go-istio-microservices/proto"
    "google.golang.org/grpc"
)

func main() {
    r := gin.Default()

    r.GET("/customers/:id/orders", func(c *gin.Context) {
        customerID := c.Param("id")
        if customerID == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
            return
        }

        conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
        if err != nil {
            fmt.Printf("Could not connect: %v\n", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Connection failed"})
            return
        }
        defer conn.Close()

        client := pb.NewOrderServiceClient(conn)
        resp, err := client.GetOrdersByCustomerId(context.Background(), &pb.CustomerRequest{CustomerId: customerID})
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "gRPC error"})
            return
        }
        c.JSON(http.StatusOK, resp.Orders)
    })

    _ = r.Run("127.0.0.1:8080")
}
