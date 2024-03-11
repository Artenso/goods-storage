package model

import (
	"database/sql"
	"time"
)

type ProductInfo struct {
	Name        string `db:"name"`
	Description string `db:"description"`
}

type UpdateProductInfo struct {
	Name        sql.NullString
	Description sql.NullString
}

type Product struct {
	ID        int64        `db:"id"`
	Info      ProductInfo  `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}
