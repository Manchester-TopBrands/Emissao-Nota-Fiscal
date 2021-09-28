package apihandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// respEnotas
type respEnotas struct {
	Tipo                          string      `json:"tipo"`
	EmpresaID                     string      `json:"empresaId"`
	NfeID                         string      `json:"nfeId"`
	NfeStatus                     string      `json:"nfeStatus"`
	NfeMotivoStatus               string      `json:"nfeMotivoStatus"`
	NfeLinkDanfe                  string      `json:"nfeLinkDanfe"`
	NfeLinkXML                    string      `json:"nfeLinkXml"`
	NfeLinkConsultaPorChaveAcesso string      `json:"nfeLinkConsultaPorChaveAcesso"`
	ConteudoQRCode                interface{} `json:"conteudoQRCode"`
	NfeNumero                     string      `json:"nfeNumero"`
	NfeSerie                      string      `json:"nfeSerie"`
	NfeChaveAcesso                interface{} `json:"nfeChaveAcesso"`
	NfeDataEmissao                time.Time   `json:"nfeDataEmissao"`
	NfeDataAutorizacao            interface{} `json:"nfeDataAutorizacao"`
	NfeNumeroProtocolo            interface{} `json:"nfeNumeroProtocolo"`
	NfeDigestValue                interface{} `json:"nfeDigestValue"`
} // modelo para resposta do pedido de emissão, controlada pelo cefaz
var enotasResponses map[string]chan respEnotas = make(map[string]chan respEnotas)
var enotarResponsesLocker sync.Mutex

// handlerEnotas ...
func Enotas(w http.ResponseWriter, r *http.Request) {
	var rsp respEnotas
	if err := json.NewDecoder(r.Body).Decode(&rsp); err != nil {
		log.Println(err)
		return
	}

	enotarResponsesLocker.Lock()
	defer enotarResponsesLocker.Unlock()
	if _, ok := enotasResponses[rsp.NfeID]; ok {
		// pretty.Println(rsp)
		enotasResponses[rsp.NfeID] <- rsp
		return
	}
	log.Println("id not found")
	fmt.Println(r.URL.String())
}
