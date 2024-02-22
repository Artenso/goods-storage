package goods_storage

import (
	"context"

	"github.com/Artenso/goods-storage/internal/model"
)

// UpdateProduct implements goods_storage.IGoodsStorageService.
func (s *Service) UpdateProduct(ctx context.Context, id int64, info *model.ProductInfo) (*model.Product, error) {
	return s.goodsRepository.UpdateProduct(ctx, id, info)
}
