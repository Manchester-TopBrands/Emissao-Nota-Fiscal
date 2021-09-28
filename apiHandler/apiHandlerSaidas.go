package apihandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"xmlconvert/models"
)

type req struct {
	ID       string `json:"id,omitempty"`
	Producao bool   `json:"producao,omitempty"`
}

func Saidas(w http.ResponseWriter, r *http.Request) {
	var err error

	var resp req
	if err = json.NewDecoder(r.Body).Decode(&resp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.RespLink{
			Status: "formatacao json errado",
			PDF:    "",
			XML:    "",
			Error:  err.Error(),
		})
		return
	}
	fmt.Println(resp)

	ambiente := "Homologacao"
	if resp.Producao {
		ambiente = "Producao"
	}

	var nfs = tempNfs[resp.ID]

	var nfResults map[string]*models.RespLink = make(map[string]*models.RespLink)

	for name, nf := range nfs {

		contribuinte := "Contribuinte"
		if nf.Contribuinte == "0" {
			contribuinte = "NaoContribuinte"
		}

		enotasNF := models.NfeStruct{
			Tipo:             "NF-e",
			ID:               fmt.Sprintf("%s_%s", name, time.Now().Format("2006-01-02_150405")),
			AmbienteEmissao:  ambiente,
			NaturezaOperacao: "Venda",
			TipoOperacao:     "Saida",
			Finalidade:       "Normal",
			ConsumidorFinal:  false,
			EnviarPorEmail:   false,
			Cliente: &models.ClientCadastro{
				TipoPessoa:                "J",
				IndicadorContribuinteICMS: contribuinte,
				Nome:                      nf.Name,
				Email:                     nf.Email,
				Telefone:                  "0000000000",
				CpfCnpj:                   nf.Cnpj,
				InscricaoEstadual:         "1234454324543",
			},
			Itens:                 []models.ProdutoCadastro{},
			InformacoesAdicionais: nf.AddInfo,
		}

		enotasNF.Transporte.Frete.Modalidade = nf.ModFrete
		enotasNF.Transporte.Volume.Especie = nf.EspecieVol
		enotasNF.Transporte.Volume.Numeracao = strconv.Itoa(nf.QtdVol)
		enotasNF.Transporte.Volume.Quantidade = nf.QtdVol
		enotasNF.Transporte.Volume.PesoLiq = nf.PesoLiq
		enotasNF.Transporte.Volume.PesoBruto = nf.PesoBruto

		nf.Uf = strings.ToUpper(nf.Uf)

		enotasNF.Cliente.Endereco.UF = nf.Uf
		enotasNF.Cliente.Endereco.Cidade = nf.Cidade
		enotasNF.Cliente.Endereco.Logradouro = nf.Logradouro
		enotasNF.Cliente.Endereco.Numero = nf.Nr
		enotasNF.Cliente.Endereco.Complemento = nf.Complemento
		enotasNF.Cliente.Endereco.Bairro = nf.Bairro
		enotasNF.Cliente.Endereco.CEP = nf.Cep

		cfop := "6102"
		if nf.Uf == "SC" {
			cfop = "5102"
		}

		for code, item := range nf.Items {
			enotasItem := models.ProdutoCadastro{
				Codigo:        code,
				Quantidade:    item.Qty,
				ValorUnitario: item.Value,
				Cfop:          cfop,
				Descricao:     item.Desc,
				Ncm:           item.Ncm,
				Unidademedida: item.Unmedida,
				Frete:         0.00,
			}
			enotasItem.Impostos.Icms.SituacaoTributaria = "00"
			enotasItem.Impostos.Icms.Origem = "1"
			enotasItem.Impostos.Icms.Aliquota = 4.0
			enotasItem.Impostos.Pis.SituacaoTributaria = "01"
			enotasItem.Impostos.Pis.PorAliquota.Aliquota = 0.65
			enotasItem.Impostos.Cofins.SituacaoTributaria = "01"
			enotasItem.Impostos.Cofins.PorAliquota.Aliquota = 3

			enotasNF.Itens = append(enotasNF.Itens, enotasItem)

		}
		var rsp *respEnotas
		rsp, err = EmitirNFe(&enotasNF, resp.Producao)

		if err != nil {
			nfResults[enotasNF.ID] = &models.RespLink{
				Status: "erro interno",
				PDF:    "",
				XML:    "",
				Error:  err.Error(),
			}
			continue
		}

		nfResults[enotasNF.ID] = &models.RespLink{
			Status: rsp.NfeStatus,
			PDF:    rsp.NfeLinkDanfe,
			XML:    rsp.NfeLinkXML,
		}

	}

	// io.Copy(os.Stdout, r.Body)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("access-control-expose-headers", "*")
	w.Header().Set("Content-Type", "application/octet-stream")
	// if rsp
	json.NewEncoder(w).Encode(models.Response{
		Status: "OK",
		Error:  "",
		Data:   nfResults,
	})

}
