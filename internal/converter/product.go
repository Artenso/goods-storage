package converter

import (
	"database/sql"

	"github.com/Artenso/goods-storage/internal/model"
	desc "github.com/Artenso/goods-storage/pkg/goods_storage/github.com/Artenso/goods_storage/pkg/goods_storage"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToProductInfo(req *desc.AddProductRequest) *model.ProductInfo {
	return &model.ProductInfo{
		Name:        req.GetProductInfo().GetName(),
		Description: req.GetProductInfo().GetDescription(),
	}
}

func ToDescProduct(product *model.Product) *desc.Product {
	descProduct := &desc.Product{
		Id: product.ID,
		ProductInfo: &desc.ProductInfo{
			Name:        product.Info.Name,
			Description: product.Info.Description,
		},
		CreatedAt: timestamppb.New(product.CreatedAt),
	}

	if product.UpdatedAt.Valid {
		descProduct.UpdatedAt = timestamppb.New(product.UpdatedAt.Time)
	}

	return descProduct
}

func ToAddProductResponse(product *model.Product) *desc.AddProductResponse {
	return &desc.AddProductResponse{
		Id:        product.ID,
		CreatedAt: timestamppb.New(product.CreatedAt),
	}
}

func ToGetProductResponse(product *model.Product) *desc.GetProductResponse {

	descProduct := &desc.GetProductResponse{
		Product: ToDescProduct(product),
	}

	return descProduct
}

func ToListProductResponse(productsList []*model.Product) *desc.ListProductResponse {

	responseProductsList := make([]*desc.Product, 0, len(productsList))

	for _, product := range productsList {

		descProduct := ToDescProduct(product)

		responseProductsList = append(responseProductsList, descProduct)
	}

	return &desc.ListProductResponse{
		ProductsList: responseProductsList,
	}
}

func ToProductInfoForUpdate(req *desc.UpdateProductRequest) *model.UpdateProductInfo {
	return &model.UpdateProductInfo{
		Name: sql.NullString{
			String: req.GetName().GetValue(),
			Valid:  req.GetName() != nil,
		},
		Description: sql.NullString{
			String: req.GetDescription().GetValue(),
			Valid:  req.GetDescription() != nil,
		},
	}
}

func ToUpdateProductResponse(product *model.Product) *desc.UpdateProductResponse {
	return &desc.UpdateProductResponse{
		ProductInfo: &desc.ProductInfo{
			Name:        product.Info.Name,
			Description: product.Info.Description,
		},
		CreatedAt: timestamppb.New(product.CreatedAt),
		UpdatedAt: timestamppb.New(product.UpdatedAt.Time),
	}
}
