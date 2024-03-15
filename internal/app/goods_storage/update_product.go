package goods_storage

import (
	"context"

	"github.com/Artenso/goods-storage/internal/converter"
	desc "github.com/Artenso/goods-storage/pkg/goods_storage/github.com/Artenso/goods_storage/pkg/goods_storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UpdateProduct updates product name and description
func (i *Implementation) UpdateProduct(ctx context.Context, req *desc.UpdateProductRequest) (*desc.UpdateProductResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err.Error())
	}

	product, err := i.goodsStorageSrv.UpdateProduct(ctx, req.GetId(), converter.ToProductInfoForUpdate(req))
	if err != nil {
		return nil, err
	}

	return converter.ToUpdateProductResponse(product), nil
}
