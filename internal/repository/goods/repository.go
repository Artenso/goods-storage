package goods

import (
	"context"
	"sync"

	"github.com/Artenso/goods-storage/internal/model"
)

type IRepository interface {
	AddProduct(_ context.Context, info *model.ProductInfo) (*model.Product, error)
	GetProduct(_ context.Context, id int64) (*model.Product, error)
	ListProduct(_ context.Context, limit, offset int64) ([]*model.Product, error)
	UpdateProduct(_ context.Context, id int64, info *model.ProductInfo) (*model.Product, error)
	DeleteProduct(_ context.Context, id int64) error
}

type Repository struct {
	mu     sync.RWMutex
	goods  []*model.Product
	lastId int64
}

func New() *Repository {
	return &Repository{
		mu:     sync.RWMutex{},
		goods:  make([]*model.Product, 0, 10),
		lastId: 0,
	}
}
