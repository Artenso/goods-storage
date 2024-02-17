package goods

import (
	"context"
	"sync"

	"github.com/Artenso/goods-storage/internal/model"
)

type Repository struct {
	mu    sync.RWMutex
	goods []*model.Product
}

func New() *Repository {
	return &Repository{
		mu:    sync.RWMutex{},
		goods: make([]*model.Product, 0, 10),
	}
}

func (r *Repository) AddProduct(_ context.Context, info *model.ProductInfo) (*model.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := int64(len(r.goods)) + 1
	product := &model.Product{
		ID:   id,
		Info: *info,
	}

	r.goods = append(r.goods, product)

	return product, nil
}
