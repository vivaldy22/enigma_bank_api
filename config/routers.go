package config

import (
	"database/sql"
	"log"
	"net/http"

	_lDeliv "github.com/vivaldy22/enigma_bank/master/login/delivery"
	_lRepo "github.com/vivaldy22/enigma_bank/master/login/repository"
	_lUsc "github.com/vivaldy22/enigma_bank/master/login/usecase"
	_tDeliv "github.com/vivaldy22/enigma_bank/master/transaction/delivery"
	_tRepo "github.com/vivaldy22/enigma_bank/master/transaction/repository"
	_tUsc "github.com/vivaldy22/enigma_bank/master/transaction/usecase"
	_uDeliv "github.com/vivaldy22/enigma_bank/master/user/delivery"
	_uRepo "github.com/vivaldy22/enigma_bank/master/user/repository"
	_uUsc "github.com/vivaldy22/enigma_bank/master/user/usecase"
	"github.com/vivaldy22/enigma_bank/middleware"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	return mux.NewRouter()
}

func RunServer(r *mux.Router, host, port string) {
	log.Printf("Starting Web Server at %v port: %v\n", host, port)

	if err := http.ListenAndServe(host+": "+port, r); err != nil {
		log.Fatal(err)
	}
}

func InitRouters(db *sql.DB, r *mux.Router) {
	r.Use(middleware.ActivityLogMiddleware)

	lRepo := _lRepo.NewLoginRepo(db)
	lUsc := _lUsc.NewLoginUseCase(lRepo)
	_lDeliv.NewLoginHandler(lUsc, r)

	uRepo := _uRepo.NewUserRepo(db)
	uUsc := _uUsc.NewUserUseCase(uRepo)
	_uDeliv.NewUserHandler(uUsc, r)

	tRepo := _tRepo.NewTransactionRepo(db)
	tUsc := _tUsc.NewTransactionUseCase(tRepo)
	_tDeliv.NewTransactionHandler(tUsc, r)

}
