package service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thapasubham/go-learn/cmd/datatypes"
	"github.com/thapasubham/go-learn/cmd/utils"
)

type Handler struct {
	store datatypes.UserStore
}

func NewHandler(store datatypes.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", h.handleRegister).Methods("post")
	router.HandleFunc("/users", h.handleGet).Methods("get")
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload datatypes.Employee

	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	fmt.Println(payload.Name + " " + payload.Address)

	rowAffected, err := h.store.CreateUser(datatypes.Employee{
		Name:    payload.Name,
		Address: payload.Address,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)

	}
	utils.WriteJson(w, http.StatusAccepted, rowAffected)
}

func (h *Handler) handleGet(w http.ResponseWriter, r *http.Request) {
	var payload datatypes.Employee

	err := h.store.GetUsers(&payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJson(w, http.StatusOK, payload)

}
