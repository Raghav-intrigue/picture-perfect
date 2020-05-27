package viewmodel

import (
	"fmt"

	"clumio.com/roger/picture-perfect/webapp/model"
)

// Shop : represents the shop page view model
type Shop struct {
	Title      string
	Active     string
	Categories []Category
}

// Category : represents the category sections of of the shop page
type Category struct {
	URL           string
	ImageURL      string
	Title         string
	Description   string
	IsOrientRight bool
}

// NewShop : creates a Shop page viewmodel, given the list of categories
func NewShop(categories []model.Category) Shop {
	result := Shop{
		Title:  "Lemonade Stand Supply - Shop",
		Active: "shop",
	}
	result.Categories = make([]Category, len(categories))
	for i := 0; i < len(categories); i++ {
		vm := categoryToVM(categories[i])
		vm.IsOrientRight = i%2 == 1
		result.Categories[i] = vm
	}
	return result
}

// categoryToVM : converts the raw model data to a view model
func categoryToVM(c model.Category) Category {
	return Category{
		URL:         fmt.Sprintf("/shop/%v", c.ID),
		ImageURL:    c.ImageURL,
		Title:       c.Title,
		Description: c.Description,
	}
}
