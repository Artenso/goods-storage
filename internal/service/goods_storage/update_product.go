package goods_storage

import (
	"context"
	"errors"

	"github.com/Artenso/goods-storage/internal/model"
	"github.com/jackc/pgx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdateProduct implements goods_storage.IGoodsStorageService.
func (s *Service) UpdateProduct(ctx context.Context, id int64, info *model.UpdateProductInfo) (*model.Product, error) {
	product, err := s.goodsRepository.UpdateProduct(ctx, id, info)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Errorf(codes.InvalidArgument, "invalid request: bad id: %s", model.ErrProductNotFound.Error())
		}
		return nil, err
	}

	return product, nil
}
