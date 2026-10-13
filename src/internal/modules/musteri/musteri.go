package musteri

import (
	"net/http"

	"github.com/farukkayra-18/customer_debts/src/internal/modules"
	"github.com/farukkayra-18/customer_debts/src/internal/modules/musteri/handler"
)

type module struct {
	handler handler.Handler
}

func New(handler handler.Handler) modules.Module {
	return &module{
		handler: handler,
	}
}

func (m module) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /musteri", m.handler.Get)
	mux.HandleFunc("POST /musteri", m.handler.Post)
}
