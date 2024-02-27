package goods_storage

import (
	"context"
	"errors"

	"github.com/Artenso/goods-storage/internal/model"
)

type IGoodsRepository interface {
	AddProduct(_ context.Context, info *model.ProductInfo) (*model.Product, error)
	GetProduct(_ context.Context, id int64) (*model.Product, error)
	ListProduct(_ context.Context, limit, offset int64) ([]*model.Product, error)
	UpdateProduct(_ context.Context, id int64, info *model.UpdateProductInfo) (*model.Product, error)
	DeleteProduct(_ context.Context, id int64) error
}

type Service struct {
	goodsRepository IGoodsRepository
}

func New(goodsRepo IGoodsRepository) *Service {
	return &Service{
		goodsRepository: goodsRepo,
	}
}

func (s *Service) IsNotFoundError(err error) bool {
	return errors.Is(err, model.ErrProductNotFound)
}
