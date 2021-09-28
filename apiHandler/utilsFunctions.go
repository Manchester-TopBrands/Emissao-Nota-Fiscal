package apihandler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"
	"xmlconvert/models"
	"xmlconvert/sql"
)

const (
	apiUrl   string = "https://api.enotasgw.com.br"
	apikey   string = `Basic OGVjMTE5M2EtYThkOS00NTMzLTkzMDctMjVlOGRhNGUwNzAw`
	resource string = `/v2/empresas/59C5E342-EDFA-4838-866A-4A58E14E0700/nf-e`
)

var tempNfs map[string]map[string]*models.Nf = map[string]map[string]*models.Nf{}
var locker sync.Mutex

var sqlConn *sql.SQLStr

func init() {
	rand.Seed(time.Now().UnixNano())
}

//SetSQLConn ...
func SetSQLConn(c *sql.SQLStr) {
	sqlConn = c
}

// genKey ...
func genKey() string {
	b := make([]byte, 12)
	rand.Read(b)

	return base64.URLEncoding.EncodeToString(b)
}

// UnMarshal ...
func UnMarshal(object []byte) models.XmlFormat {
	data := models.XmlFormat{}
	err := xml.Unmarshal(object, &data)
	if nil != err {
		fmt.Println("Error unmarshalling from XML", err)
	}
	return data
}

// responseExcel ...
func responseExcel(nfs map[string]*models.Nf, stocks map[string]float64, idnfe string) models.StructExcel {
	var t models.StructExcel
	t.Clientes = make([]string, len(nfs))
	t.Items = make([]*models.StructExcelItem, len(stocks))
	t.ID = idnfe
	i := 0
	for client := range nfs {
		t.Clientes[i] = client
		i++
	}
	i = 0
	for code, stock := range stocks {
		t.Items[i] = &models.StructExcelItem{
			CodeDesc: code,
			Price:    0,
			Qties:    make([]float64, len(nfs)+1),
		}
		t.Items[i].Qties[0] = stock
		for clientCode, clientName := range t.Clientes {
			if item, ok := nfs[clientName].Items[code]; ok {
				t.Items[i].Price = item.Value
				t.Items[i].Qties[clientCode+1] = item.Qty
			}
		}
		i++
	}
	//pretty.Print(t)
	return t
}

// EmitirNFeProducao ...
func EmitirNFe(nf *models.NfeStruct, prod bool, username string) (*respEnotas, error) {
	client := &http.Client{}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(nf)
	req, _ := http.NewRequest(http.MethodPost, apiUrl+resource, b)
	req.Header.Set("Authorization", apikey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	resp, _ := client.Do(req)
	fmt.Println("emitir nf:", resp.Status)

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("nota invalida: %s", resp.Status)
	}

	c := make(chan respEnotas)

	enotarResponsesLocker.Lock()
	enotasResponses[nf.ID] = c
	enotarResponsesLocker.Unlock()

	var response respEnotas
	var err error

	t := time.NewTimer(time.Second * 200)

	select {
	case response = <-c:
	case <-t.C:
		err = fmt.Errorf("timeout")
	}

	enotarResponsesLocker.Lock()
	delete(enotasResponses, nf.ID)
	enotarResponsesLocker.Unlock()

	if err != nil {
		return nil, err
	}

	err = DownloadFile(response.NfeLinkXML, prod, username)
	return &response, err
}

// DownloadFileProducao ...
func DownloadFile(url string, prod bool, username string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if !prod {
		return nil
	}

	b, _ := ioutil.ReadAll(resp.Body)
	data := UnMarshal(b)

	return sqlConn.AddSaidas(&data, username)
}
