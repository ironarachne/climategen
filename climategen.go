package climategen

import (
	"github.com/ironarachne/utility"
)

// Climate is a climate
type Climate struct {
	Name        string
	Temperature int
	Humidity    int
	Resources   []string
}

// Generate generates a climate
func Generate() Climate {
	climateOption := utility.RandomItem(climates)

	climate := climateData[climateOption]

	return climate
}

// GetClimate returns a specific climate
func GetClimate(name string) Climate {
	climate := climateData[name]

	return climate
}
