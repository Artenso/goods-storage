package goods

import (
	"context"
	"fmt"
	"time"

	"github.com/Artenso/goods-storage/internal/model"
)

// UpdateProduct implements goods_storage.IGoodsRepository.
func (r *Repository) UpdateProduct(_ context.Context, id int64, info *model.ProductInfo) (*model.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i := 0; i < len(r.goods); i++ {
		if r.goods[i].ID == id {

			if info.Name != "" {
				r.goods[i].Info.Name = info.Name
			}

			if info.Description != "" {
				r.goods[i].Info.Description = info.Description
			}

			r.goods[i].UpdatedAt.Time = time.Now().UTC()
			r.goods[i].UpdatedAt.Valid = true
			return r.goods[i], nil
		}
	}

	return nil, fmt.Errorf("no such product with id: %v", id)
}
