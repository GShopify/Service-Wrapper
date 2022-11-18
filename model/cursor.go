package model

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type Cursor struct {
	namespace Gid
	LastId    string `json:"last_id"`
	LastValue string `json:"last_value"`
}

func NewSimpleCursor(id string, namespace Gid) (*Cursor, error) {
	rawId, err := ParseId(namespace, id)
	if err != nil {
		return nil, err
	}

	return &Cursor{
		namespace: namespace,
		LastId:    rawId,
		LastValue: rawId,
	}, nil
}

func ParseCursor(s *string) (*Cursor, error) {
	if s == nil {
		return nil, fmt.Errorf("empty cursor")
	}

	d, err := base64.StdEncoding.DecodeString(*s)
	if err != nil {
		return nil, fmt.Errorf("could not parse cursor data: %v", err)
	}

	cursor := Cursor{}
	if err = json.Unmarshal(d, &cursor); err != nil {
		return nil, fmt.Errorf("could not parse cursor data: %v", err)
	}

	return &cursor, nil
}

func (c *Cursor) String() *string {
	d, err := json.Marshal(c)
	if err != nil {
		return nil
	}

	s := base64.StdEncoding.EncodeToString(d)
	return &s
}
