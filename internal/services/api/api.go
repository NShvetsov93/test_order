package api

import (
	"context"
)

//Storage ...
type Storage interface {
	InsertProduct(ctx context.Context, productId int32, quantity int32) error
	GetProduct(ctx context.Context, productId int32, quantity int32) (error, error)
	SelectProduct(ctx context.Context, productId int32) (int32, error)
}

//Service ...
type Service struct {
	storage Storage
}

//NewService ...
func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) AddProduct(ctx context.Context, productId int32, quantity int32) error {
	return s.storage.InsertProduct(ctx, productId, quantity)
}

func (s *Service) PullProduct(ctx context.Context, productId int32, quantity int32) (error, error) {
	return s.storage.GetProduct(ctx, productId, quantity)
}

func (s *Service) Quantity(ctx context.Context, productId int32) (int32, error) {
	return s.storage.SelectProduct(ctx, productId)
}
