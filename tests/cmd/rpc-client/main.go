package main

import (
	"context"
	"fmt"

	v1 "github.com/najeal/rpc-fusion/tests/gen/coreapi/v1"
	coreapiv1fusion "github.com/najeal/rpc-fusion/tests/gen/coreapi/v1/coreapifusion"
)

func main() {
	ctx := context.Background()
	rpcClient := coreapiv1fusion.NewJsonrpcCoreApiServiceClient("http://localhost:8083/rpc")
	res, err := rpcClient.Ping(ctx, &v1.PingRequest{
		Value: 10,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Value)
}
