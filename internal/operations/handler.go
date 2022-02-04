package operations

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
	router.GET("/users/:uuid/operations", h.GetBalance)
	router.PUT("/users/:uuid/operations/payment", h.IncreaseBalance)
	router.PUT("/users/:uuid/operations/charge", h.DecreaseBalance)
	router.PUT("/operations/transfer ", h.TransferMoney)
}

func (h *handler) GetBalance(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("this is operations")))
}

func (h *handler) IncreaseBalance(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("this is operations")))
}

func (h *handler) DecreaseBalance(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("this is operations")))
}

func (h *handler) TransferMoney(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte(fmt.Sprintf("this is operations")))
}
