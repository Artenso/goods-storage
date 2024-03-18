package goods_pg

import (
	"context"
	"time"

	"github.com/Artenso/goods-storage/internal/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

const (
	table = "goods"

	idCol          = "id"
	nameCol        = "name"
	descriptionCol = "description"
	createdAtCol   = "created_at"
	updatedAtCol   = "updated_at"
)

// Repository ...
type Repository struct {
	dbConn *pgx.Conn
}

// New creates new repository object
func New(dbConn *pgx.Conn) *Repository {
	return &Repository{
		dbConn: dbConn,
	}
}

// AddProduct adds product
func (r *Repository) AddProduct(ctx context.Context, info *model.ProductInfo) (*model.Product, error) {
	builder := sq.Insert(table).
		Columns(nameCol, descriptionCol, createdAtCol).
		Values(info.Name, info.Description, time.Now().UTC()).
		Suffix("RETURNING id, created_at").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build sql query")
	}

	product := new(model.Product)

	if err = r.dbConn.QueryRow(ctx, query, args...).Scan(&product.ID, &product.CreatedAt); err != nil {
		return nil, errors.Wrap(err, "failed to insert product to db")
	}

	product.Info = *info

	return product, nil
}

// GetProduct gets product by id
func (r *Repository) GetProduct(ctx context.Context, id int64) (*model.Product, error) {
	builder := sq.Select("*").
		From(table).
		Where(sq.Eq{idCol: id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build sql query")
	}

	var products []*model.Product

	row, err := r.dbConn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	if err = pgxscan.ScanAll(&products, row); err != nil {
		return nil, err
	}

	return products[0], nil
}

// ListProduct gets products from offset to limit
func (r *Repository) ListProduct(ctx context.Context, limit, offset int64) ([]*model.Product, error) {
	builder := sq.Select("*").
		From(table).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		GroupBy(idCol).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "failed to build sql query")
	}

	rows, err := r.dbConn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var products []*model.Product

	if err = pgxscan.ScanAll(&products, rows); err != nil {
		return nil, err
	}

	return products, nil
}

// UpdateProduct updates product name or/and description
func (r *Repository) UpdateProduct(ctx context.Context, id int64, info *model.UpdateProductInfo) (*model.Product, error) {
	builder := sq.Update(table).
		Set(updatedAtCol, time.Now().UTC()).
		Where(sq.Eq{idCol: id}).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar)

	if info.Name.Valid {
		builder = builder.Set(nameCol, info.Name.String)
	}
	if info.Description.Valid {
		builder = builder.Set(descriptionCol, info.Description.String)
	}

	query, args, err := builder.ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "failed to build sql query")
	}

	var products []*model.Product

	row, err := r.dbConn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	if err = pgxscan.ScanAll(&products, row); err != nil {
		return nil, err
	}

	return products[0], nil
}

// DeleteProduct deletes product from database
func (r *Repository) DeleteProduct(ctx context.Context, id int64) error {
	builder := sq.Delete(table).
		Where(sq.Eq{idCol: id}).
		Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return errors.Wrap(err, "failed to build sql query")
	}

	var deletedID uint64

	err = r.dbConn.QueryRow(ctx, query, args...).Scan(&deletedID)
	if err != nil {
		return err
	}

	return nil
}
