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
					"blackberries",
					[]string{"fruit", "dye", "food", "plant"},
				},
				Resource{
					"blueberries",
					[]string{"fruit", "dye", "food", "plant"},
				},
				Resource{
					"raspberries",
					[]string{"fruit", "dye", "food", "plant"},
				},
				Resource{
					"chamois",
					[]string{"hide", "meat", "food"},
				},
				Resource{
					"goose",
					[]string{"feathers", "meat", "food"},
				},
				Resource{
					"softwood",
					[]string{"wood"},
				},
				Resource{
					"brahmi",
					[]string{"herb", "medicine", "plant"},
				},
				Resource{
					"nutmeg",
					[]string{"spice", "plant"},
				},
				Resource{
					"parsley",
					[]string{"spice", "herb", "plant"},
				},
				Resource{
					"potato",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"onion",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"carrot",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"eggplant",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"cabbage",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"kohlrabi",
					[]string{"vegetable", "plant", "food"},
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
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"squirrel",
					[]string{"meat", "food"},
				},
				Resource{
					"wolf",
					[]string{"hide"},
				},
				Resource{
					"basil",
					[]string{"spice", "herb", "plant"},
				},
				Resource{
					"cilantro",
					[]string{"spice", "herb", "plant"},
				},
				Resource{
					"potato",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"fiddlehead",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"broccoli",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"cauliflower",
					[]string{"vegetable", "plant", "food"},
				},
			},
		},
		"desert": Climate{
			Name:        "desert",
			Temperature: 9,
			Humidity:    0,
			Resources: []Resource{
				Resource{
					"camel",
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
					"scorpion",
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
				Resource{
					"saffron",
					[]string{"spice", "plant"},
				},
				Resource{
					"yam",
					[]string{"vegetable", "plant", "food"},
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
					"chicken",
					[]string{"feathers", "eggs", "meat", "food"},
				},
				Resource{
					"coyote",
					[]string{"hide"},
				},
				Resource{
					"flowers",
					[]string{"flowers", "plant"},
				},
				Resource{
					"horse",
					[]string{"mount", "meat", "food"},
				},
				Resource{
					"cow",
					[]string{"hide", "meat", "milk", "food"},
				},
				Resource{
					"parsley",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"rosemary",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"thyme",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"black pepper",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"potato",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"onion",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"tomato",
					[]string{"vegetable", "plant", "food"},
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
					"duck",
					[]string{"feathers", "eggs", "meat", "food"},
				},
				Resource{
					"flamingo",
					[]string{"feathers", "meat", "food"},
				},
				Resource{
					"snake",
					[]string{"hide", "meat", "food", "poison"},
				},
				Resource{
					"herbs",
					[]string{"herbs", "plant", "medicine"},
				},
				Resource{
					"reed",
					[]string{"plant", "basket"},
				},
				Resource{
					"vine",
					[]string{"plant", "basket"},
				},
				Resource{
					"oregano",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"mint",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"cabbage",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"rice",
					[]string{"grain", "plant", "food"},
				},
				Resource{
					"water chestnuts",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"bamboo shoots",
					[]string{"vegetable", "plant", "food"},
				},
			},
		},
		"tropical": Climate{
			Name:        "tropical",
			Temperature: 9,
			Humidity:    7,
			Resources: []Resource{
				Resource{
					"avocado",
					[]string{"fruit", "plant", "food"},
				},
				Resource{
					"banana",
					[]string{"fruit", "plant", "food"},
				},
				Resource{
					"coconut",
					[]string{"fruit", "plant", "food"},
				},
				Resource{
					"coffee",
					[]string{"drink", "plant", "food"},
				},
				Resource{
					"elephant",
					[]string{"hide", "meat", "food"},
				},
				Resource{
					"ivory",
					[]string{"ivory", "decoration"},
				},
				Resource{
					"jaguar",
					[]string{"hide"},
				},
				Resource{
					"mango",
					[]string{"fruit", "plant", "food"},
				},
				Resource{
					"pineapple",
					[]string{"fruit", "plant", "food"},
				},
				Resource{
					"tiger",
					[]string{"hide"},
				},
				Resource{
					"yam",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"lemongrass",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"cinnamon",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"coriander",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"cloves",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"basil",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"chili pepper",
					[]string{"spice", "plant"},
				},
				Resource{
					"bell pepper",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"lettuce",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"bok choy",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"leek",
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
					"diamond",
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
				Resource{
					"rosemary",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"thyme",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"sage",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"parsley",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"onion",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"brussel sprouts",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"green beans",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"pinto beans",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"black beans",
					[]string{"vegetable", "plant", "food"},
				},
			},
		},
		"rainforest": Climate{
			Name:        "rainforest",
			Temperature: 9,
			Humidity:    9,
			Resources: []Resource{
				Resource{
					"avocado",
					[]string{"fruit", "plant", "food"},
				},
				Resource{
					"banana",
					[]string{"fruit", "plant", "food"},
				},
				Resource{
					"coconut",
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
					"pineapple",
					[]string{"fruit", "plant", "food"},
				},
				Resource{
					"softwood",
					[]string{"wood"},
				},
				Resource{
					"yam",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"lemongrass",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"cinnamon",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"turmeric",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"vanilla",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"galangal",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"ginger",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"kaffir lime",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"chili pepper",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"yam",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"bok choy",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"leek",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"bell pepper",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"onion",
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
					"cheetah",
					[]string{"hide"},
				},
				Resource{
					"copper",
					[]string{"metal", "precious metal"},
				},
				Resource{
					"diamond",
					[]string{"gemstone"},
				},
				Resource{
					"elephant",
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
					"hippo",
					[]string{"hide", "meat", "food"},
				},
				Resource{
					"lion",
					[]string{"hide", "meat", "food"},
				},
				Resource{
					"nickel",
					[]string{"metal"},
				},
				Resource{
					"ruby",
					[]string{"gemstone"},
				},
				Resource{
					"sapphire",
					[]string{"gemstone"},
				},
				Resource{
					"silver",
					[]string{"metal", "precious metal"},
				},
				Resource{
					"zebra",
					[]string{"hide", "mount", "pack animal", "meat", "food"},
				},
				Resource{
					"salt",
					[]string{"spice"},
				},
				Resource{
					"yam",
					[]string{"vegetable", "plant", "food"},
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
					"donkey",
					[]string{"hide", "meat", "food", "pack animal", "mount"},
				},
				Resource{
					"flowers",
					[]string{"flowers", "decoration"},
				},
				Resource{
					"horse",
					[]string{"mount", "meat", "food", "pack animal"},
				},
				Resource{
					"rabbit",
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
					"wolf",
					[]string{"hide"},
				},
				Resource{
					"rosemary",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"thyme",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"potato",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"yam",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"onion",
					[]string{"vegetable", "plant", "food"},
				},
			},
		},
		"taiga": Climate{
			Name:        "taiga",
			Temperature: 3,
			Humidity:    3,
			Resources: []Resource{
				Resource{
					"ermine",
					[]string{"hide"},
				},
				Resource{
					"fox",
					[]string{"hide"},
				},
				Resource{
					"iron",
					[]string{"metal", "weapon metal", "armor metal"},
				},
				Resource{
					"mink",
					[]string{"hide"},
				},
				Resource{
					"mudhen",
					[]string{"feathers", "meat", "eggs", "food"},
				},
				Resource{
					"oil",
					[]string{"oil", "fuel"},
				},
				Resource{
					"rabbit",
					[]string{"hide", "meat", "food"},
				},
				Resource{
					"salmon",
					[]string{"fish", "meat", "food"},
				},
				Resource{
					"squirrel",
					[]string{"hide", "meat", "food"},
				},
				Resource{
					"trout",
					[]string{"fish", "meat", "food"},
				},
				Resource{
					"wolf",
					[]string{"hide"},
				},
				Resource{
					"sage",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"parsley",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"black pepper",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"salt",
					[]string{"spice"},
				},
				Resource{
					"potato",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"fiddlehead",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"cabbage",
					[]string{"vegetable", "plant", "food"},
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
					"falcon",
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
					"rabbit",
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
				Resource{
					"sage",
					[]string{"herb", "spice", "plant"},
				},
				Resource{
					"salt",
					[]string{"spice"},
				},
				Resource{
					"potato",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"sweet potato",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"carrot",
					[]string{"vegetable", "plant", "food"},
				},
				Resource{
					"yam",
					[]string{"vegetable", "plant", "food"},
				},
			},
		},
	}
)
