package session

import (
	"context"
	"time"
)

type Session interface {
	PutToken(parent context.Context, sid string, token string, ttl time.Duration) error
	Token(parent context.Context, sid string) (string, error)
	DeleteToken(parent context.Context, sid string) error
}
