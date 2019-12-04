package route

import "net/http"

type Routes struct {
	Pattern string
	Handle http.Handler
}