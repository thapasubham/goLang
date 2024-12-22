package user

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

// routes
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/user", h.handleInsert).Methods("post")
	router.HandleFunc("/login", h.logIn).Methods("post")

}

// create user
func (h *Handler) handleInsert(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	var payload datatypes.User

	if err := utils.ParseJson(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	fmt.Println(payload.Username + " and " + payload.Password)

	rowAffected, err := h.store.CreateUser(datatypes.User{
		Username: payload.Username,
		Password: payload.Password,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)

	}
	utils.WriteJson(w, http.StatusAccepted, rowAffected)
}

func (h *Handler) logIn(w http.ResponseWriter, r *http.Request) {
	utils.EnableCors(&w)
	var payload datatypes.User
	err := utils.ParseJson(r, &payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	id, err := h.store.LogIn(payload)
	if err != nil {
		utils.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	secret := []byte(utils.LoadEnv("SECRET"))
	fmt.Println(secret)
	token, err := utils.CreateJWT(secret, id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"token": token})

}
