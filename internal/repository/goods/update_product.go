package goods

import (
	"context"
	"time"

	"github.com/Artenso/goods-storage/internal/model"
)

// UpdateProduct implements goods_storage.IGoodsRepository.
func (r *Repository) UpdateProduct(_ context.Context, id int64, info *model.UpdateProductInfo) (*model.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, product := range r.goods {
		if product.ID == id {

			if info.Name.Valid {
				product.Info.Name = info.Name.String
			}

			if info.Description.Valid {
				product.Info.Description = info.Description.String
			}

			product.UpdatedAt.Time = time.Now().UTC()
			product.UpdatedAt.Valid = true
			return product, nil
		}
	}

	return nil, model.ErrProductNotFound
}
