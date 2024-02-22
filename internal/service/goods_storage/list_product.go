package goods_storage

import (
	"context"

	"github.com/Artenso/goods-storage/internal/model"
)

// ListProduct implements goods_storage.IGoodsStorageService.
func (s *Service) ListProduct(ctx context.Context, limit int64, offset int64) ([]*model.Product, error) {
	return s.goodsRepository.ListProduct(ctx, limit, offset)
}
