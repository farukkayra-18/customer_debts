package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/farukkayra-18/customer_debts/src/internal/modules/musteri/service"
)

type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service service.Service
}

func New(service service.Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	ret, err := h.service.TumMusterileriGetir()
	if err != nil {
		log.Printf("error occured: %s", err.Error())
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&ret); err != nil {
		log.Printf("error occured when encoding result: %s \n", err.Error())
	}
}

func (h *handler) Post(w http.ResponseWriter, r *http.Request) {

}
