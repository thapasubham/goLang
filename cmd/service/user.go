package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/thapasubham/go-learn/cmd/datatypes"
	"github.com/thapasubham/go-learn/cmd/utils"
)

type Handler struct {
	store datatypes.UserStore
}

// new handler
func NewHandler(store datatypes.UserStore) *Handler {
	return &Handler{store: store}
}

// routes
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/user", h.handleInsert).Methods("post")
	router.HandleFunc("/user", h.handleGet).Methods("get")
	router.HandleFunc("/user/{id}", h.handleEdit).Methods("put")
	router.HandleFunc("/user/{id}", h.handleDelete).Methods("delete")
}

// create user
func (h *Handler) handleInsert(w http.ResponseWriter, r *http.Request) {
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

// get single user
func (h *Handler) handleGet(w http.ResponseWriter, r *http.Request) {
	var payload datatypes.Employee

	err := h.store.GetUser(&payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJson(w, http.StatusOK, payload)

}

// edit
func (h *Handler) handleEdit(w http.ResponseWriter, r *http.Request) {
	var payload datatypes.Employee
	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	vars := mux.Vars(r)
	temp_ID := vars["id"]
	id, _ := strconv.Atoi(temp_ID)
	err := h.store.EditUser(id, payload)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	} else {
		utils.WriteJson(w, http.StatusAccepted, "Updated")
	}
}

// delete user
func (h *Handler) handleDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if err := h.store.DeleteUser(id); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	} else {
		utils.WriteJson(w, http.StatusOK, "Sucessfully Deleted")
	}

}
