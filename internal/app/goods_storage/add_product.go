package goods_storage

import (
	"context"

	"github.com/Artenso/goods-storage/internal/converter"
	desc "github.com/Artenso/goods-storage/pkg/goods_storage/github.com/Artenso/goods_storage/pkg/goods_storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AddProduct adds product
func (i *Implementation) AddProduct(ctx context.Context, req *desc.AddProductRequest) (*desc.AddProductResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err.Error())
	}

	product, err := i.goodsStorageSrv.AddProduct(ctx, converter.ToProductInfo(req))
	if err != nil {
		return nil, err
	}

	return converter.ToAddProductResponse(product), nil
}
