package auth

import (
	"context"
	"fmt"
	"github.com/gshopify/service-wrapper/session"
	v3 "go.etcd.io/etcd/client/v3"
	"sync"
	"time"
)

var (
	once     sync.Once
	instance session.Session
)

type impl struct {
	cli     *v3.Client
	timeout time.Duration
}

func Init(etcd *v3.Client, timeout time.Duration) error {
	if etcd == nil {
		return fmt.Errorf("v3.Client must be provided to initialize SessionManager")
	}

	once.Do(func() {
		instance = &impl{
			cli:     etcd,
			timeout: timeout,
		}
	})

	return nil
}

func SessionManager() session.Session {
	return instance
}

func (s *impl) PutToken(parent context.Context, sid string, token string, ttl time.Duration) error {
	ctx, cancel := context.WithTimeout(parent, s.timeout)
	defer cancel()

	var leaseId v3.LeaseID

	if r, err := s.cli.Get(ctx, s.session(sid), v3.WithLastRev()...); err == nil && r.Count > 0 {
		leaseId = v3.LeaseID(r.Kvs[0].Lease)
	} else {
		lease, err := s.cli.Grant(parent, int64(ttl.Seconds()))
		if err != nil {
			return err
		}

		leaseId = lease.ID
	}

	if _, err := s.cli.Put(ctx, s.session(sid), token, v3.WithLease(leaseId)); err != nil {
		return err
	}

	return nil
}

func (s *impl) Token(parent context.Context, sid string) (string, error) {
	ctx, cancel := context.WithTimeout(parent, s.timeout)
	defer cancel()

	r, err := s.cli.Get(ctx, s.session(sid), v3.WithLastRev()...)
	if err != nil {
		return "", fmt.Errorf("could not retrieve SID: `%s`", sid)
	}

	if r.Count < 1 {
		return "", fmt.Errorf("could not retrieve SID: `%s`", sid)
	}

	return string(r.Kvs[0].Value), nil
}

func (s *impl) DeleteToken(parent context.Context, sid string) error {
	ctx, cancel := context.WithTimeout(parent, s.timeout)
	defer cancel()

	_, err := s.cli.Delete(ctx, s.session(sid))
	return err
}

func (s *impl) session(sid string) string {
	return fmt.Sprintf("component/session/%s", sid)
}
