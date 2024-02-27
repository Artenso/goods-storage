package goods

import (
	"context"

	"github.com/Artenso/goods-storage/internal/model"
)

// DeleteProduct implements goods_storage.IGoodsRepository.
func (r *Repository) DeleteProduct(_ context.Context, id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, product := range r.goods {
		if product.ID == id {
			r.goods = append(r.goods[:i], r.goods[i+1:]...)
			return nil
		}
	}

	return model.ErrProductNotFound
}
