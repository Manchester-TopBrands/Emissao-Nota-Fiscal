package apihandler

import (
	"encoding/json"
	"log"
	"net/http"
	"xmlconvert/models"

	auth "github.com/korylprince/go-ad-auth/v3"

	c "xmlconvert/config"
)

// Token ...
var Tokens map[string]byte = map[string]byte{}

func Login(w http.ResponseWriter, r *http.Request) {
	config := &auth.Config{
		Server:   c.Config.AUTH.Server,
		Port:     c.Config.AUTH.Port,
		BaseDN:   c.Config.AUTH.BaseDN,
		Security: auth.SecurityNone,
	}
	client := models.Login{}
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Status: "Bad Request",
			Error:  "",
			Data:   err.Error(),
		})
		return
	}

	status, _, _, err := auth.AuthenticateExtended(config, client.Username, client.Userpassword, nil, []string{"Manchester"})
	if err != nil {
		log.Fatal(err)
	}
	if !status {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.Response{
			Status: "Unauthorized",
			Error:  "",
			Data:   "Senha ou username inválidos",
		})
		return
	}
	delete(Tokens, "1")
	token := genKey()
	Tokens[token] = 1
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("access-control-expose-headers", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Token", token)
	w.Header().Set("Name", client.Username)
	json.NewEncoder(w).Encode(models.Response{
		Status: "OK",
		Error:  "",
	})
}
