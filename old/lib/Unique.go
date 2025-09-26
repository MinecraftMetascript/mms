package lib

func Unique[T comparable](arr []T) []T {
	seen := make(map[T]bool) // A map to store seen elements
	result := make([]T, 0)   // The array to store unique elements

	for _, v := range arr {
		if _, ok := seen[v]; !ok { // If element not in map (not seen yet)
			seen[v] = true             // Mark as seen
			result = append(result, v) // Add to result
		}
	}
	return result
}
