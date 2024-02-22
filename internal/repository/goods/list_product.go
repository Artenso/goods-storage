package goods

import (
	"context"
	"fmt"

	"github.com/Artenso/goods-storage/internal/model"
)

// ListProduct implements goods_storage.IGoodsRepository.
func (r *Repository) ListProduct(_ context.Context, limit int64, offset int64) ([]*model.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if offset >= int64(len(r.goods)) {
		return nil, fmt.Errorf("offset out of range, %v products are available now", len(r.goods))
	}

	if offset+limit > int64(len(r.goods)) {
		return r.goods[offset:], nil
	}

	return r.goods[offset : offset+limit], nil

}
