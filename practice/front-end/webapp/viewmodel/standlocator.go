package viewmodel

// StandLocator : represents the Stand Locator page view model
type StandLocator struct {
	Title  string
	Active string
}

// NewStandLocator : returns the Stand Locator page viewmodel
func NewStandLocator() StandLocator {
	result := StandLocator{
		Active: "standlocator",
		Title:  "Lemonade Stand Supply - Stand Locator",
	}
	return result
}
