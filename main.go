package main

import (
	"embed"
	"flag"
	"io/fs"
	"log"
	"net/http"
	apihandler "xmlconvert/apiHandler"
	"xmlconvert/auth"
	c "xmlconvert/config"
	"xmlconvert/sql"
)

//go:embed html
var content embed.FS

var createConfig bool

func init() {
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
}
func main() {

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

	Auth := auth.New(nil)

	http.Handle("/download", Auth.Layer(apihandler.Download, "/html/login.html", true))
	http.Handle("/addExcel", Auth.Layer(apihandler.Excel(connection), "/html/login.html", true))
	http.Handle("/addEntradas", Auth.Layer(apihandler.Entradas(connection), "/html/login.html", true))
	http.Handle("/addSaidas", Auth.Layer(apihandler.Saidas(connection), "/html/login.html", true))
	//http.Handle("/addSaidas/homologacao", Auth.Layer(apihandler.SaidasHomologacao})
	http.HandleFunc("/addLogin", apihandler.Login(Auth))
	http.HandleFunc("/", redirect)
	// http.Handle("/addLogout", Auth.Layer(apihandler.Logout, true})
	// fs := http.FileServer(http.Dir("html"))
	fs := http.FileServer(http.FS(htmlFS))
	http.Handle("/html/", http.StripPrefix("/html/", Auth.Layer(func(w http.ResponseWriter, r *http.Request, u string) { fs.ServeHTTP(w, r) }, "/login.html", false)))
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
