package goods

import (
	"context"
	"fmt"

	"github.com/Artenso/goods-storage/internal/model"
)

// GetProduct implements goods_storage.IGoodsRepository.
func (r *Repository) GetProduct(_ context.Context, id int64) (*model.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for i := 0; i < len(r.goods); i++ {
		if r.goods[i].ID == id {
			return r.goods[i], nil
		}
	}

	return nil, fmt.Errorf("no such product with id: %v", id)
}
