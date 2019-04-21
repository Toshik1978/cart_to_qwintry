package cartparsers

import "io"

// CartParser declare interface for parsing cart
type CartParser interface {
	Parse(reader io.Reader) ([]CartItem, error)
}

// CartItem declare cart item
type CartItem struct {
	ProductUrl   string
	ProductName  string
	ProductStyle string
	ProductSize  string
	ProductColor string
	ProductQty   int
	ProductPrice float64
}
