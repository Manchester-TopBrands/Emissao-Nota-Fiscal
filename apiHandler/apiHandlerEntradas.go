package apihandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"xmlconvert/models"
)

// Entradas ...
func Entradas(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	object, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	data := UnMarshal(object)

	if err = sqlConn.AddEntradas(&data); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(models.Response{
			Status: "Erro ao dar entrada",
			Error:  err.Error(),
		}); err != nil {
			fmt.Println(err)
		}
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("access-control-expose-headers", "*")
	w.Header().Set("Content-Type", "application/octet-stream")

	fn := strings.Split(header.Filename, ".")
	if len(fn) > 1 {
		fn[len(fn)-1] = "json"
	}
	w.Header().Set("File-Name", strings.Join(fn, "."))

	if err := json.NewEncoder(w).Encode(models.Response{
		Status: "OK",
		Data:   data,
	}); err != nil {
		fmt.Println(err)
	}
}
