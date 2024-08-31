package cmd

import (
	"context"
	"fmt"

	coreapiv1 "github.com/najeal/rpc-fusion/tests/gen/coreapi/v1"
)

func NewCommonServer() *CommonServer {
	return &CommonServer{}
}

type CommonServer struct{}

func (s *CommonServer) Ping(ctx context.Context, arg *coreapiv1.PingRequest, res *coreapiv1.PingResponse) error {
	res.Value = arg.GetValue()
	return nil
}

func (s *CommonServer) Order(ctx context.Context, arg *coreapiv1.OrderRequest, res *coreapiv1.OrderResponse) error {
	return fmt.Errorf("not implemented")
}

func (s *CommonServer) Cancel(ctx context.Context, arg *coreapiv1.CancelRequest, res *coreapiv1.CancelResponse) error {
	return fmt.Errorf("not implemented")
}
