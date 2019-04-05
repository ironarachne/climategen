package climategen

func getAllMinerals() []Mineral {
	var minerals []Mineral

	gems := getGems()
	metal := getMetal()
	other := getOtherMinerals()
	stone := getStone()

	minerals = append(minerals, gems...)
	minerals = append(minerals, metal...)
	minerals = append(minerals, other...)
	minerals = append(minerals, stone...)

	return minerals
}

func getGems() []Mineral {
	diamond := Mineral{
		Name:         "diamond",
		PluralName:   "diamonds",
		Hardness:     10,
		HasOre:       true,
		IsEdible:     false,
		IsGem:        true,
		IsMetal:      false,
		IsPrecious:   true,
		IsStone:      false,
		Malleability: 1,
	}

	garnet := Mineral{
		Name:         "garnet",
		PluralName:   "garnets",
		Hardness:     6,
		HasOre:       true,
		IsEdible:     false,
		IsGem:        true,
		IsMetal:      false,
		IsPrecious:   true,
		IsStone:      false,
		Malleability: 3,
	}

	ruby := Mineral{
		Name:         "ruby",
		PluralName:   "rubies",
		Hardness:     6,
		HasOre:       true,
		IsEdible:     false,
		IsGem:        true,
		IsMetal:      false,
		IsPrecious:   true,
		IsStone:      false,
		Malleability: 3,
	}

	sapphire := Mineral{
		Name:         "sapphire",
		PluralName:   "sapphires",
		Hardness:     6,
		HasOre:       true,
		IsEdible:     false,
		IsGem:        true,
		IsMetal:      false,
		IsPrecious:   true,
		IsStone:      false,
		Malleability: 3,
	}

	return []Mineral{diamond, garnet, ruby, sapphire}
}

func getMetal() []Mineral {
	copper := Mineral{
		Name:         "copper",
		PluralName:   "copper",
		Hardness:     5,
		HasOre:       true,
		IsEdible:     false,
		IsGem:        false,
		IsMetal:      true,
		IsPrecious:   true,
		IsStone:      false,
		Malleability: 5,
	}

	gold := Mineral{
		Name:         "gold",
		PluralName:   "gold",
		Hardness:     1,
		HasOre:       true,
		IsEdible:     true,
		IsGem:        false,
		IsMetal:      true,
		IsPrecious:   true,
		IsStone:      false,
		Malleability: 8,
	}

	iron := Mineral{
		Name:         "iron",
		PluralName:   "iron",
		Hardness:     8,
		HasOre:       true,
		IsEdible:     false,
		IsGem:        false,
		IsMetal:      true,
		IsPrecious:   false,
		IsStone:      false,
		Malleability: 4,
	}

	nickel := Mineral{
		Name:         "nickel",
		PluralName:   "nickel",
		Hardness:     3,
		HasOre:       true,
		IsEdible:     false,
		IsGem:        false,
		IsMetal:      true,
		IsPrecious:   false,
		IsStone:      false,
		Malleability: 4,
	}

	silver := Mineral{
		Name:         "silver",
		PluralName:   "silver",
		Hardness:     1,
		HasOre:       true,
		IsEdible:     false,
		IsGem:        false,
		IsMetal:      true,
		IsPrecious:   true,
		IsStone:      false,
		Malleability: 7,
	}

	return []Mineral{copper, gold, iron, nickel, silver}
}

func getOtherMinerals() []Mineral {
	coal := Mineral{
		Name:         "coal",
		PluralName:   "coal",
		Hardness:     1,
		HasOre:       true,
		IsEdible:     false,
		IsGem:        false,
		IsMetal:      false,
		IsPrecious:   false,
		IsStone:      false,
		Malleability: 4,
	}

	salt := Mineral{
		Name:         "salt",
		PluralName:   "salt",
		Hardness:     1,
		HasOre:       false,
		IsEdible:     true,
		IsGem:        false,
		IsMetal:      false,
		IsPrecious:   false,
		IsStone:      false,
		Malleability: 1,
	}

	return []Mineral{coal, salt}
}

func getStone() []Mineral {
	granite := Mineral{
		Name:         "granite",
		PluralName:   "granite",
		Hardness:     6,
		HasOre:       false,
		IsEdible:     false,
		IsGem:        false,
		IsMetal:      false,
		IsPrecious:   false,
		IsStone:      true,
		Malleability: 2,
	}

	limestone := Mineral{
		Name:         "limestone",
		PluralName:   "limestone",
		Hardness:     4,
		HasOre:       false,
		IsEdible:     false,
		IsGem:        false,
		IsMetal:      false,
		IsPrecious:   false,
		IsStone:      true,
		Malleability: 4,
	}

	marble := Mineral{
		Name:         "marble",
		PluralName:   "marble",
		Hardness:     5,
		HasOre:       false,
		IsEdible:     false,
		IsGem:        false,
		IsMetal:      false,
		IsPrecious:   false,
		IsStone:      true,
		Malleability: 2,
	}

	sandstone := Mineral{
		Name:         "sandstone",
		PluralName:   "sandstone",
		Hardness:     4,
		HasOre:       false,
		IsEdible:     false,
		IsGem:        false,
		IsMetal:      false,
		IsPrecious:   false,
		IsStone:      true,
		Malleability: 4,
	}

	return []Mineral{granite, limestone, marble, sandstone}
}
