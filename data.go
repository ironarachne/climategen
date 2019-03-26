package climategen

var (
	climates = []string{
		"coniferous",
		"deciduous",
		"desert",
		"grassland",
		"marshland",
		"tropical",
		"mountain",
		"rainforest",
		"savanna",
		"steppe",
		"taiga",
		"tundra",
	}

	climateData = map[string]Climate{
		"coniferous": Climate{
			Name:        "coniferous forest",
			Temperature: 4,
			Humidity:    6,
			Resources: []Resource{
				Resource{
					"bear",
					[]string{"hide", "meat", "food"},
				},
				Resource{
					"beaver",
					[]string{"hide", "meat", "food"},
				},
				Resource{
					"berries",
					[]string{"fruit", "dye", "food", "plant"},
				},
				Resource{
					"chamois",
					[]string{"textile"},
				},
				Resource{
					"geese",
					[]string{"feathers", "meat", "food"},
				},
				Resource{
					"softwood",
					[]string{"wood"},
				},
			},
		},
		"deciduous": Climate{
			Name:        "deciduous forest",
			Temperature: 5,
			Humidity:    5,
			Resources: []Resource{
				Resource{
					"berries",
					[]string{"fruit", "dye", "food", "plant"},
				},
				Resource{
					"deer",
					[]string{"hide", "meat", "food"},
				},
				Resource{
					"eagle",
					[]string{"feathers"},
				},
				Resource{
					"hardwood",
					[]string{"wood"},
				},
				Resource{
					"herbs",
					[]string{"herbs", "plant", "medicine"},
				},
				Resource{
					"mushrooms",
					[]string{"plant", "food"},
				},
				Resource{
					"squirrel",
					[]string{"meat", "food"},
				},
				Resource{
					"wolves",
					[]string{"hide"},
				},
			},
		},
		"desert": Climate{
			Name:        "desert",
			Temperature: 9,
			Humidity:    0,
			Resources: []Resource{
				Resource{
					"camels",
					[]string{"pack animal", "mount", "meat", "food"},
				},
				Resource{
					"copper",
					[]string{"metal", "precious metal"},
				},
				Resource{
					"glass",
					[]string{"glass"},
				},
				Resource{
					"gold",
					[]string{"metal", "precious metal"},
				},
				Resource{
					"iron",
					[]string{"metal", "weapon metal", "armor metal"},
				},
				Resource{
					"oil",
					[]string{"oil"},
				},
				Resource{
					"salt",
					[]string{"spice"},
				},
				Resource{
					"scorpions",
					[]string{"meat", "food", "poison"},
				},
				Resource{
					"silver",
					[]string{"metal", "precious metal"},
				},
				Resource{
					"turquoise",
					[]string{"gemstone"},
				},
			},
		},
		"grassland": Climate{
			Name:        "grassland",
			Temperature: 5,
			Humidity:    3,
			Resources: []Resource{
				Resource{
					"bison",
					[]string{"hide", "meat", "food"},
				},
				Resource{
					"bobcat",
					[]string{"hide"},
				},
				Resource{
					"chickens",
					[]string{"feathers", "eggs", "meat", "food"},
				},
				Resource{
					"coyotes",
					[]string{"hide"},
				},
				Resource{
					"flowers",
					[]string{"flowers", "plant"},
				},
				Resource{
					"horses",
					[]string{"mount", "meat", "food"},
				},
				Resource{
					"cattle",
					[]string{"hide", "meat", "milk", "food"},
				},
			},
		},
		"marshland": Climate{
			Name:        "marshland",
			Temperature: 7,
			Humidity:    9,
			Resources: []Resource{
				Resource{
					"alligator",
					[]string{"hide", "meat", "food"},
				},
				Resource{
					"ducks",
					[]string{"feathers", "eggs", "meat", "food"},
				},
				Resource{
					"flamingos",
					[]string{"feathers", "meat", "food"},
				},
				Resource{
					"snakes",
					[]string{"hide", "meat", "food", "poison"},
				},
				Resource{
					"herbs",
					[]string{"herbs", "plant", "medicine"},
				},
				Resource{
					"reeds",
					[]string{"plant", "basket"},
				},
				Resource{
					"vines",
					[]string{"plant", "basket"},
				},
			},
		},
		"tropical": Climate{
			Name:        "tropical",
			Temperature: 9,
			Humidity:    7,
			Resources: []Resource{
				Resource{
					"avocados",
					[]string{"fruit", "plant", "food"},
				},
				Resource{
					"bananas",
					[]string{"fruit", "plant", "food"},
				},
				Resource{
					"coconuts",
					[]string{"fruit", "plant", "food"},
				},
				Resource{
					"coffee",
					[]string{"drink", "plant", "food"},
				},
				Resource{
					"elephants",
					[]string{"hide", "meat", "food"},
				},
				Resource{
					"ivory",
					[]string{"ivory", "decoration"},
				},
				Resource{
					"jaguars",
					[]string{"hide"},
				},
				Resource{
					"mangos",
					[]string{"fruit", "plant", "food"},
				},
				Resource{
					"pineapples",
					[]string{"fruit", "plant", "food"},
				},
				Resource{
					"tigers",
					[]string{"hide"},
				},
				Resource{
					"yams",
					[]string{"vegetable", "plant", "food"},
				},
			},
		},
		"mountain": Climate{
			Name:        "mountain",
			Temperature: 4,
			Humidity:    4,
			Resources: []Resource{
				Resource{
					"coal",
					[]string{"fuel", "dye"},
				},
				Resource{
					"copper",
					[]string{"metal", "precious metal"},
				},
				Resource{
					"diamonds",
					[]string{"gemstone"},
				},
				Resource{
					"gold",
					[]string{"metal", "precious metal"},
				},
				Resource{
					"granite",
					[]string{"stone"},
				},
				Resource{
					"iron",
					[]string{"metal", "weapon metal"},
				},
				Resource{
					"limestone",
					[]string{"stone"},
				},
				Resource{
					"marble",
					[]string{"stone"},
				},
				Resource{
					"sandstone",
					[]string{"stone"},
				},
				Resource{
					"silver",
					[]string{"metal", "precious metal"},
				},
			},
		},
		"rainforest": Climate{
			Name:        "rainforest",
			Temperature: 9,
			Humidity:    9,
			Resources: []Resource{
				Resource{
					"avocados",
					[]string{"fruit", "plant", "food"},
				},
				Resource{
					"bananas",
					[]string{"fruit", "plant", "food"},
				},
				Resource{
					"coconuts",
					[]string{"fruit", "plant", "food"},
				},
				Resource{
					"flowers",
					[]string{"flowers", "decoration", "plant"},
				},
				Resource{
					"hardwood",
					[]string{"wood"},
				},
				Resource{
					"herbs",
					[]string{"herbs", "plant", "medicine"},
				},
				Resource{
					"mushrooms",
					[]string{"plant", "food", "medicine", "poison"},
				},
				Resource{
					"pineapples",
					[]string{"fruit", "plant", "food"},
				},
				Resource{
					"softwood",
					[]string{"wood"},
				},
				Resource{
					"yams",
					[]string{"vegetable", "plant", "food"},
				},
			},
		},
		"savanna": Climate{
			Name:        "savanna",
			Temperature: 9,
			Humidity:    5,
			Resources: []Resource{
				Resource{
					"cheetahs",
					[]string{"hide"},
				},
				Resource{
					"copper",
					[]string{"metal", "precious metal"},
				},
				Resource{
					"diamonds",
					[]string{"gemstone"},
				},
				Resource{
					"elephants",
					[]string{"hide", "meat", "food", "mount", "pack animal"},
				},
				Resource{
					"ivory",
					[]string{"ivory", "decoration"},
				},
				Resource{
					"gazelles",
					[]string{"hide", "meat", "food"},
				},
				Resource{
					"gold",
					[]string{"metal", "precious metal"},
				},
				Resource{
					"hippos",
					[]string{"hide"},
				},
				Resource{
					"lions",
					[]string{"hide"},
				},
				Resource{
					"nickel",
					[]string{"metal"},
				},
				Resource{
					"rubies",
					[]string{"gemstone"},
				},
				Resource{
					"sapphires",
					[]string{"gemstone"},
				},
				Resource{
					"silver",
					[]string{"metal", "precious metal"},
				},
				Resource{
					"zebras",
					[]string{"hide", "mount", "pack animal", "meat", "food"},
				},
			},
		},
		"steppe": Climate{
			Name:        "steppe",
			Temperature: 7,
			Humidity:    3,
			Resources: []Resource{
				Resource{
					"antelope",
					[]string{"hide", "meat", "food"},
				},
				Resource{
					"donkeys",
					[]string{"hide", "meat", "food", "pack animal", "mount"},
				},
				Resource{
					"flowers",
					[]string{"flowers", "decoration"},
				},
				Resource{
					"horses",
					[]string{"mount", "meat", "food", "pack animal"},
				},
				Resource{
					"rabbits",
					[]string{"hide", "meat", "food"},
				},
				Resource{
					"sheep",
					[]string{"wool", "textiles", "meat", "food"},
				},
				Resource{
					"tea",
					[]string{"drink", "plant"},
				},
				Resource{
					"wolves",
					[]string{"hide"},
				},
			},
		},
		"taiga": Climate{
			Name:        "taiga",
			Temperature: 3,
			Humidity:    3,
			Resources: []Resource{
				Resource{
					"ermines",
					[]string{"hide"},
				},
				Resource{
					"foxes",
					[]string{"hide"},
				},
				Resource{
					"iron",
					[]string{"metal", "weapon metal", "armor metal"},
				},
				Resource{
					"minks",
					[]string{"hide"},
				},
				Resource{
					"oil",
					[]string{"oil", "fuel"},
				},
				Resource{
					"rabbits",
					[]string{"hide", "meat", "food"},
				},
				Resource{
					"salmon",
					[]string{"fish", "meat", "food"},
				},
				Resource{
					"squirrels",
					[]string{"hide", "meat", "food"},
				},
				Resource{
					"trout",
					[]string{"fish", "meat", "food"},
				},
				Resource{
					"wolves",
					[]string{"hide"},
				},
			},
		},
		"tundra": Climate{
			Name:        "tundra",
			Temperature: 1,
			Humidity:    3,
			Resources: []Resource{
				Resource{
					"coal",
					[]string{"fuel", "dye"},
				},
				Resource{
					"cod",
					[]string{"fish", "meat", "food"},
				},
				Resource{
					"copper",
					[]string{"metal", "precious metal"},
				},
				Resource{
					"falcons",
					[]string{"feathers", "hunting animals"},
				},
				Resource{
					"iron",
					[]string{"metal", "weapon metal", "armor metal"},
				},
				Resource{
					"nickel",
					[]string{"metal"},
				},
				Resource{
					"oil",
					[]string{"oil", "fuel"},
				},
				Resource{
					"rabbits",
					[]string{"hide", "meat", "food"},
				},
				Resource{
					"reindeer",
					[]string{"hide", "meat", "food", "mount", "pack animal"},
				},
				Resource{
					"trout",
					[]string{"fish", "meat", "food"},
				},
			},
		},
	}
)
