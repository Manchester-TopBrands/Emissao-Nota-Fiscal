package sql

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"strings"
	"xmlconvert/models"

	_ "github.com/denisenkom/go-mssqldb" //bblablalba
)

const (
	productFmt string = "INSERT INTO LINX_TBFG..NFE_%s_PRODUTOS (ID_NFE_%s,ID_PRODUTO,CODIGO_PRODUTO,CEAN,DESC_PROD,NCM,CFOP,UCOM,QCOM,VUNCOM,VPROD,CEANT_TRIB,UTRIB,QTRIB,VUN_TRIB,IND_TOT,ICMS00_ORIG,ICMS00_CST,ICMS00_MODBC,ICMS00_VBC,ICMS00_PICMS,ICMS00_VICMS,IPI_CENQ,IPI_TRIB_CST,IPI_TRIB_VBC,IPI_TRIB_PIPI,IPI_TRIB_VIPI,PIS_CST,PIS_VBC,PIS_PPIS,PIS_VPIS,COFINS_CST,COFINS_VBC,COFINS_PCOFINS,COFINS_VCOFINS) VALUES('%s','%d','%d','%s','%s','%s','%s','%s','%f','%f','%f','%s','%s','%f','%f','%s','%s','%s','%s','%f','%f','%f','%s','%s','%f','%f','%f','%s','%f','%f','%f','%s','%f','%f','%f')"
	nfFmt      string = "INSERT INTO LINX_TBFG..NFE_%s (ID_NFE_%s,NUM_SERIE_NFE,CNPJ_NFE,CLIFOR,NUM_NOTA,DATA_EMISSAO,QTD_ITEMS) VALUES('%s','%s','%s','%s','%s','%s','%d')"
	// saidaFmt   string = "INSERT INTO LINX_TBFG..NFE_SAIDAS (ID_NFE_SAIDAS,NUM_SERIE_NFE,CNPJ_NFE,NOME_CLIENTE,NUM_NOTA,DATA_EMISSAO,QTD_ITEMS) VALUES('%s','%s','%s','%s','%s','%s','%d')"
)

// SQLStr ...
type SQLStr struct {
	url *url.URL
	db  *sql.DB
}

// addNFF ...
func (s *SQLStr) addNFF(nfentrada *models.XmlFormat, kind string) error {
	if s.db.Ping() != nil {
		if err := s.connect(); err != nil {
			return err
		}
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf(nfFmt, kind, kind, nfentrada.NFe.InfNFe.ID, nfentrada.NFe.InfNFe.Ide.Serie, nfentrada.NFe.InfNFe.Emit.Cnpj, nfentrada.NFe.InfNFe.Emit.Xnome, nfentrada.NFe.InfNFe.Ide.Nnf, nfentrada.NFe.InfNFe.Ide.Dhsaient.Format("02-01-06 15:04:05"), len(nfentrada.NFe.InfNFe.Det))

	if _, err := tx.ExecContext(context.Background(), query); err != nil {
		tx.Rollback()
		return err
	}
	//fmt.Println(query)

	for _, det := range nfentrada.NFe.InfNFe.Det {
		if len(det.Prod.DescProduto) > 100 {
			det.Prod.DescProduto = det.Prod.DescProduto[:100]
		}
		query2 := fmt.Sprintf(productFmt, kind, kind, nfentrada.NFe.InfNFe.ID, det.IDProd, det.Prod.Cprod, det.Prod.CEAN, det.Prod.DescProduto, det.Prod.NCM, det.Prod.Cfop, det.Prod.Ucom, det.Prod.Quantidade, det.Prod.ValorUni, det.Prod.Vprod, det.Prod.CEantrib,
			det.Prod.Utrib, det.Prod.Qtrib, det.Prod.Vuntrib, det.Prod.Indtot, det.Imposto.Icms.Icms00.Orig, det.Imposto.Icms.Icms00.Cst, det.Imposto.Icms.Icms00.Modbc, det.Imposto.Icms.Icms00.Vbc, det.Imposto.Icms.Icms00.Picms, det.Imposto.Icms.Icms00.Vicms,
			det.Imposto.Ipi.Cenq, det.Imposto.Ipi.Ipitrib.Cst, det.Imposto.Ipi.Ipitrib.Vbc, det.Imposto.Ipi.Ipitrib.Pipi, det.Imposto.Ipi.Ipitrib.Vipi, det.Imposto.Pis.PisOutr.Cst, det.Imposto.Pis.PisOutr.Vbc, det.Imposto.Pis.PisOutr.Ppis, det.Imposto.Pis.PisOutr.Vpis, det.Imposto.Cofins.CofinsOutr.Cst, det.Imposto.Cofins.CofinsOutr.Vbc, det.Imposto.Cofins.CofinsOutr.Pconfins, det.Imposto.Cofins.CofinsOutr.Vconfins)
		_, err := tx.ExecContext(context.Background(), query2)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

// addNF ...
func (s *SQLStr) addNF(nfentrada *models.XmlFormat, kind string, username string) error {
	if s.db.Ping() != nil {
		if err := s.connect(); err != nil {
			return err
		}
	}

	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	query := fmt.Sprintf(nfFmt, kind, kind, nfentrada.NFe.InfNFe.ID, nfentrada.NFe.InfNFe.Ide.Serie, nfentrada.NFe.InfNFe.Emit.Cnpj, nfentrada.NFe.InfNFe.Emit.Xnome, nfentrada.NFe.InfNFe.Ide.Nnf, nfentrada.NFe.InfNFe.Ide.Dhsaient.Format("02-01-06 15:04:05"), len(nfentrada.NFe.InfNFe.Det))

	if _, err := tx.ExecContext(context.Background(), query); err != nil {
		tx.Rollback()
		return err
	}
	//fmt.Println(query)

	for _, det := range nfentrada.NFe.InfNFe.Det {
		if len(det.Prod.DescProduto) > 100 {
			det.Prod.DescProduto = det.Prod.DescProduto[:100]
		}
		query2 := fmt.Sprintf(productFmt, kind, kind, nfentrada.NFe.InfNFe.ID, det.IDProd, det.Prod.Cprod, det.Prod.CEAN, det.Prod.DescProduto, det.Prod.NCM, det.Prod.Cfop, det.Prod.Ucom, det.Prod.Quantidade, det.Prod.ValorUni, det.Prod.Vprod, det.Prod.CEantrib,
			det.Prod.Utrib, det.Prod.Qtrib, det.Prod.Vuntrib, det.Prod.Indtot, det.Imposto.Icms.Icms00.Orig, det.Imposto.Icms.Icms00.Cst, det.Imposto.Icms.Icms00.Modbc, det.Imposto.Icms.Icms00.Vbc, det.Imposto.Icms.Icms00.Picms, det.Imposto.Icms.Icms00.Vicms,
			det.Imposto.Ipi.Cenq, det.Imposto.Ipi.Ipitrib.Cst, det.Imposto.Ipi.Ipitrib.Vbc, det.Imposto.Ipi.Ipitrib.Pipi, det.Imposto.Ipi.Ipitrib.Vipi, det.Imposto.Pis.PisOutr.Cst, det.Imposto.Pis.PisOutr.Vbc, det.Imposto.Pis.PisOutr.Ppis, det.Imposto.Pis.PisOutr.Vpis, det.Imposto.Cofins.CofinsOutr.Cst, det.Imposto.Cofins.CofinsOutr.Vbc, det.Imposto.Cofins.CofinsOutr.Pconfins, det.Imposto.Cofins.CofinsOutr.Vconfins)
		_, err := tx.ExecContext(context.Background(), query2)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

// AddEntradas ...
func (s *SQLStr) AddEntradas(nfentrada *models.XmlFormat) error {
	return s.addNFF(nfentrada, "ENTRADAS")
}

// addSaidas ...
func (s *SQLStr) AddSaidas(nfentrada *models.XmlFormat, username string) error {
	return s.addNF(nfentrada, "SAIDAS", username)
}

// GetStcok ...
func (s *SQLStr) GetStock(codes []string) (map[string]*models.ProductInfo, error) {
	if s.db.Ping() != nil {
		if err := s.connect(); err != nil {
			return nil, err
		}
	}

	rows, err := s.db.QueryContext(context.Background(), fmt.Sprintf(`SELECT A.CODIGO_PRODUTO, A.QUANTIDADE_ESTOQUE , B.CFOP,B.DESC_PROD,B.NCM,B.UCOM FROM LINX_TBFG..ESTOQUE_PRODUTO A
	LEFT JOIN LINX_TBFG..NFE_ENTRADAS_PRODUTOS B ON A.CODIGO_PRODUTO = B.CODIGO_PRODUTO
	WHERE A.CODIGO_PRODUTO IN ('%s')`, strings.Join(codes, "', '")))
	if err != nil {
		return nil, err
	}
	rst := make(map[string]*models.ProductInfo)
	for rows.Next() {
		var code string
		info := new(models.ProductInfo)
		if err := rows.Scan(&code, &info.Qtd, &info.Cfop, &info.DescProd, &info.Ncm, &info.UnMedida); err != nil {
			return nil, err
		}
		rst[code] = info
	}
	return rst, nil
}

// MakeSQL ...
func MakeSQL(host, port, username, password string) (*SQLStr, error) {

	s := &SQLStr{}
	s.url = &url.URL{
		Scheme:   "sqlserver",
		User:     url.UserPassword(username, password),
		Host:     fmt.Sprintf("%s:%s", host, port),
		RawQuery: url.Values{}.Encode(),
	}
	return s, s.connect()
}

func (s *SQLStr) connect() error {
	var err error
	if s.db, err = sql.Open("sqlserver", s.url.String()); err != nil {
		return err
	}
	return s.db.PingContext(context.Background())
}
