package util

// GetAndConvert gets a value from a map[string]interface of a particular type, if it exists
func GetAndConvert[T any](name string, data map[string]interface{}) (T, bool) {
	var zero T
	raw, exists := data[name]
	if !exists {
		return zero, false
	}

	out, ok := raw.(T)
	return out, ok
}
