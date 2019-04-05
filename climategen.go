package climategen

import (
	"github.com/ironarachne/random"
)

// Animal is an animal
type Animal struct {
	Name         string
	PluralName   string
	AnimalType   string
	EatsAnimals  bool
	EatsPlants   bool
	GivesBone    bool
	GivesEggs    bool
	GivesHide    bool
	GivesHorn    bool
	GivesMeat    bool
	GivesMilk    bool
	IsMount      bool
	IsPackAnimal bool
	IsScavenger  bool
	IsVenomous   bool
}

// Climate is a climate
type Climate struct {
	Name        string
	Temperature int
	Humidity    int
	Seasons     map[string]Season
	Resources   []Resource
	Needs       []Resource
}

// Mineral is a mineral
type Mineral struct {
	Name         string
	PluralName   string
	Hardness     int
	HasOre       bool
	IsEdible     bool
	IsGem        bool
	IsMetal      bool
	IsPrecious   bool
	IsStone      bool
	Malleability int
}

// Plant is a plant
type Plant struct {
	Name        string
	PluralName  string
	HasWood     bool
	IsFiber     bool
	IsFruit     bool
	IsGrain     bool
	IsHerb      bool
	IsMedicine  bool
	IsNut       bool
	IsRoot      bool
	IsSpice     bool
	IsToxic     bool
	IsTree      bool
	IsVegetable bool
}

// Resource is a resource
type Resource struct {
	Name  string
	Types []string
}

// Season is a season
type Season struct {
	Name              string
	TemperatureChange int
	HumidityChange    int
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

func (climate Climate) getCurrentHumidity(season Season) int {
	return climate.Humidity + season.HumidityChange
}

func (climate Climate) getCurrentTemperature(season Season) int {
	return climate.Temperature + season.TemperatureChange
}

func getListOfNeeds(climate Climate) []Resource {
	resourcesOwned := climate.Resources
	resourcesNeeded := []Resource{}

	climateToCheck := Climate{}

	for _, otherClimate := range climates {
		climateToCheck = climateData[otherClimate]

		for _, resource := range climateToCheck.Resources {
			if !IsResourceInSlice(resource, resourcesOwned) {
				if !IsResourceInSlice(resource, resourcesNeeded) {
					resourcesNeeded = append(resourcesNeeded, resource)
				}
			}
		}
	}

	return resourcesNeeded
}
