package auth

import (
	"encoding/json"
	"net/http"
	"strings"
	"xmlconvert/models"
)

// Auth ..
type Auth struct {
	key []byte
}

func (a *Auth) Layer(handler func(w http.ResponseWriter, r *http.Request, user string), redirectURL string, api bool) *Layer {
	return &Layer{
		a:           a,
		Handler:     handler,
		Api:         api,
		redirectURL: redirectURL,
	}
}

type Layer struct {
	a           *Auth
	Handler     func(w http.ResponseWriter, r *http.Request, user string)
	Api         bool
	redirectURL string
}

func (l *Layer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")

	boole := false
	username := ""
	if token, err := r.Cookie("Token"); err == nil {
		boole, username, _ = l.a.VerifyToken(token.Value, 3600)
	}

	if boole {
		l.Handler(w, r, username)
		return
	}

	if l.Api {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Status: "NotAuthorized",
			Error:  "",
		})
		return
	}

	urlS := strings.Split(r.URL.String(), ".")

	if urlS[len(urlS)-1] != "html" || r.URL.String() == "login.html" {
		l.Handler(w, r, "")
		return
	}

	redirect(w, r, l.redirectURL)
}

// Jwt ...
type Jwt struct {
	Username string `json:"username"`
	Iat      int64  `json:"iat"`
}

func redirect(w http.ResponseWriter, r *http.Request, url string) {
	// remove/add not default ports from req.Host
	target := "http://" + r.Host + url
	// log.Println(target)
	if len(r.URL.RawQuery) > 0 {
		target += "?" + r.URL.RawQuery
	}
	// log.Printf("redirect to: %s", target)
	http.Redirect(w, r, target,
		// see comments below and consider the codes 308, 302, or 301
		http.StatusTemporaryRedirect)
}
