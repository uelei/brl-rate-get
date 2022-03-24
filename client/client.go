package client

import (
	"encoding/json"
	"fmt"
	"github.com/uelei/go_dolar_get/lib"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

// Get
func GetDataFromUrl(url string) ([]byte, error) {

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	response, err := Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

func MakeRangeRequest(startDate time.Time, endDate time.Time) []lib.DayValues {

	formatedUrl := fmt.Sprintf("https://olinda.bcb.gov.br/olinda/servico/PTAX/versao/v1/odata/CotacaoDolarPeriodo(dataInicial=@dataInicial,dataFinalCotacao=@dataFinalCotacao)?@dataInicial='%s'&@dataFinalCotacao='%s'&$top=1000&$format=json&$select=cotacaoCompra,cotacaoVenda,dataHoraCotacao", lib.FormatDateToBC(startDate), lib.FormatDateToBC(endDate))

	resp, err := GetDataFromUrl(formatedUrl)

	lib.Check(err)

	var result lib.RangeResponse

	err := json.Unmarshal(resp, &result)

	lib.Check(err)

	log.Printf("Found %d records of price..\n", len(result.Value))

	return result.Value

}
