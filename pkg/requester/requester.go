package requester

import (
	"context"

	hypersdkrequester "github.com/ava-labs/hypersdk/requester"
)

type Requester interface {
	SendRequest(ctx context.Context, method string, arg interface{}, res interface{}, opts ...hypersdkrequester.Option) error
}

func Send[K any](ctx context.Context, requester Requester, method string, arg interface{}) (res *K, err error) {
	res = new(K)
	if err = requester.SendRequest(ctx, method, arg, res); err != nil {
		return nil, err
	}
	return res, nil
}
