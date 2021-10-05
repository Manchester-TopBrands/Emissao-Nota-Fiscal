package main

import (
	"embed"
	"flag"
	"io/fs"
	"log"
	"net/http"
	"strings"
	"time"
	apihandler "xmlconvert/apiHandler"
	c "xmlconvert/config"
	"xmlconvert/sql"
)

//go:embed html
var content embed.FS

var createConfig bool

func main() {

	flag.BoolVar(&createConfig, "c", false, "create config.yaml file")
	flag.Parse()

	if createConfig {
		c.CreateConfigFile()
		return
	}

	log.Print("loading config file")
	if err := c.LoadConfig(); err != nil {
		log.Fatal(err)
	}
	log.Print("connecting sql ...")
	connection, err := sql.MakeSQL(c.Config.SQL.Host, c.Config.SQL.Port, c.Config.SQL.User, c.Config.SQL.Password)
	if err != nil {
		log.Println(err)
		return
	}
	htmlFS, err := fs.Sub(content, "html")
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("starting server '%s' at port: %s", c.Config.API.Host, c.Config.API.Port)

	http.Handle("/download", &Auth{apihandler.Download, true})
	http.Handle("/addExcel", &Auth{apihandler.Excel(connection), true})
	http.Handle("/addEntradas", &Auth{apihandler.Entradas(connection), true})
	http.Handle("/addSaidas", &Auth{apihandler.Saidas(connection), true})
	//http.Handle("/addSaidas/homologacao", &Auth{apihandler.SaidasHomologacao})
	http.HandleFunc("/addLogin", apihandler.Login)
	http.HandleFunc("/", redirect)
	http.Handle("/addLogout", &Auth{apihandler.Logout, true})
	// fs := http.FileServer(http.Dir("html"))
	fs := http.FileServer(http.FS(htmlFS))
	http.Handle("/html/", http.StripPrefix("/html/", &Auth{fs.ServeHTTP, false}))
	http.HandleFunc("/enotasNF", apihandler.Enotas)
	log.Fatal(http.ListenAndServe(":8083", nil))
}

func redirect(w http.ResponseWriter, r *http.Request) {
	// remove/add not default ports from req.Host
	target := "http://" + r.Host + "/html/login.html"
	// log.Println(target)
	if len(r.URL.RawQuery) > 0 {
		target += "?" + r.URL.RawQuery
	}
	// log.Printf("redirect to: %s", target)
	http.Redirect(w, r, target,
		// see comments below and consider the codes 308, 302, or 301
		http.StatusTemporaryRedirect)
}

type Auth struct {
	handler http.HandlerFunc
	all     bool
}

func (a *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")

	if token, err := r.Cookie("Token"); err == nil {
		if _, ok := apihandler.Tokens[token.Value]; ok {
			if token.Expires.Before(time.Now()) {
				a.handler(w, r)
				return
			}
		}
	}

	urlS := strings.Split(r.URL.String(), ".")
	if a.all || (urlS[len(urlS)-1] == "html" && r.URL.String() != "login.html") {
		redirect(w, r)
		return
	}
	a.handler(w, r)
}
