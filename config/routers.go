package config

import (
	"database/sql"
	"github.com/vivaldy22/enigma_bank/middleware"
	"log"
	"net/http"

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

	//mRepo := _mRepo.NewMovieRepo(db)
	//mUsc := _mUsc.NewMovieUseCase(mRepo)
	//_mDeliv.NewMovieHandler(mUsc, r)

}
