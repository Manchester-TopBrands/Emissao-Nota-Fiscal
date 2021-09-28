package models

//Modalidade ...
var Modalidade map[int]string = map[int]string{
	0: "ContratacaoPorContaDoRemetente",
	1: "ContratacaoPorContaDoDestinatario",
	2: "ContratacaoPorContaDeTerceiros",
	3: "TransporteProprioPorContaDoRemetente",
	4: "TransporteProprioPorContaDoDestinatario",
	9: "SemOcorrenciaDeTransporte",
}

// NfeStruct ...
type NfeStruct struct {
	Tipo                  string            `json:"tipo,omitempty"`
	ID                    string            `json:"id,omitempty"`
	AmbienteEmissao       string            `json:"ambienteEmissao,omitempty"`
	NaturezaOperacao      string            `json:"naturezaOperacao,omitempty"`
	TipoOperacao          string            `json:"tipoOperacao,omitempty"`
	Finalidade            string            `json:"finalidade,omitempty"`
	ConsumidorFinal       bool              `json:"consumidorFinal,omitempty"`
	EnviarPorEmail        bool              `json:"enviarPorEmail,omitempty"`
	Cliente               *ClientCadastro   `json:"cliente,omitempty"`
	Itens                 []ProdutoCadastro `json:"itens,omitempty"`
	InformacoesAdicionais string            `json:"informacoesAdicionais,omitempty"`
	Transporte            struct {
		Frete struct {
			Modalidade string `json:"modalidade,omitempty"`
		} `json:"frete,omitempty"`
		Volume struct {
			Quantidade int     `json:"quantidade,omitempty"`
			Especie    string  `json:"especie,omitempty"`
			Numeracao  string  `json:"numeracao,omitempty"`
			PesoBruto  float64 `json:"pesoBruto,omitempty"`
			PesoLiq    float64 `json:"pesoLiquido,omitempty"`
		} `json:"volume,omitempty"`
	} `json:"transporte,omitempty"`
} // modelo para solicitar um pedido de emissão de NFE pela api E-notas

// ClientCadastro ...
type ClientCadastro struct {
	TipoPessoa                string `json:"tipoPessoa,omitempty"`
	IndicadorContribuinteICMS string `json:"indicadorContribuinteICMS,omitempty"`
	InscricaoEstadual         string `json:"inscricaoEstadual,omitempty"`
	Nome                      string `json:"nome,omitempty"`
	Email                     string `json:"email,omitempty"`
	Telefone                  string `json:"telefone,omitempty"`
	CpfCnpj                   string `json:"cpfCnpj,omitempty"`
	Endereco                  struct {
		UF          string `json:"uf,omitempty"`
		Cidade      string `json:"cidade,omitempty"`
		Logradouro  string `json:"logradouro,omitempty"`
		Numero      string `json:"numero,omitempty"`
		Complemento string `json:"complemento,omitempty"`
		Bairro      string `json:"bairro,omitempty"`
		CEP         string `json:"cep,omitempty"`
	} `json:"endereco,omitempty"`
} // modelo cliente pedido NFE

// ProdutoCadastro ...
type ProdutoCadastro struct {
	Cfop          string  `json:"cfop,omitempty"`
	Codigo        string  `json:"codigo,omitempty"`
	Descricao     string  `json:"descricao,omitempty"`
	Ncm           string  `json:"ncm,omitempty"`
	Quantidade    float64 `json:"quantidade,omitempty"`
	Unidademedida string  `json:"unidadeMedida,omitempty"`
	ValorUnitario float64 `json:"valorUnitario,omitempty"`
	Frete         float64 `json:"frete,omitempty"`
	Impostos      struct {
		Icms struct {
			SituacaoTributaria string  `json:"situacaoTributaria,omitempty"`
			Origem             string  `json:"origem,omitempty"`
			Aliquota           float64 `json:"aliquota,omitempty"`
		} `json:"icms,omitempty"`
		Pis struct {
			SituacaoTributaria string `json:"situacaoTributaria,omitempty"`
			PorAliquota        struct {
				Aliquota float64 `json:"aliquota,omitempty"`
			} `json:"porAliquota,omitempty"`
		} `json:"pis,omitempty"`
		Cofins struct {
			SituacaoTributaria string `json:"situacaoTributaria,omitempty"`
			PorAliquota        struct {
				Aliquota float64 `json:"aliquota,omitempty"`
			} `json:"porAliquota,omitempty"`
		} `json:"cofins,omitempty"`
	} `json:"impostos,omitempty"`
} // modelo Items pedido NFE

// respEnotas ...
