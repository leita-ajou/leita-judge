package utils

import . "golang.org/x/exp/constraints"

func All[T comparable](s ...T) bool {
	var zero T
	for _, v := range s {
		if v == zero {
			return false
		}
	}
	return true
}

func Sum[T Integer | Float](s ...T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}
