package templater

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed test_files/exp_service.txt
var expectedService string

//go:embed test_files/exp_file_basic.txt
var expectedFileBasic string

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
	content, err := generateContent(data, additionalTemplate{templateName: serviceTemplateName, templateContent: serviceTemplate})
	require.NoError(t, err)
	require.Equal(t, expectedService, string(content))
}

func TestGenerateFile(t *testing.T) {
	data := File{
		PackageName: "fusioner",
		PackageImports: map[string]struct{}{
			"\"github.com/najeal/rpc-fusion/gen1\"": {},
			"\"github.com/najeal/rpc-fusion/gen2\"": {},
		},
	}

	content, err := generateContent(data, additionalTemplate{templateName: fileTemplateName, templateContent: fileTemplate})
	require.NoError(t, err)
	require.Equal(t, expectedFileBasic, string(content))
}
