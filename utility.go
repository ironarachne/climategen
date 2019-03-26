package climategen

func isResourceInSlice(resource Resource, resources []Resource) bool {
	for _, r := range resources {
		if resource.Name == r.Name {
			return true
		}
	}

	return false
}
