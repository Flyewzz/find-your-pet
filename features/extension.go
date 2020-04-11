package features

import "strings"

func GetExtension(path string) string {
	elements := strings.Split(path, ".")
	extension := elements[len(elements)-1]
	return extension
}

func IsExtensionPicture(extension string) bool {
	normalizeExtension := strings.ToLower(extension)
	if normalizeExtension != "png" && normalizeExtension != "jpg" {
		return false
	}
	return true
}
