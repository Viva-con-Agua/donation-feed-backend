package broadcastChannel

// Filter the given slice by applying the given filter function.
// If the filter function returns true, the element will be kept, otherwise it will be filtered out.
func filterSlice[T any](slice []T, filter func(elem T) bool) []T {
	result := make([]T, 0)
	for _, elem := range slice {
		if filter(elem) {
			result = append(result, elem)
		}
	}
	return result
}
