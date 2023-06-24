package route

import "strings"

// AsRouteName get a string and transform into a internal route name representation
func AsRouteName(raw string) string {
	return strings.Trim(raw, "/")
}

func isMapSubset(baseMap, mapToCheck map[string]string) bool {
	// Check headers
	for key, value := range mapToCheck {
		key = strings.ToLower(key)
		headerValue, exists := baseMap[key]
		if !exists {
			return false
		}
		if headerValue != value {
			return false
		}
	}
	return true
}
