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

// New ...
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

	res := new(model.Product)

	if err = r.dbConn.QueryRow(ctx, query, args...).Scan(&res.ID, &res.CreatedAt); err != nil {
		return nil, errors.Wrap(err, "failed to insert product to db")
	}

	res.Info = *info

	return res, nil
}

func (r *Repository) GetProduct(ctx context.Context, id int64) (*model.Product, error) {
	return nil, nil
}
func (r *Repository) ListProduct(ctx context.Context, limit, offset int64) ([]*model.Product, error) {
	builder := sq.Select("*").
		From(table).
		Limit(uint64(limit)).
		Offset(uint64(offset)).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "failed to build sql query")
	}

	rows, err := r.dbConn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	//res := make([]*model.Product, 0, limit-offset)
	//
	//for rows.Next() {
	//	product := new(model.Product)
	//	err = rows.Scan(&product.ID, &product.Info.Name, &product.Info.Description, &product.CreatedAt, &product.UpdatedAt)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	res = append(res, product)
	//}

	var res []*model.Product

	if err = pgxscan.ScanAll(&res, rows); err != nil {
		return nil, err
	}

	return res, nil
}
func (r *Repository) UpdateProduct(ctx context.Context, id int64, info *model.UpdateProductInfo) (*model.Product, error) {
	return nil, nil
}
func (r *Repository) DeleteProduct(ctx context.Context, id int64) error {
	return nil
}
