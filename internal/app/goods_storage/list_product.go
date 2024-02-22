package goods_storage

import (
	"context"

	"github.com/Artenso/goods-storage/internal/converter"
	desc "github.com/Artenso/goods-storage/pkg/goods_storage/github.com/Artenso/goods_storage/pkg/goods_storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ListProduct gets products list with pagination
func (i *Implementation) ListProduct(ctx context.Context, req *desc.ListProductRequest) (*desc.ListProductResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err.Error())
	}

	productsList, err := i.goodsStorageSrv.ListProduct(ctx, req.GetLimit(), req.GetOffset())
	if err != nil {
		return nil, status.Errorf(codes.OutOfRange, "bad offset: %s", err.Error())
	}
	return converter.ToListProductResponse(productsList), nil
}
