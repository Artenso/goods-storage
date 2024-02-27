package goods_storage

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DeleteProduct implements goods_storage.IGoodsStorageService.
func (s *Service) DeleteProduct(ctx context.Context, id int64) error {
	err := s.goodsRepository.DeleteProduct(ctx, id)
	if s.IsNotFoundError(err) {
		return status.Errorf(codes.NotFound, "bad id: %v, %s", id, err)
	}
	return err
}
