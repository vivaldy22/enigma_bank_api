package delivery

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vivaldy22/enigma_bank/models"
)

type UserHandler struct {
	UUseCase models.UserUseCase
}

func NewUserHandler(u models.UserUseCase, r *mux.Router) {
	handler := &UserHandler{u}
	r.HandleFunc("/users", handler.ShowUsers).Methods(http.MethodGet)

	userPref := r.PathPrefix("/user").Subrouter()
	userPref.HandleFunc("", handler.CreateUser).Methods(http.MethodPost)
	userPref.HandleFunc("/{id}", handler.GetUserByID).Methods(http.MethodGet)
	userPref.HandleFunc("/{id}", handler.UpdateUser).Methods(http.MethodPut)
	userPref.HandleFunc("/{id}", handler.RemoveUser).Methods(http.MethodDelete)
}
