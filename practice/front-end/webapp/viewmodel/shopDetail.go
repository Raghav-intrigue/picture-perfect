package viewmodel

import (
	"clumio.com/roger/picture-perfect/webapp/model"
)

// ShopDetail : shop details page's view model
type ShopDetail struct {
	Title    string
	Active   string
	Products []Product
}

// NewShopDetail : returns a ShopDetails viewmodel, given the list of products
func NewShopDetail(products []model.Product) ShopDetail {
	result := ShopDetail{
		Title:    "Lemonade Stand Supply",
		Active:   "shop",
		Products: []Product{},
	}
	for _, p := range products {
		result.Products = append(result.Products, productToVM(&p))
	}
	return result
}
