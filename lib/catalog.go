package lib

type catalog struct {
	Images map[string][]image
}

type image struct {
	Default      bool   // Whether the image is currently marked as the default
	Distribution string // The distribution this image belongs to
	Path         string // Path to the adapted image file
	Hash         string // Hash for the image file
}

// Load a catalog from the specified location
func LoadCatalog(path string) catalog {
	return catalog{
		Images: make(map[string][]image),
	}
}
