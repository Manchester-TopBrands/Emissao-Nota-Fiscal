package apihandler

import (
	_ "embed" //f
	"net/http"
)

//go:embed modelpedido.xlsx
var ExcelPedido []byte

// Download ...
func Download(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("access-control-expose-headers", "*")
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(ExcelPedido)
}
