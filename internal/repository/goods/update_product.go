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

	for _, product := range r.goods {
		if product.ID == id {

			if info.Name.Valid {
				product.Info.Name = info.Name
			}

			if info.Description.Valid {
				product.Info.Description = info.Description
			}

			product.UpdatedAt.Time = time.Now().UTC()
			product.UpdatedAt.Valid = true
			return product, nil
		}
	}

	return nil, fmt.Errorf("no such product with id: %v", id)
}
