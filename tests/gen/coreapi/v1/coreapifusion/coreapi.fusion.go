// Code generated by rpc-fuction. DO NOT EDIT.

package coreapiv1fusion

import (
	"context"
	"net/http"
	hypersdkrequester "github.com/ava-labs/hypersdk/requester"
	"github.com/najeal/rpc-fusion/pkg/requester"
	v1 "github.com/najeal/rpc-fusion/tests/gen/coreapi/v1"
)

type CoreApiServiceCommonHandler interface {
	Ping(ctx context.Context, arg *v1.PingRequest, res *v1.PingResponse) error
	Order(ctx context.Context, arg *v1.OrderRequest, res *v1.OrderResponse) error
	Cancel(ctx context.Context, arg *v1.CancelRequest, res *v1.CancelResponse) error
}

type CoreApiServiceJsonrpcHandler interface {
	Ping(req *http.Request, arg *v1.PingRequest, res *v1.PingResponse) error
	Order(req *http.Request, arg *v1.OrderRequest, res *v1.OrderResponse) error
	Cancel(req *http.Request, arg *v1.CancelRequest, res *v1.CancelResponse) error
}

type CoreApiServiceGrpcHandler interface {
	Ping(ctx context.Context, arg *v1.PingRequest) (res *v1.PingResponse, err error)
	Order(ctx context.Context, arg *v1.OrderRequest) (res *v1.OrderResponse, err error)
	Cancel(ctx context.Context, arg *v1.CancelRequest) (res *v1.CancelResponse, err error)
}

func NewCoreApiServiceGrpcServer(commonHandler CoreApiServiceCommonHandler) *CoreApiServiceGrpcServer {
	return &CoreApiServiceGrpcServer{
		commonHandler: commonHandler,
	}
}

type CoreApiServiceGrpcServer struct {
	commonHandler CoreApiServiceCommonHandler
	v1.UnimplementedCoreApiServiceServer
}

func (s *CoreApiServiceGrpcServer) Ping(ctx context.Context, arg *v1.PingRequest) (res *v1.PingResponse, err error) {
	res = new(v1.PingResponse)
	if err := s.commonHandler.Ping(ctx, arg, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CoreApiServiceGrpcServer) Order(ctx context.Context, arg *v1.OrderRequest) (res *v1.OrderResponse, err error) {
	res = new(v1.OrderResponse)
	if err := s.commonHandler.Order(ctx, arg, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CoreApiServiceGrpcServer) Cancel(ctx context.Context, arg *v1.CancelRequest) (res *v1.CancelResponse, err error) {
	res = new(v1.CancelResponse)
	if err := s.commonHandler.Cancel(ctx, arg, res); err != nil {
		return nil, err
	}
	return res, nil
}

func NewCoreApiServiceJsonrpcServer(commonHandler CoreApiServiceCommonHandler) *CoreApiServiceJsonrpcServer {
	return &CoreApiServiceJsonrpcServer{
		commonHandler: commonHandler,
	}
}

type CoreApiServiceJsonrpcServer struct {
	commonHandler CoreApiServiceCommonHandler
}

func (s *CoreApiServiceJsonrpcServer) Ping(req *http.Request, arg *v1.PingRequest, res *v1.PingResponse) error {
	ctx := req.Context()
	return s.commonHandler.Ping(ctx, arg, res)
}

func (s *CoreApiServiceJsonrpcServer) Order(req *http.Request, arg *v1.OrderRequest, res *v1.OrderResponse) error {
	ctx := req.Context()
	return s.commonHandler.Order(ctx, arg, res)
}

func (s *CoreApiServiceJsonrpcServer) Cancel(req *http.Request, arg *v1.CancelRequest, res *v1.CancelResponse) error {
	ctx := req.Context()
	return s.commonHandler.Cancel(ctx, arg, res)
}
func NewJsonrpcCoreApiServiceClient(uri string) *JsonrpcCoreApiServiceClient {
	req := hypersdkrequester.New(uri, "CoreApiService")
	return &JsonrpcCoreApiServiceClient{
		requester: req,
	}
}

type JsonrpcCoreApiServiceClient struct {
	requester requester.Requester
}

func (c *JsonrpcCoreApiServiceClient) Ping(ctx context.Context, arg *v1.PingRequest) (res *v1.PingResponse, err error) {
	return requester.Send[v1.PingResponse](ctx, c.requester, "Ping", arg)
}

func (c *JsonrpcCoreApiServiceClient) Order(ctx context.Context, arg *v1.OrderRequest) (res *v1.OrderResponse, err error) {
	return requester.Send[v1.OrderResponse](ctx, c.requester, "Order", arg)
}

func (c *JsonrpcCoreApiServiceClient) Cancel(ctx context.Context, arg *v1.CancelRequest) (res *v1.CancelResponse, err error) {
	return requester.Send[v1.CancelResponse](ctx, c.requester, "Cancel", arg)
}