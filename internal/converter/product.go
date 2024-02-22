package converter

import (
	"github.com/Artenso/goods-storage/internal/model"
	desc "github.com/Artenso/goods-storage/pkg/goods_storage/github.com/Artenso/goods_storage/pkg/goods_storage"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToProductInfo(req *desc.AddProductRequest) *model.ProductInfo {
	return &model.ProductInfo{
		Name:        req.GetProductInfo().GetName(),
		Description: req.GetProductInfo().GetDescription(),
	}
}

func ToAddProductResponse(product *model.Product) *desc.AddProductResponse {
	return &desc.AddProductResponse{
		Id:        product.ID,
		CreatedAt: timestamppb.New(product.CreatedAt),
	}
}

func ToGetProductResponse(product *model.Product) *desc.GetProductResponse {

	if product.UpdatedAt.Valid {
		return &desc.GetProductResponse{
			ProductInfo: &desc.ProductInfo{
				Name:        product.Info.Name,
				Description: product.Info.Description,
			},
			CreatedAt: timestamppb.New(product.CreatedAt),
			UpdatedAt: timestamppb.New(product.UpdatedAt.Time),
		}
	}
	return &desc.GetProductResponse{
		ProductInfo: &desc.ProductInfo{
			Name:        product.Info.Name,
			Description: product.Info.Description,
		},
		CreatedAt: timestamppb.New(product.CreatedAt),
	}
}

func ToDeleteProductResponse() *desc.DeleteProductResponse {
	return &desc.DeleteProductResponse{
		EmptyStruct: &emptypb.Empty{},
	}
}

func ToListProductResponse(productsList []*model.Product) *desc.ListProductResponse {

	responseProductsList := make([]*desc.ListProductResponse_Product, 0, len(productsList))

	for i := 0; i < len(productsList); i++ {
		product := desc.ListProductResponse_Product{
			Id: productsList[i].ID,
			ProductInfo: &desc.ProductInfo{
				Name:        productsList[i].Info.Name,
				Description: productsList[i].Info.Description,
			},
			CreatedAt: timestamppb.New(productsList[i].CreatedAt),
		}

		if productsList[i].UpdatedAt.Valid {
			product.UpdatedAt = timestamppb.New(productsList[i].UpdatedAt.Time)
		}

		responseProductsList = append(responseProductsList, &product)
	}

	return &desc.ListProductResponse{
		ProductsList: responseProductsList,
	}
}

func ToProductInfoForUpdate(req *desc.UpdateProductRequest) *model.ProductInfo {
	return &model.ProductInfo{
		Name:        req.GetName(),
		Description: req.GetDescription(),
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
