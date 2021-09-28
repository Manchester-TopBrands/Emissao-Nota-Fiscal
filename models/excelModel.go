package models

// Nf ...
type Nf struct {
	ID           string
	Nf           string
	Name         string
	Cnpj         string
	Ie           string
	Contribuinte string
	Email        string
	Uf           string
	Cidade       string
	Logradouro   string
	Nr           string
	Complemento  string
	Bairro       string
	Cep          string
	AddInfo      string
	ModFrete     string
	QtdVol       int
	EspecieVol   string
	PesoLiq      float64
	PesoBruto    float64
	Items        map[string]*Item
}

// Item ...
type Item struct {
	Code     string
	Qty      float64
	Value    float64
	Cfop     string
	Desc     string
	Ncm      string
	Unmedida string
}
