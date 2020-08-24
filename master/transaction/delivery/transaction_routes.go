package delivery

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vivaldy22/enigma_bank/models"
)

type TransactionHandler struct {
	TUseCase models.TransactionUseCase
}

func NewTransactionHandler(u models.TransactionUseCase, r *mux.Router) {
	handler := &TransactionHandler{u}
	r.HandleFunc("/transactions", handler.ShowTransactions).Methods(http.MethodGet)

	transactionPref := r.PathPrefix("/transaction").Subrouter()
	transactionPref.HandleFunc("", handler.CreateTransaction).Methods(http.MethodPost)
	transactionPref.HandleFunc("/{id}", handler.GetTransactionByID).Methods(http.MethodGet)
	transactionPref.HandleFunc("/user/{id}", handler.GetTransactionByUserOwnerID).Methods(http.MethodGet)
	transactionPref.HandleFunc("/{id}", handler.UpdateTransaction).Methods(http.MethodPut)
	transactionPref.HandleFunc("/{id}", handler.RemoveTransaction).Methods(http.MethodDelete)
}
