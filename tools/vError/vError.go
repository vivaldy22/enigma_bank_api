package vError

import (
	"log"
	"net/http"
)

func WriteError(message string, err error, w *http.ResponseWriter) {
	(*w).Write([]byte(message))
	log.Println(err)
}
