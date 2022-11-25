package fun

import (
	"fmt"
	"strings"
)

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

	return strings.TrimSpace(s[:*size]), nil
}
