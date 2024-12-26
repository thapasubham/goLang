package expense

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/thapasubham/go-learn/cmd/datatypes"
	"github.com/thapasubham/go-learn/cmd/utils"
)

type Handler struct {
	store datatypes.ExpenseStore
}

func NewHandler(store datatypes.ExpenseStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {

	router.HandleFunc("/expenses/", h.Expenses).Methods("get")
	router.HandleFunc("/expense/{id}", h.Expense).Methods("get")
}

func (h *Handler) Expenses(w http.ResponseWriter, r *http.Request) {

	var expenses []datatypes.Expense

	user_id, err := utils.GetIDJwt(r)

	if err != nil {
		utils.WriteError(w, http.StatusForbidden, err)
		return
	}

	err = h.store.GetExpenses(&expenses, user_id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJson(w, http.StatusAccepted, expenses)
}

func (h *Handler) Expense(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var expenses datatypes.Expense

	err := h.store.GetExpense(&expenses, id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJson(w, http.StatusAccepted, expenses)
}
