package respJson

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(resp interface{}, w http.ResponseWriter) {
	jsonResp, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
