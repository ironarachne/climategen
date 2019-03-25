package climategen

import (
	"github.com/ironarachne/random"
	"github.com/ironarachne/utility"
)

// Climate is a climate
type Climate struct {
	Name        string
	Temperature int
	Humidity    int
	Resources   []string
	Needs       []string
}

// Generate generates a climate
func Generate() Climate {
	climateOption := random.Item(climates)

	climate := climateData[climateOption]
	climate.Needs = getListOfNeeds(climate)

	return climate
}

// GetClimate returns a specific climate
func GetClimate(name string) Climate {
	climate := climateData[name]
	climate.Needs = getListOfNeeds(climate)

	return climate
}

func getListOfNeeds(climate Climate) []string {
	resourcesOwned := climate.Resources
	resourcesNeeded := []string{}

	climateToCheck := Climate{}

	for _, otherClimate := range climates {
		climateToCheck = climateData[otherClimate]

		for _, resource := range climateToCheck.Resources {
			if !utility.ItemInCollection(resource, resourcesOwned) {
				if !utility.ItemInCollection(resource, resourcesNeeded) {
					resourcesNeeded = append(resourcesNeeded, resource)
				}
			}
		}
	}

	return resourcesNeeded
}
