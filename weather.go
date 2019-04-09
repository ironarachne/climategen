package climategen

import (
	"math/rand"
)

// WeatherProfile is a summary of a type of weather
type WeatherProfile struct {
	CloudCover          int
	CloudFrequency      int
	PrecipitationAmount int
	PrecipitationTypes  []string
	WindLevel           int
}

func (climate Climate) generateWeatherProfileForSeason(season Season) WeatherProfile {
	var weatherProfile WeatherProfile

	temp := climate.getCurrentTemperature(season)
	humidity := climate.getCurrentHumidity(season)

	if temp < 4 {
		weatherProfile.PrecipitationTypes = []string{"snow"}
	} else if temp < 5 {
		weatherProfile.PrecipitationTypes = []string{"snow", "sleet"}
	} else if temp < 6 {
		weatherProfile.PrecipitationTypes = []string{"hail", "sleet", "rain"}
	} else {
		weatherProfile.PrecipitationTypes = []string{"rain"}
	}

	weatherProfile.PrecipitationAmount = humidity

	weatherProfile.WindLevel = rand.Intn(10) + 1
	weatherProfile.CloudCover = rand.Intn(10) + 1
	weatherProfile.CloudFrequency = rand.Intn(10) + 1

	return weatherProfile
}
