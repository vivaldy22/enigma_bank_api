package varMux

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetVarsMux(vars string, r *http.Request) string {
	muxVar := mux.Vars(r)
	return muxVar[vars]
}
