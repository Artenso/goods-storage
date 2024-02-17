package goods_storage

import (
	"context"

	"github.com/Artenso/goods-storage/internal/model"
)

// AddProduct ...
func (s *Service) AddProduct(ctx context.Context, info *model.ProductInfo) (*model.Product, error) {
	return s.goodsRepository.AddProduct(ctx, info)
}
