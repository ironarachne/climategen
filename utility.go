package climategen

// IsResourceInSlice checks to see if a given resource is in a given slice of resources
func IsResourceInSlice(resource Resource, resources []Resource) bool {
	for _, r := range resources {
		if resource.Name == r.Name {
			return true
		}
	}

	return false
}

// IsTypeInResources checks to see if a resource type is present in a slice of resources
func IsTypeInResources(resourceType string, resources []Resource) bool {
	for _, r := range resources {
		for _, t := range r.Types {
			if t == resourceType {
				return true
			}
		}
	}

	return false
}

// ListResourcesOfType returns a slice of resources that match the type given
func ListResourcesOfType(resourceType string, resources []Resource) []Resource {
	var found []Resource

	for _, r := range resources {
		for _, t := range r.Types {
			if t == resourceType {
				found = append(found, r)
			}
		}
	}

	return found
}
