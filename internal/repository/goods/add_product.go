package goods

import (
	"context"
	"time"

	"github.com/Artenso/goods-storage/internal/model"
)

// AddProduct implements goods_storage.IGoodsRepository.
func (r *Repository) AddProduct(_ context.Context, info *model.ProductInfo) (*model.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.lastId++

	id := r.lastId
	product := &model.Product{
		ID:        id,
		Info:      *info,
		CreatedAt: time.Now().UTC(),
	}

	r.goods = append(r.goods, product)

	return product, nil
}
