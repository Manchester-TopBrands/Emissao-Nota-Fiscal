package models

import (
	"time"
)

//ProductInfo ...
type ProductInfo struct {
	Qtd      float64
	Cfop     string
	DescProd string
	Ncm      string
	UnMedida string
}

//XmlFormat ...
type XmlFormat struct {
	NFe struct {
		InfNFe InfNFe `xml:"infNFe"`
	} `xml:"NFe"`
}

// Ide ...
type Ide struct {
	Cuf      string    `xml:"cUF"`
	Cnf      string    `xml:"cNF"`
	Natop    string    `xml:"natOp"`
	Mod      string    `xml:"mod"`
	Serie    string    `xml:"serie"`
	Nnf      string    `xml:"nNF"`
	Dhemi    time.Time `xml:"dhEmi"`
	Dhsaient time.Time `xml:"dhSaiEnt"`
	Tpnf     string    `xml:"tpNF"`
	Iddest   string    `xml:"idDest"`
	Cmunfg   string    `xml:"cMunFG"`
	Tpimp    string    `xml:"tpImp"`
	Tpemis   string    `xml:"tpEmis"`
	Cdv      string    `xml:"cDV"`
	Tpamb    string    `xml:"tpAmb"`
	Finnfe   string    `xml:"finNFe"`
	Indfinal string    `xml:"indFinal"`
	Indpres  string    `xml:"indPres"`
	Procemi  string    `xml:"procEmi"`
	Verproc  string    `xml:"verProc"`
	Nfref    struct {
		Refnfe string `xml:"refNFe"`
	} `xml:"NFref"`
}

// Prod ...
type Prod struct {
	Cprod       int     `xml:"cProd"`
	CEAN        string  `xml:"cEAN"`
	NCM         string  `xml:"NCM,omitempty"`
	DescProduto string  `xml:"xProd,omitempty"`
	Cfop        string  `xml:"CFOP"`
	Ucom        string  `xml:"uCom"`
	Quantidade  float64 `xml:"qCom,omitempty"`
	ValorUni    float64 `xml:"vUnCom"`
	Vprod       float64 `xml:"vProd"`
	CEantrib    string  `xml:"cEANTrib"`
	Utrib       string  `xml:"uTrib"`
	Qtrib       float64 `xml:"qTrib"`
	Vuntrib     float64 `xml:"vUnTrib"`
	Indtot      string  `xml:"indTot"`
}

// Det ...
type Det struct {
	Prod    Prod  `xml:"prod"`
	Imposto Impos `xml:"imposto"`
	IDProd  int   `xml:"nItem,attr"`
}

// Impos ...
type Impos struct {
	Icms struct {
		Icms00 Icms00 `xml:"ICMS00"`
	} `xml:"ICMS"`
	Ipi struct {
		Cenq    string  `xml:"cEnq"`
		Ipitrib IpiTrib `xml:"IPITrib"`
	} `xml:"IPI"`
	Pis struct {
		PisOutr PisOutr `xml:"PISOutr"`
	} `xml:"PIS"`
	Cofins struct {
		CofinsOutr CofinsOutr `xml:"COFINSOutr"`
	} `xml:"COFINS"`
}

// PisOutr ...
type PisOutr struct {
	Cst  string  `xml:"CST"`
	Vbc  float64 `xml:"vBC"`
	Ppis float64 `xml:"pPIS"`
	Vpis float64 `xml:"vPIS"`
}

// IpiTrib ...
type IpiTrib struct {
	Cst  string  `xml:"CST"`
	Vbc  float64 `xml:"vBC"`
	Pipi float64 `xml:"pIPI"`
	Vipi float64 `xml:"vIPI"`
}

// CofinsOutr ...
type CofinsOutr struct {
	Cst      string  `xml:"CST"`
	Vbc      float64 `xml:"vBC"`
	Pconfins float64 `xml:"pCONFINS"`
	Vconfins float64 `xml:"vCONFINS"`
}

// Icms00 ...
type Icms00 struct {
	Orig  string  `xml:"orig"`
	Cst   string  `xml:"CST"`
	Modbc string  `xml:"modBC"`
	Vbc   float64 `xml:"vBC"`
	Picms float64 `xml:"pICMS"`
	Vicms float64 `xml:"vICMS"`
}

// Enderdest ...
type Enderdest struct {
	Xlgr    string `xml:"xLgr"`
	Nro     string `xml:"nro"`
	Xbairro string `xml:"xBairro"`
	Cmun    string `xml:"cMun"`
	Xmun    string `xml:"Xmun"`
	Uf      string `xml:"UF"`
	Cep     string `xml:"CEP"`
	Cpais   string `xml:"cPais"`
	Xpais   string `xml:"xPais"`
}

// Emit ...
type Emit struct {
	Cnpj      string    `xml:"CNPJ"`
	Xnome     string    `xml:"xNome"`
	Enderdest Enderdest `xml:"enderDest"`
	Indiedest string    `xml:"indIEDest"`
	Ie        string    `xml:"IE"`
}

// InfNFE ...
type InfNFe struct {
	Det  []Det  `xml:"det"`
	Ide  Ide    `xml:"ide"`
	Emit Emit   `xml:"emit"`
	ID   string `xml:"Id,attr"`
}
