package fun

import "fmt"

func TruncateAt(s string, size *int) (string, error) {
	l := len(s)
	if size == nil {
		return s, nil
	}

	if *size < 0 {
		return "", fmt.Errorf("`truncateAt` must be greater then 0")
	}

	if *size > l {
		*size = l
	}

	return s[:*size], nil
}
