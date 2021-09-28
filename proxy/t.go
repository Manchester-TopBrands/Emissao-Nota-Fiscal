package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/kr/pretty"
)

const (
	apiUrl   string = "https://api.enotasgw.com.br/v2/empresas/"
	resource string = `59C5E342-EDFA-4838-866A-4A58E14E0700/nf-e`
)

func mainn() {
	client := &http.Client{}

	b := new(bytes.Buffer)
	json.NewDecoder(b).Decode("oi")
	req, _ := http.NewRequest(http.MethodPost, apiUrl+resource, b)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Basic OGVjMTE5M2EtYThkOS00NTMzLTkzMDctMjVlOGRhNGUwNzAw")
	req.Header.Set("Content-Type", "application/json")

	resp, _ := client.Do(req)
	pretty.Println(resp.Request)
}
