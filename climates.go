package climategen

import "math/rand"

// Climate is a climate
type Climate struct {
	Name           string
	Temperature    int
	Humidity       int
	Seasons        map[string]Season
	Animals        []Animal
	Plants         []Plant
	Stones         []Mineral
	Gems           []Mineral
	CommonMetals   []Mineral
	PreciousMetals []Mineral
}

func getClimateByName(name string) Climate {
	climates := getAllClimates()

	for _, c := range climates {
		if c.Name == name {
			return c
		}
	}

	panic("No such climate")
}

func getRandomClimate() Climate {
	climates := getAllClimates()

	return climates[rand.Intn(len(climates)-1)]
}

func (climate Climate) getCurrentHumidity(season Season) int {
	return climate.Humidity + season.HumidityChange
}

func (climate Climate) getCurrentTemperature(season Season) int {
	return climate.Temperature + season.TemperatureChange
}

func (climate Climate) populate() Climate {
	animals := climate.getFilteredAnimals()
	gems := getGems()
	commonMetals := getCommonMetals()
	preciousMetals := getPreciousMetals()
	plants := climate.getFilteredPlants()
	stones := getStones()

	climate.Seasons = climate.getSeasons()

	climate.Animals = getRandomAnimals(10, animals)
	climate.Gems = getRandomMinerals(3, gems)
	climate.CommonMetals = getRandomMinerals(2, commonMetals)
	climate.PreciousMetals = getRandomMinerals(3, preciousMetals)
	climate.Plants = getRandomPlants(12, plants)
	climate.Stones = getRandomMinerals(3, stones)

	return climate
}

func getAllClimates() []Climate {
	climates := []Climate{
		Climate{
			Name:        "coniferous forest",
			Temperature: 4,
			Humidity:    6,
		},
		Climate{
			Name:        "deciduous forest",
			Temperature: 5,
			Humidity:    5,
		},
		Climate{
			Name:        "desert",
			Temperature: 9,
			Humidity:    0,
		},
		Climate{
			Name:        "grassland",
			Temperature: 5,
			Humidity:    3,
		},
		Climate{
			Name:        "marshland",
			Temperature: 7,
			Humidity:    9,
		},
		Climate{
			Name:        "tropical",
			Temperature: 9,
			Humidity:    7,
		},
		Climate{
			Name:        "mountain",
			Temperature: 4,
			Humidity:    4,
		},
		Climate{
			Name:        "rainforest",
			Temperature: 9,
			Humidity:    9,
		},
		Climate{
			Name:        "savanna",
			Temperature: 9,
			Humidity:    5,
		},
		Climate{
			Name:        "steppe",
			Temperature: 7,
			Humidity:    3,
		},
		Climate{
			Name:        "taiga",
			Temperature: 3,
			Humidity:    3,
		},
		Climate{
			Name:        "tundra",
			Temperature: 1,
			Humidity:    3,
		},
	}

	return climates
}
