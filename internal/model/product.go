package model

import (
	"database/sql"
	"time"
)

type ProductInfo struct {
	Name        string
	Description string
}

type Product struct {
	ID        int64
	Info      ProductInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
