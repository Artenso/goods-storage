package goods_storage

import (
	"context"

	"github.com/Artenso/goods-storage/internal/converter"
	desc "github.com/Artenso/goods-storage/pkg/goods_storage/github.com/Artenso/goods_storage/pkg/goods_storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetProduct gets product
func (i *Implementation) GetProduct(ctx context.Context, req *desc.GetProductRequest) (*desc.GetProductResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err.Error())
	}

	product, err := i.goodsStorageSrv.GetProduct(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return converter.ToGetProductResponse(product), nil
}
