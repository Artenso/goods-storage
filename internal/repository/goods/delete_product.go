package goods

import (
	"context"
	"fmt"
)

// DeleteProduct implements goods_storage.IGoodsRepository.
func (r *Repository) DeleteProduct(_ context.Context, id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i := 0; i < len(r.goods); i++ {
		if r.goods[i].ID == id {
			r.goods = append(r.goods[:i], r.goods[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("no such product with id: %v", id)
}
