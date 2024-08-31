package main

import (
	"context"
	"fmt"

	hypersdkrequester "github.com/ava-labs/hypersdk/requester"
	"github.com/najeal/rpc-fusion/pkg/requester"
	v1 "github.com/najeal/rpc-fusion/tests/gen/coreapi/v1"
	coreapiv1fusion "github.com/najeal/rpc-fusion/tests/gen/coreapi/v1/coreapifusion"
)

const (
	Endpoint    = "/pinger"
	ServiceName = "PingerService"
)

func NewJsonrpcClient(uri string) *JsonrpcClient {
	req := hypersdkrequester.New(uri, ServiceName)
	return &JsonrpcClient{
		requester: req,
	}
}

type JsonrpcClient struct {
	requester requester.Requester
}

func (c *JsonrpcClient) Ping(ctx context.Context, arg *v1.PingRequest) (res *v1.PingResponse, err error) {
	return requester.Send[v1.PingResponse](ctx, c.requester, "Ping", arg)
}

func (c *JsonrpcClient) Order(ctx context.Context, arg *v1.OrderRequest) (res *v1.OrderResponse, err error) {
	return requester.Send[v1.OrderResponse](ctx, c.requester, "Order", arg)
}

func (c *JsonrpcClient) Cancel(ctx context.Context, arg *v1.OrderRequest) (res *v1.OrderResponse, err error) {
	return requester.Send[v1.OrderResponse](ctx, c.requester, "Cancel", arg)
}

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
