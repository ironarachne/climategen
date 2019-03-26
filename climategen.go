package climategen

import (
	"github.com/ironarachne/random"
)

// Climate is a climate
type Climate struct {
	Name        string
	Temperature int
	Humidity    int
	Resources   []Resource
	Needs       []Resource
}

// Resource is a resource
type Resource struct {
	Name  string
	Types []string
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

func getListOfNeeds(climate Climate) []Resource {
	resourcesOwned := climate.Resources
	resourcesNeeded := []Resource{}

	climateToCheck := Climate{}

	for _, otherClimate := range climates {
		climateToCheck = climateData[otherClimate]

		for _, resource := range climateToCheck.Resources {
			if !isResourceInSlice(resource, resourcesOwned) {
				if !isResourceInSlice(resource, resourcesNeeded) {
					resourcesNeeded = append(resourcesNeeded, resource)
				}
			}
		}
	}

	return resourcesNeeded
}
