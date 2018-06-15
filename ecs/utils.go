package ecs

import "reflect"

// IntInSlice will return true if the integer value provided is present in the slice provided, false otherwise.
func IntInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// TypeInSlice will return true if the reflect.Type provided is present in the slice provided, false otherwise.
func TypeInSlice(a reflect.Type, list []reflect.Type) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
