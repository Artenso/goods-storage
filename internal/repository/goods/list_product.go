package goods

import (
	"context"

	"github.com/Artenso/goods-storage/internal/model"
)

// ListProduct implements goods_storage.IGoodsRepository.
func (r *Repository) ListProduct(_ context.Context, limit int64, offset int64) ([]*model.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if offset >= int64(len(r.goods)) {
		return r.goods[len(r.goods):], nil
	}

	end := offset + limit
	if end > int64(len(r.goods)) {
		return r.goods[offset:], nil
	}

	return r.goods[offset:end], nil

}
