package api

import (
	"context"
)

// Implementation ...
type Implementation struct {
	Api API
}

//API ...
type API interface {
	AddProduct(ctx context.Context, productId int32, quantity int32) error
	PullProduct(ctx context.Context, productId int32, quantity int32) (error, error)
	Quantity(ctx context.Context, productId int32) (int32, error)
}

//NewAPI ...
func NewAPI(api API) *Implementation {
	return &Implementation{
		Api: api,
	}
}
