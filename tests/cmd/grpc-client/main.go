package main

import (
	"context"
	"fmt"
	"time"

	coreapiv1 "github.com/najeal/rpc-fusion/tests/gen/coreapi/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	client := coreapiv1.NewCoreApiServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.Ping(ctx, &coreapiv1.PingRequest{Value: 11})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(res.GetValue())
}
