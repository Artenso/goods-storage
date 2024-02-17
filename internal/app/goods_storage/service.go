package goods_storage

import (
	"context"

	"github.com/Artenso/goods-storage/internal/model"
	desc "github.com/Artenso/goods-storage/pkg/goods_storage/github.com/Artenso/goods_storage/pkg/goods_storage"
)

// IGoodsStorageService сервисный слой
type IGoodsStorageService interface {
	AddProduct(ctx context.Context, info *model.ProductInfo) (*model.Product, error)
}

type Implementation struct {
	desc.UnimplementedGoodsStorageServer

	goodsStorageSrv IGoodsStorageService
}

func NewGoodsStorage(goodsStorageSrv IGoodsStorageService) *Implementation {
	return &Implementation{
		UnimplementedGoodsStorageServer: desc.UnimplementedGoodsStorageServer{},

		goodsStorageSrv: goodsStorageSrv,
	}
}
