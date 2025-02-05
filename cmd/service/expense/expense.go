package expense

import (
	"fmt"
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

	router.HandleFunc("/expenses", h.Expenses).Methods("get")
	router.HandleFunc("/expense/{id}", h.Expense).Methods("get")
	router.HandleFunc("/addExpense", h.AddExpense).Methods("post")
	router.HandleFunc("/editExpense/{id}", h.AddExpense).Methods("put")
	router.HandleFunc("/deleteExpense/{id}", h.deleteExpense).Methods("delete")
}

// get all the expenses of a user
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

// get a single expense made by the user
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

// Add a new expense
func (h *Handler) AddExpense(w http.ResponseWriter, r *http.Request) {

	userID, err := utils.GetIDJwt(r)
	if err != nil {
		utils.WriteError(w, http.StatusForbidden, err)
		return
	}

	var expense datatypes.Expense
	err = utils.ParseJson(r, &expense)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request body: %v", err))
		return
	}

	err = h.store.AddExpense(&expense, userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to add expense: %v", err))
		return
	}

	utils.WriteJson(w, http.StatusCreated, "Expense added successfully")
}

// Edit an existing expense
func (h *Handler) EditExpense(w http.ResponseWriter, r *http.Request) {

	userID, err := utils.GetIDJwt(r)
	if err != nil {
		utils.WriteError(w, http.StatusForbidden, err)
		return
	}

	var expense datatypes.Expense
	err = utils.ParseJson(r, &expense)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request body: %v", err))
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	expense.ID = id
	err = h.store.EditExpense(&expense, userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to edit expense: %v", err))
		return
	}

	utils.WriteJson(w, http.StatusOK, "Expense updated successfully")
}

// delete expense
func (h *Handler) deleteExpense(w http.ResponseWriter, r *http.Request) {

	user_id, err := utils.GetIDJwt(r)

	if err != nil {
		utils.WriteError(w, http.StatusForbidden, err)
		return
	}
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)

	}

	err = h.store.DeleteExpense(id, user_id)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJson(w, http.StatusOK, "sucessfully deleted the expense")
}
