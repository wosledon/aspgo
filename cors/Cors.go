package cors

import (
	"aspgo/core"
	"net/http"
)

type CorsMiddleware struct {
	core.Middleware

	Allow *AccessControlAllow
}

type AccessControlAllow struct {
	Origin  string
	Methods string
	Headers string
}

func NewAccessControlAllow(origin, methods, headers string) *AccessControlAllow {
	return &AccessControlAllow{
		Origin:  origin,
		Methods: methods,
		Headers: headers,
	}
}

func NewCorsMiddleware(allow *AccessControlAllow) *CorsMiddleware {
	return &CorsMiddleware{
		Allow: allow,
	}
}

func (c *CorsMiddleware) Invoke(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", c.Allow.Origin)
		w.Header().Set("Access-Control-Allow-Methods", c.Allow.Methods)
		w.Header().Set("Access-Control-Allow-Headers", c.Allow.Headers)
		next.ServeHTTP(w, r)
	})
}
