package middleware

import (
	"context"
	"fmt"
	"github.com/gshopify/service-wrapper/model"
	"net/http"
	"strings"
)

const (
	keyGShopify = Key("Context-Key: GShopify/InContext")
)

type GShopifyContext struct {
	RequestId string
	Country   model.CountryCode
	Language  model.LanguageCode
}

func (c *GShopifyContext) String() string {
	return fmt.Sprintf("🍑 RequestId: %s\n🍑 Country: %s\n🍑 Language: %s\n",
		c.RequestId,
		c.Country,
		c.Language)
}

func GShopify() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := GShopifyContext{}

			if rid := r.Header.Get("X-Request-Id"); rid != "" {
				ctx.RequestId = strings.TrimSpace(rid)
			} else {
				w.WriteHeader(http.StatusPreconditionFailed)
				return
			}

			if s := r.Header.Get("X-GShopify-Country"); s != "" {
				ctx.Country = model.CountryCode(s)
			} else {
				ctx.Country = model.CountryCodeUs
			}

			if s := r.Header.Get("X-GShopify-Language"); s != "" {
				ctx.Language = model.LanguageCode(s)
			} else {
				ctx.Language = model.LanguageCodeEn
			}

			if !ctx.Country.IsValid() || !ctx.Language.IsValid() {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			next.ServeHTTP(w, r.WithContext(
				context.WithValue(r.Context(), keyGShopify, ctx)))
		})
	}
}

func InContext(ctx context.Context) (*GShopifyContext, error) {
	if ctx == nil {
		return nil, fmt.Errorf("`context.Context` must not be a nil")
	}

	if v, ok := ctx.Value(keyGShopify).(GShopifyContext); ok {
		return &v, nil
	}

	return nil, fmt.Errorf("there is no defined `GShopifyContext` was attached to Context")
}
