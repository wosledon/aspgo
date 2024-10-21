package core

import "net/http"

type Middleware interface {
	Invoke(http.Handler) http.Handler
}
