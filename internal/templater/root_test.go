package templater

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed test_files/exp_service.txt
var expectedService string

func TestGenerateService(t *testing.T) {
	data := Service{
		ServiceName: "Book",
		MethodNames: []string{
			"Order",
			"CancelOrder",
		},
		ResponseTypes: []string{
			"OrderResponse",
			"CancelOrderResponse",
		},
		CommonMethods: []string{
			"Order(ctx context.Context, arg *connect.Request[OrderRequest], res *connect.Response[OrderResponse]) error",
			"CancelOrder(ctx context.Context, arg *connect.Request[CancelOrderRequest], res *connect.Response[CancelOrderResponse]) error",
		},
		JsonrpcMethods: []string{
			"Order(req *http.Request, arg *connect.Request[OrderRequest], res *connect.Response[OrderResponse]) error",
			"CancelOrder(req *http.Request, arg *connect.Request[CancelOrderRequest], res *connect.Response[CancelOrderResponse]) error",
		},
		GrpcMethods: []string{
			"Order(ctx context.Context, arg *connect.Request[OrderRequest]) (res *connect.Response[OrderResponse], err error)",
			"CancelOrder(ctx context.Context, arg *connect.Request[CancelOrderRequest]) (res *connect.Response[CancelOrderResponse], err error)",
		},
	}
	content, err := generateContent(data, additionalTemplate{templateName: serviceTemplate, templateContent: serviceTemplate})
	require.NoError(t, err)
	require.Equal(t, expectedService, string(content))
}
