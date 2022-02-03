package balance

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/picunada/http-rest-balance/internal/handlers"
	"net/http"
)

type handler struct {
}

func NewHandler() hanlders.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/users/:uuid/balance", h.GetBalance)
	router.PUT("/users/:uuid/balance/payment", h.IncreaseBalance)
	router.PUT("/users/:uuid/balance/charge", h.DecreaseBalance)
	router.PUT("/balance/transfer ", h.TransferMoney)
}

func (h *handler) GetBalance(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("this is balance")))
}

func (h *handler) IncreaseBalance(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("this is balance")))
}

func (h *handler) DecreaseBalance(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("this is balance")))
}

func (h *handler) TransferMoney(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("this is balance")))
}
