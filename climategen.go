package climategen

// Generate generates a climate
func Generate() Climate {
	climate := getRandomClimate()
	climate = climate.populate()

	return climate
}

// GetClimate returns a specific climate
func GetClimate(name string) Climate {
	climate := getClimateByName(name)
	climate = climate.populate()

	return climate
}
