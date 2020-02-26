package api

import (
	"encoding/json"
	"log"
	"net/http"
)

//PingPongHandler test handler
func PingPongHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("It's alive!")
	res := map[string]string{"ping": "pong"}
	json.NewEncoder(w).Encode(res)
}
