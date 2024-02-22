package goods_storage

import (
	"context"

	"github.com/Artenso/goods-storage/internal/model"
)

// GetProduct implements goods_storage.IGoodsStorageService.
func (s *Service) GetProduct(ctx context.Context, id int64) (*model.Product, error) {
	return s.goodsRepository.GetProduct(ctx, id)
}
