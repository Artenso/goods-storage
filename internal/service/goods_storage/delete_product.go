package goods_storage

import "context"

// DeleteProduct implements goods_storage.IGoodsStorageService.
func (s *Service) DeleteProduct(ctx context.Context, id int64) error {
	return s.goodsRepository.DeleteProduct(ctx, id)
}
