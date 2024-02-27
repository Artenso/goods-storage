package goods_storage

import (
	"context"

	"github.com/Artenso/goods-storage/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdateProduct implements goods_storage.IGoodsStorageService.
func (s *Service) UpdateProduct(ctx context.Context, id int64, info *model.UpdateProductInfo) (*model.Product, error) {
	product, err := s.goodsRepository.UpdateProduct(ctx, id, info)
	if s.IsNotFoundError(err) {
		return product, status.Errorf(codes.NotFound, "bad id: %v, %s", id, err)
	}
	return product, err
}
