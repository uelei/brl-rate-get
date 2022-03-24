package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func Check(e error) {
	if e != nil {
		// log.Fatalln(e)
		panic(e)

	}
}

type DayValues struct {
	CotacaoCompra   float64 `json:"cotacaoCompra"`
	CotacaoVenda    float64 `json:"cotacaoVenda"`
	DataHoraCotacao string  `json:"dataHoraCotacao"`
}

type RangeResponse struct {
	OdataContext string      `json:"@odata.context"`
	Value        []DayValues `json:"value"`
}

func WriteResultsToCSV(results []DayValues, filename string) {

	log.Printf("Writing a file.")
	os.Remove(filename)

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	Check(err)

	datawriter := bufio.NewWriter(file)

	header := "Day,Buy,Sell\n"

	_, _ = datawriter.WriteString(header)
	for _, data := range results {
		datecut := data.DataHoraCotacao[:10]
		res := fmt.Sprintf("%s,%0.4f,%0.4f\n", datecut, data.CotacaoCompra, data.CotacaoVenda)

		_, _ = datawriter.WriteString(res)
	}

	datawriter.Flush()
	file.Close()
}

func FormatDateToBC(date time.Time) string {

	return date.Format("01-02-2006")

}
