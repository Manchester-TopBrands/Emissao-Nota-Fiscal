package models

// StructExcel ...
type StructExcel struct {
	ID       string             `json:"id"`
	Clientes []string           `json:"clientes"`
	Items    []*StructExcelItem `json:"items"`
}

// StructExcelItem ...
type StructExcelItem struct {
	CodeDesc string    `json:"cod-Desc,omitempty"`
	Price    float64   `json:"valor,omitempty"`
	Qties    []float64 `json:"qtds,omitempty"`
}
