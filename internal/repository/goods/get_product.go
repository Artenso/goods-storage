package goods

import (
	"context"

	"github.com/Artenso/goods-storage/internal/model"
)

// GetProduct implements goods_storage.IGoodsRepository.
func (r *Repository) GetProduct(_ context.Context, id int64) (*model.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, product := range r.goods {
		if product.ID == id {
			return product, nil
		}
	}

	return nil, model.ErrProductNotFound
}
