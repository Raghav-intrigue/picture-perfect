package viewmodel

import (
	"clumio.com/roger/picture-perfect/webapp/model"
)

// ProductList : product list page's view model
type ProductList struct {
	Title    string
	Active   string
	Products []Product
}

// NewProductList : returns the viewmodel, given the list of products
func NewProductList(products []model.Product) ProductList {
	result := ProductList{
		Title:    "Lemonade Stand Supply",
		Active:   "shop",
		Products: []Product{},
	}
	for _, p := range products {
		result.Products = append(result.Products, productToVM(&p))
	}
	return result
}
