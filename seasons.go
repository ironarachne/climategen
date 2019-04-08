package climategen

// Season is a season
type Season struct {
	Name              string
	TemperatureChange int
	HumidityChange    int
}

func (climate Climate) getSeasons() map[string]Season {
	fourSeasons := map[string]Season{
		"spring": Season{
			Name:              "spring",
			HumidityChange:    0,
			TemperatureChange: 0,
		},
		"summer": Season{
			Name:              "summer",
			HumidityChange:    1,
			TemperatureChange: 2,
		},
		"autumn": Season{
			Name:              "autumn",
			HumidityChange:    -1,
			TemperatureChange: -1,
		},
		"winter": Season{
			Name:              "winter",
			HumidityChange:    -2,
			TemperatureChange: -2,
		},
	}

	twoSeasons := map[string]Season{
		"dry": Season{
			Name:              "dry",
			HumidityChange:    -1,
			TemperatureChange: +1,
		},
		"wet": Season{
			Name:              "wet",
			HumidityChange:    +1,
			TemperatureChange: 0,
		},
	}

	if climate.Humidity > 5 && climate.Temperature > 6 {
		return twoSeasons
	}

	return fourSeasons
}
