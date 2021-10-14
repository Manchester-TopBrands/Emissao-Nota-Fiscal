package apihandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"xmlconvert/models"
	"xmlconvert/sql"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func Excel(s *sql.SQLStr) func(w http.ResponseWriter, r *http.Request, user string) {
	return func(w http.ResponseWriter, r *http.Request, user string) {
		file, header, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		f, err := excelize.OpenReader(file)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		stocks := make(map[string]float64)
		nfs := make(map[string]*models.Nf)

		if f.GetSheetIndex("faturamento") == 0 || f.GetSheetIndex("produtos") == 0 {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(models.Response{
				Status: "excel errado",
				Error:  "nao existe a aba faturamento ou produtos",
			})
			return
		}

		rows := f.GetRows("faturamento")
		for i := 1; i < len(rows); i++ {
			if len(rows[i]) < 14 {
				rows[i] = append(rows[i], make([]string, 14-len(rows[i]))...)
			}

			if _, ok := nfs[rows[i][0]]; ok {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(models.Response{
					Status: "excel errado",
					Error:  "ref de pedido duplicada",
				})
				return
			}

			mod, err := strconv.ParseFloat(rows[i][14], 64)
			if err != nil {
				log.Println(err)
			}

			qtd, err := strconv.ParseFloat(rows[i][15], 64)
			if err != nil {
				log.Println(err)
			}
			pl, err := strconv.ParseFloat(rows[i][17], 64)
			if err != nil {
				log.Println(err)
			}
			pb, err := strconv.ParseFloat(rows[i][18], 64)
			if err != nil {
				log.Println(err)
			}

			nfs[rows[i][0]] = &models.Nf{
				Nf:           rows[i][0],
				Name:         rows[i][1],
				Cnpj:         rows[i][2],
				Ie:           rows[i][3],
				Contribuinte: rows[i][4],
				Email:        rows[i][5],
				Uf:           rows[i][6],
				Cidade:       rows[i][7],
				Logradouro:   rows[i][8],
				Nr:           rows[i][9],
				Complemento:  rows[i][10],
				Bairro:       rows[i][11],
				Cep:          rows[i][12],
				AddInfo:      rows[i][13],
				ModFrete:     models.Modalidade[int(mod)],
				QtdVol:       int(qtd),
				EspecieVol:   rows[i][16],
				PesoLiq:      pl,
				PesoBruto:    pb,
				Items:        make(map[string]*models.Item),
			}
		}

		rows = f.GetRows("produtos")
		for i := 1; i < len(rows); i++ {
			if len(rows[i]) < 4 {
				rows[i] = append(rows[i], make([]string, 4-len(rows[i]))...)
			}

			nf, ok := nfs[rows[i][0]]
			if !ok {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(models.Response{
					Status: "excel errado",
					Error:  "ref de pedido em \"produtos\" nao existente",
				})
				return
			}

			if _, ok := nf.Items[rows[i][1]]; ok {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(models.Response{
					Status: "excel errado",
					Error:  "ref do item em \"produtos\" duplicado na mesmo pedido",
				})
				return
			}

			qty, _ := strconv.ParseFloat(rows[i][2], 64)
			value, _ := strconv.ParseFloat(rows[i][3], 64)
			nf.Items[rows[i][1]] = &models.Item{
				Code:  rows[i][1],
				Qty:   qty,
				Value: value,
			}

			stocks[rows[i][1]] = 0

		}
		fmt.Println(header.Filename)
		// pretty.Println(nfs)

		codes := make([]string, len(stocks))
		i := 0
		for code := range stocks {
			codes[i] = code
			i++
		}

		productSql, err := s.GetStock(codes)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.Response{
				Status: "SQL Error",
				Error:  err.Error(),
			})
			return
		}

		for code, product := range productSql {
			stocks[code] = product.Qtd
			for _, nf := range nfs {
				nf.Items[code].Cfop = product.Cfop
				nf.Items[code].Desc = product.DescProd
				nf.Items[code].Ncm = product.Ncm
				nf.Items[code].Unmedida = product.UnMedida
			}
		}

		idnfe := genKey()

		locker.Lock()
		tempNfs[idnfe] = nfs
		locker.Unlock()
		go func(idnfe string) {
			time.Sleep(time.Minute * 5)
			locker.Lock()
			delete(tempNfs, idnfe)
			locker.Unlock()
		}(idnfe)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("access-control-expose-headers", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseExcel(nfs, stocks, idnfe))
	}
}
