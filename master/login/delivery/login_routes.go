package delivery

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vivaldy22/enigma_bank/models"
)

type LoginHandler struct {
	LUseCase models.LoginUseCase
}

func NewLoginHandler(u models.LoginUseCase, r *mux.Router) {
	handler := &LoginHandler{u}
	r.HandleFunc("/logins", handler.ShowLogins).Methods(http.MethodGet)

	loginPref := r.PathPrefix("/login").Subrouter()
	loginPref.HandleFunc("", handler.CreateLogin).Methods(http.MethodPost)
	loginPref.HandleFunc("/{id}", handler.GetLoginByID).Methods(http.MethodGet)
	loginPref.HandleFunc("/{id}", handler.UpdateLogin).Methods(http.MethodPut)
	loginPref.HandleFunc("/{id}", handler.RemoveLogin).Methods(http.MethodDelete)
}
