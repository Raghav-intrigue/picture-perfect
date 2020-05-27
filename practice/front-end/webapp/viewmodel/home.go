package viewmodel

// Home : home page view model
type Home struct {
	Title  string
	Active string
}

// NewHome : returns the Home page viewmodel
func NewHome() Home {
	result := Home{
		Active: "home",
		Title:  "Lemonade Stand Supply",
	}
	return result
}
