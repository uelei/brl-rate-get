package brlrateget

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func format_url(base_url string, currency string, start_data string, end_data string) string {

	if base_url == "" {
		base_url = "https://olinda.bcb.gov.br/olinda/servico/PTAX/versao/v1/odata/CotacaoMoedaPeriodo"
	}
	// https://olinda.bcb.gov.br/olinda/servico/PTAX/versao/v1/swagger-ui3

	formatedUrl := fmt.Sprintf("%s(moeda='%s',dataInicial='%s',dataFinalCotacao='%s')?$top=1000", base_url, currency, start_data, end_data)
	return formatedUrl
}

func MakeRangeRequest(startDate time.Time, endDate time.Time, currency string) []DayValues {

	formatedUrl := format_url("", currency, FormatDateToBC(startDate), FormatDateToBC(endDate))
	resp, err := GetDataFromUrl(formatedUrl)

	if err != nil {
		panic(err)
	}
	var result RangeResponse

	if err := json.Unmarshal(resp, &result); err != nil {
		panic(err)
	}
	log.Printf("Found %d records of price..\n", len(result.Value))

	filtered := []DayValues{}

	for i := range result.Value {
		if result.Value[i].TipoBoletim == "Fechamento" {
			filtered = append(filtered, DayValues{CotacaoCompra: result.Value[i].CotacaoCompra, CotacaoVenda: result.Value[i].CotacaoVenda, DataHoraCotacao: result.Value[i].DataHoraCotacao})
		}
	}

	log.Printf("Found %d records of fechamento price..\n", len(filtered))
	return filtered
}
