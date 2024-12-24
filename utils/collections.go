package utils

type Tuple[T1 any, T2 any] struct {
	A T1
	B T2
}

type Coord struct {
	X int
	Y int
}

// Pop removes and returns the last element of the slice
func Pop[T1 any](slice *[]T1) T1 {
	if len(*slice) == 0 {
		panic("cannot pop from an empty slice")
	}

	last := (*slice)[len(*slice)-1]
	*slice = (*slice)[:len(*slice)-1]

	return last
}

// RemoveFront removes the first element from a slice and returns the removed element and the updated slice
func RemoveFront[T1 any](slice *[]T1) *T1 {
	if len(*slice) == 0 {
		// Handle empty slice case
		return nil // Return default zero value for int and the empty slice
	}

	removed := (*slice)[0]
	*slice = (*slice)[1:]

	return &removed
}

func Bool2Int(v bool) int {
	if v {
		return 1
	}

	return 0
}
