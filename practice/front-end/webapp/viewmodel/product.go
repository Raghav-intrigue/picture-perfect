package viewmodel

import (
	"html/template"

	"clumio.com/roger/picture-perfect/webapp/model"
)

// ProductViewModel : represents the product page view model
type ProductViewModel struct {
	Title   string
	Active  string
	Product Product
}

// Product : represents each product
type Product struct {
	Name             string
	DescriptionShort string
	DescriptionLong  template.HTML
	PricePerLiter    float32
	PricePer10Liter  float32
	Origin           string
	IsOrganic        bool
	ImageURL         string
	ID               int
}

// NewProduct : returns Product page viewmodel, given the Product
func NewProduct(product *model.Product) ProductViewModel {
	return ProductViewModel{
		Title:   "Lemonade Stand Supply - Shop",
		Active:  "shop",
		Product: productToVM(product),
	}
}

// productToVM : converts the raw model data to a view model
func productToVM(product *model.Product) Product {
	return Product{
		Name:             product.Name,
		DescriptionShort: product.DescriptionShort,
		DescriptionLong:  template.HTML(product.DescriptionLong),
		PricePerLiter:    product.PricePerLiter,
		PricePer10Liter:  product.PricePer10Liter,
		Origin:           product.Origin,
		IsOrganic:        product.IsOrganic,
		ImageURL:         product.ImageURL,
		ID:               product.ID,
	}
}
