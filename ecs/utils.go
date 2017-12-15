package ecs


// IntInSlice will return true if the integer value provided is present in the slide provided, false otherwise.
func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
