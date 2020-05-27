package model

// StandCoordinate : coordinates for a Stand
type StandCoordinate struct {
	Title     string  `json:"title"`
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lng"`
}

// GetStandCoords : returns Stand Coordinates
func GetStandCoords() []StandCoordinate {
	return coords
}

var coords []StandCoordinate = []StandCoordinate{
	{
		Latitude:  37.409,
		Longitude: -122.06,
		Title:     "Bobby's stand",
	},
	{
		Latitude:  37.4092,
		Longitude: -122.061,
		Title:     "Macy's stand",
	},
	{
		Latitude:  37.4094,
		Longitude: -122.06,
		Title:     "Juan's stand",
	},
	{
		Latitude:  37.41,
		Longitude: -122.065,
		Title:     "Allison's stand",
	},
	{
		Latitude:  37.415,
		Longitude: -122.07,
		Title:     "Chen's stand",
	},
	{
		Latitude:  37.4217,
		Longitude: -122.075,
		Title:     "Matthew's stand",
	},
	{
		Latitude:  37.4206,
		Longitude: -122.08,
		Title:     "Alice's stand",
	},
	{
		Latitude:  37.4205,
		Longitude: -122.083,
		Title:     "Kara's stand",
	},
	{
		Latitude:  37.414,
		Longitude: -122.09,
		Title:     "Fred's stand",
	},
	{
		Latitude:  37.412,
		Longitude: -122.09,
		Title:     "Jake's stand",
	},
	{
		Latitude:  37.41,
		Longitude: -122.093,
		Title:     "Wallace's stand",
	},
	{
		Latitude:  37.416,
		Longitude: -122.095,
		Title:     "Gromit's stand",
	},
	{
		Latitude:  37.41,
		Longitude: -122.1,
		Title:     "Kirk's stand",
	},
	{
		Latitude:  37.41,
		Longitude: -122.102,
		Title:     "Lorelei's stand",
	},
	{
		Latitude:  37.412,
		Longitude: -122.099,
		Title:     "Rebecca's stand",
	},
	{
		Latitude:  37.407,
		Longitude: -122.1025,
		Title:     "Chris's stand",
	},
	{
		Latitude:  37.423,
		Longitude: -122.1025,
		Title:     "Carson's stand",
	},
}
