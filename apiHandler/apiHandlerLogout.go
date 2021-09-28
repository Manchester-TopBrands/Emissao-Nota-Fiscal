package apihandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"xmlconvert/models"
)

//Logout ...
func Logout(w http.ResponseWriter, r *http.Request) {

	if data, err := r.Cookie("Token"); err == nil {
		fmt.Println(data.Value)
		delete(Tokens, data.Value)
	} else {
		log.Println(err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("access-control-expose-headers", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	fmt.Println(Tokens)
	json.NewEncoder(w).Encode(models.Response{
		Status: "OK",
		Error:  "",
	})
}
