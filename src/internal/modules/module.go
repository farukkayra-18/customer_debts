package modules

import "net/http"

type Module interface {
	RegisterRoutes(*http.ServeMux)
}
