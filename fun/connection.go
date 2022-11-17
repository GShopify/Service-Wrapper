package fun

func Reverse[T any](slice []T) []T {
	var reversed []T
	for i := len(slice) - 1; i >= 0; i-- {
		reversed = append(reversed, slice[i])
	}

	return reversed
}

func First[T any](slice []T, first *int) []T {
	if len(slice) == 0 || first == nil {
		return slice
	}

	if *first > len(slice) {
		*first = len(slice)
	}

	return slice[:*first]
}

func Last[T any](slice []T, last *int) []T {
	if len(slice) == 0 || last == nil {
		return slice
	}

	if *last >= len(slice) {
		*last = 0
	} else if *last == 0 {
		*last = len(slice)
	}

	return slice[*last:]
}
