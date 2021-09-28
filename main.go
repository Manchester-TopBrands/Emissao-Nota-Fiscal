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
	apihandler.SetSQLConn(connection)
	htmlFS, err := fs.Sub(content, "html")
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("starting server '%s' at port: %s", c.Config.API.Host, c.Config.API.Port)

	http.Handle("/download", &Auth{apihandler.Download, true})
	http.Handle("/addExcel", &Auth{apihandler.Excel, true})
	http.Handle("/addEntradas", &Auth{apihandler.Entradas, true})
	http.Handle("/addSaidas", &Auth{apihandler.Saidas, true})
	//http.Handle("/addSaidas/homologacao", &Auth{apihandler.SaidasHomologacao})
	http.HandleFunc("/addLogin", apihandler.Login)
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

	// if s[len(s)-1] == "html" {
	// }

	if token, err := r.Cookie("Token"); err == nil {
		if _, ok := apihandler.Tokens[token.Value]; ok {
			if token.Expires.Before(time.Now()) {
				a.handler(w, r)
				return
			}
			//deletar token vencido- verificar se ja deleta automatico
		}
	}

	urlS := strings.Split(r.URL.String(), ".")
	if a.all || (urlS[len(urlS)-1] == "html" && r.URL.String() != "login.html") {
		// newUrl, _ := url.Parse("login.html")
		// fmt.Printf("old: %+v, new: %+v\n", r.URL, newUrl)
		// r.URL = newUrl
		redirect(w, r)
		// target := "http://" + r.Host + "/html/login.html"
		// if len(r.URL.RawQuery) > 0 {
		// 	target += "?" + r.URL.RawQuery
		// }
		// http.Redirect(w, r, target,
		// 	http.StatusTemporaryRedirect)
		return
	}
	a.handler(w, r)
}

// func IsAuthorized(handler http.Handler) http.Handler {
// 	var a *Auth
// 	return a
// }

// go func() {
// 	time.Sleep(time.Second * 2)
// 	f, _ := os.Open("PEDIDO2.xlsx")
// 	// req, _ := http.PostForm("http://localhost:8080", url.Values{"file": f})
// 	var b bytes.Buffer
// 	m := multipart.NewWriter(&b)
// 	w, _ := m.CreateFormFile("file", "pedidos.xlsx")
// 	io.Copy(w, f)
// 	m.Close()

// 	req, err := http.NewRequest("POST", "http://localhost:8082/addExcel", &b)
// 	if err != nil {
// 		return
// 	}

// Don't forget to set the content type, this will contain the boundary.
// req.Header.Set("Content-Type", m.FormDataContentType())

// // Submit the request
// res, err := (&http.Client{}).Do(req)
// if err != nil {
// 	return
// }
// var t table
// json.NewDecoder(res.Body).Decode(&t)
// res, err = http.Post("http://localhost:8082/addSaidas", "application/json", strings.NewReader(t.ID))
// if err != nil {
// 	fmt.Println(err)
// 	return
// }
// fmt.Println(res.Status)
// io.Copy(os.Stdout, res.Body)
// }()
