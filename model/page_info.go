package model

import (
	"github.com/gshopify/service-wrapper/interfaces"
)

type PageInfo struct {
	EndCursor       *string `json:"endCursor"`
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
}

func NewPageInfo[T interfaces.Node](src []T, namespace Gid, first T, last T) *PageInfo {
	var (
		size = len(src)
		info = &PageInfo{}
	)

	if size == 0 {
		return info
	}

	if c, err := NewCursor(src[0].GetID(), namespace); err == nil {
		info.StartCursor = c.String()
	}

	if c, err := NewCursor(src[size-1].GetID(), namespace); err == nil {
		info.EndCursor = c.String()
	}

	info.HasPreviousPage = first.GetID() != src[0].GetID()
	info.HasNextPage = info.StartCursor == info.EndCursor || last.GetID() != src[size-1].GetID()
	return info
}
