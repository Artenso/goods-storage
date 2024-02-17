package converter

import (
	"github.com/Artenso/goods-storage/internal/model"
	desc "github.com/Artenso/goods-storage/pkg/goods_storage/github.com/Artenso/goods_storage/pkg/goods_storage"
)

func ToProductInfo(req *desc.AddProductRequest) *model.ProductInfo {
	return &model.ProductInfo{
		Name:        req.GetName(),
		Description: req.GetDescription(),
	}
}

func ToAddProductResponse(product *model.Product) *desc.AddProductResponse {
	return &desc.AddProductResponse{Id: product.ID}
}
