package model

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type Cursor struct {
	id        string
	namespace Gid
	LastId    string `json:"last_id"`
	LastValue string `json:"last_value"`
}

func NewCursor(id string, namespace Gid) (*Cursor, error) {
	if namespace == "" {
		return NewSimpleCursor(id), nil
	}

	if !namespace.IsValid() {
		return nil, fmt.Errorf("illegal Gid: %s", namespace)
	}

	rawId, err := ParseId(namespace, id)
	if err != nil {
		return nil, err
	}

	return &Cursor{
		id:        id,
		namespace: namespace,
		LastId:    rawId,
		LastValue: rawId,
	}, nil
}

func NewSimpleCursor(id string) *Cursor {
	return &Cursor{
		id:        id,
		LastId:    id,
		LastValue: id,
	}
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
	if c.namespace == "" {
		s := base64.StdEncoding.EncodeToString([]byte(c.id))
		return &s
	}

	d, err := json.Marshal(c)
	if err != nil {
		return nil
	}

	s := base64.StdEncoding.EncodeToString(d)
	return &s
}
