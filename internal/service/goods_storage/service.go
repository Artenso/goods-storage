package goods_storage

import (
	"context"

	"github.com/Artenso/goods-storage/internal/model"
)

type IGoodsRepository interface {
	AddProduct(ctx context.Context, info *model.ProductInfo) (*model.Product, error)
}

type Service struct {
	goodsRepository IGoodsRepository
}

func New(goodsRepo IGoodsRepository) *Service {
	return &Service{
		goodsRepository: goodsRepo,
	}
}
