package fun

import (
	"fmt"
	"github.com/gshopify/service-wrapper/interfaces"
	"github.com/gshopify/service-wrapper/model"
)

func Reverse[T any](slice []T) []T {
	var reversed []T
	for i := len(slice) - 1; i >= 0; i-- {
		reversed = append(reversed, slice[i])
	}

	return reversed
}

func First[T any](slice []T, first *int) []T {
	if len(slice) == 0 || first == nil {
		return slice[:]
	}

	if *first > len(slice) {
		*first = len(slice)
	}

	return slice[:*first]
}

func Last[T any](slice []T, last *int) []T {
	if len(slice) == 0 || last == nil {
		return slice[:]
	}

	if *last >= len(slice) {
		*last = 0
	} else if *last == 0 {
		*last = len(slice)
	}

	return slice[*last:]
}

func WithCursor[T interfaces.Node](slice []T, cursor *string, namespace model.Gid, after bool) ([]T, error) {
	if cursor == nil {
		return slice[:], nil
	}

	var (
		c   *model.Cursor
		err error
	)

	if c, err = model.ParseCursor(cursor); err != nil {
		return nil, fmt.Errorf("illegal cursor: %s", *cursor)
	}

	for i, node := range slice {
		rawId, err := model.ParseId(namespace, node.GetID())
		if err != nil {
			continue
		}

		if rawId != c.LastId {
			continue
		}

		if after {
			return slice[i+1:], nil
		} else {
			return slice[:i], nil
		}
	}

	return slice[:], nil
}
