package goods_storage

import (
	"context"

	"github.com/Artenso/goods-storage/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetProduct implements goods_storage.IGoodsStorageService.
func (s *Service) GetProduct(ctx context.Context, id int64) (*model.Product, error) {
	product, err := s.goodsRepository.GetProduct(ctx, id)
	if err != nil {
		if s.IsNotFoundError(err) {
			return nil, status.Errorf(codes.NotFound, "bad id: %v, %s", id, err)
		}
		return nil, err
	}
	return product, nil
}
