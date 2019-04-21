package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/Toshik1978/cart_to_qwintry/cartparsers"
	"github.com/pkg/errors"
)

const (
	templateSheet = "Лист1"
	templateCell  = "%s%d"
	cartXlsx      = "./cart.xlsx"
)

// SaveTemplate saves template with cart items
func SaveTemplate(templatePath string, cart []cartparsers.CartItem) error {
	if len(cart) == 0 {
		return errors.New("empty cart detected")
	}

	xlsx, err := excelize.OpenFile(templatePath)
	if err != nil {
		return errors.Wrap(err, "failed to open template")
	}

	for i, cartItem := range cart {
		writeCartItem(xlsx, i+2, cartItem)
	}

	if err := xlsx.SaveAs(cartXlsx); err != nil {
		return errors.Wrap(err, "failed to save XLSX")
	}
	return nil
}

// writeCartItem save cart's item in XLSX
func writeCartItem(xlsx *excelize.File, index int, cartItem cartparsers.CartItem) {
	xlsx.SetCellValue(templateSheet, fmt.Sprintf(templateCell, "A", index), getProductDesc(cartItem))
	xlsx.SetCellValue(templateSheet, fmt.Sprintf(templateCell, "C", index), cartItem.ProductQty)
	xlsx.SetCellValue(templateSheet, fmt.Sprintf(templateCell, "D", index), cartItem.ProductPrice)
	xlsx.SetCellValue(templateSheet, fmt.Sprintf(templateCell, "F", index), cartItem.ProductUrl)
}

// getProductDesc retrieve description for template
func getProductDesc(cartItem cartparsers.CartItem) string {
	desc := cartItem.ProductName
	if len(cartItem.ProductStyle) > 0 {
		desc += ", " + cartItem.ProductStyle
	}
	if len(cartItem.ProductSize) > 0 {
		desc += ", " + cartItem.ProductSize
	}
	if len(cartItem.ProductColor) > 0 {
		desc += ", " + cartItem.ProductColor
	}
	return desc
}
