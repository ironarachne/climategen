package climategen

// Resource is a refined resource derived from raw materials
type Resource struct {
	Name   string
	Origin string
	Type   string
}

func resourcesFromAnimal(source Animal) []Resource {
	resources := []Resource{}

	if source.GivesBone {
		resources = append(resources, Resource{Name: source.Name + " bone", Origin: source.Name, Type: "bone"})
	}
	if source.GivesEggs {
		resources = append(resources, Resource{Name: source.Name + " eggs", Origin: source.Name, Type: "eggs"})
	}
	if source.GivesHide {
		resources = append(resources, Resource{Name: source.Name + " hide", Origin: source.Name, Type: "hide"})
	}
	if source.GivesHorn {
		resources = append(resources, Resource{Name: source.Name + " horn", Origin: source.Name, Type: "horn"})
	}
	if source.GivesMeat {
		resources = append(resources, Resource{Name: source.Name + " meat", Origin: source.Name, Type: "meat"})
	}
	if source.GivesMilk {
		resources = append(resources, Resource{Name: source.Name + " milk", Origin: source.Name, Type: "milk"})
	}
	if source.IsMount {
		resources = append(resources, Resource{Name: "riding " + source.Name, Origin: source.Name, Type: "mount"})
	}
	if source.IsPackAnimal {
		resources = append(resources, Resource{Name: "pack " + source.Name, Origin: source.Name, Type: "pack animal"})
	}
	if source.IsVenomous {
		resources = append(resources, Resource{Name: source.Name + " venom", Origin: source.Name, Type: "toxin"})
	}

	return resources
}

func resourcesFromMineral(source Mineral) []Resource {
	resources := []Resource{}

	if source.HasOre {
		resources = append(resources, Resource{Name: source.Name + " ore", Origin: source.Name, Type: "ore"})
	}
	if source.IsGem {
		resources = append(resources, Resource{Name: source.PluralName, Origin: source.Name, Type: "gem"})
	}
	if source.IsMetal {
		resources = append(resources, Resource{Name: source.Name + " bar", Origin: source.Name, Type: "metal bar"})
	}
	if source.IsPrecious && source.IsMetal {
		resources = append(resources, Resource{Name: source.Name + " ingot", Origin: source.Name, Type: "metal ingot"})
	}
	if source.IsStone {
		resources = append(resources, Resource{Name: source.Name + " block", Origin: source.Name, Type: "stone block"})
	}
	if source.IsEdible {
		resources = append(resources, Resource{Name: "ground " + source.Name, Origin: source.Name, Type: "spice"})
	}

	return resources
}

func resourcesFromPlant(source Plant) []Resource {
	resources := []Resource{}

	if source.HasWood {
		resources = append(resources, Resource{Name: source.Name + " planks", Origin: source.Name, Type: "wood planks"})
		resources = append(resources, Resource{Name: source.Name + " logs", Origin: source.Name, Type: "wood logs"})
	}
	if source.IsFiber {
		if source.Name == "flax" {
			resources = append(resources, Resource{Name: "linen", Origin: source.Name, Type: "fabric"})
		} else {
			resources = append(resources, Resource{Name: source.Name + " cloth", Origin: source.Name, Type: "fabric"})
		}
		resources = append(resources, Resource{Name: source.Name + " rope", Origin: source.Name, Type: "rope"})
	}
	if source.IsFruit {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "fruit"})
	}
	if source.IsGrain {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "grain"})
	}
	if source.IsHerb {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "spice"})
	}
	if source.IsMedicine {
		resources = append(resources, Resource{Name: "refined " + source.Name, Origin: source.Name, Type: "medicine"})
	}
	if source.IsNut {
		resources = append(resources, Resource{Name: source.Name + " nut", Origin: source.Name, Type: "nut"})
	}
	if source.IsToxic {
		resources = append(resources, Resource{Name: "concentrated " + source.Name, Origin: source.Name, Type: "toxin"})
	}
	if source.IsVegetable {
		resources = append(resources, Resource{Name: source.Name, Origin: source.Name, Type: "vegetable"})
	}

	return resources
}
