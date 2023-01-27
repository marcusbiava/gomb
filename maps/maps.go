package maps

func Keys[K comparable, V any](mapInstance map[K]V) []K {
	keys := make([]K, len(mapInstance))
	i := 0
	for k := range mapInstance {
		keys[i] = k
		i++
	}

	return keys
}

func Values[K comparable, V any](mapInstance map[K]V) []V {
	values := make([]V, len(mapInstance))
	i := 0
	for _, v := range mapInstance {
		values[i] = v
		i++
	}
	return values
}
