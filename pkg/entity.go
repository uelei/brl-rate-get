package brlrateget

type DayValues struct {
	CotacaoCompra   float64 `json:"cotacaoCompra"`
	CotacaoVenda    float64 `json:"cotacaoVenda"`
	DataHoraCotacao string  `json:"dataHoraCotacao"`
}

type DayValuesResponse struct {
	ParidadeCompra  float64 `json:"paridadeCompra"`
	ParidadeVenda   float64 `json:"paridadeVenda"`
	CotacaoCompra   float64 `json:"cotacaoCompra"`
	CotacaoVenda    float64 `json:"cotacaoVenda"`
	DataHoraCotacao string  `json:"dataHoraCotacao"`
	TipoBoletim     string  `json:"tipoBoletim"`
}

type RangeResponse struct {
	OdataContext string              `json:"@odata.context"`
	Value        []DayValuesResponse `json:"value"`
}
