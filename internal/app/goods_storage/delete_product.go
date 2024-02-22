package goods_storage

import (
	"context"

	"github.com/Artenso/goods-storage/internal/converter"
	desc "github.com/Artenso/goods-storage/pkg/goods_storage/github.com/Artenso/goods_storage/pkg/goods_storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DeleteProduct deletes product
func (i *Implementation) DeleteProduct(ctx context.Context, req *desc.DeleteProductRequest) (*desc.DeleteProductResponse, error) {
	if err := req.ValidateAll(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid request: %s", err.Error())
	}

	if err := i.goodsStorageSrv.DeleteProduct(ctx, req.GetId()); err != nil {
		return nil, status.Errorf(codes.NotFound, "bad id: %s", err.Error())
	}
	return converter.ToDeleteProductResponse(), nil
}
