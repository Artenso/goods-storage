package model

type ProductInfo struct {
	Name        string
	Description string
}

type Product struct {
	ID   int64
	Info ProductInfo
}
