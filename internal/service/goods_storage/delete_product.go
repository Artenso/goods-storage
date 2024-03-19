package goods_storage

import (
	"context"

	"github.com/Artenso/goods-storage/internal/model"
	"github.com/jackc/pgx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DeleteProduct implements goods_storage.IGoodsStorageService.
func (s *Service) DeleteProduct(ctx context.Context, id int64) error {
	err := s.goodsRepository.DeleteProduct(ctx, id)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return status.Errorf(codes.InvalidArgument, "invalid request: bad id: %s", model.ErrProductNotFound.Error())
		}
		return err
	}

	return nil
}
