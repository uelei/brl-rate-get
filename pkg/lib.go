package brlrateget

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func GenerateYesterdayString() time.Time {

	currentTime := time.Now()

	// Calculate yesterday's date
	yesterday := currentTime.AddDate(0, 0, -4)

	return yesterday
}

func WriteResultsToCSV(results []DayValues, filename string) {
	log.Printf("Writing a file.")
	os.Remove(filename)

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	datawriter := bufio.NewWriter(file)

	header := "Day,Buy,Sell\n"

	_, _ = datawriter.WriteString(header)
	for _, data := range results {
		datecut := strings.Replace(data.DataHoraCotacao[:10], "-", "", -1)
		res := fmt.Sprintf("%s,%0.4f,%0.4f\n", datecut, data.CotacaoCompra, data.CotacaoVenda)
		_, _ = datawriter.WriteString(res)
	}

	datawriter.Flush()
	file.Close()
}

func FormatDateToBC(date time.Time) string {
	return date.Format("01-02-2006")
}

func ParseStringToDate(dateString string) time.Time {

	date, error := time.Parse("2006-01-02", dateString)

	if error != nil {
		fmt.Println(error)
		panic(error)
	}
	return date
}

func PadNumberWithZero(value float64) string {
	return fmt.Sprintf("%.4f", value)
}

func FormatGetResult(rates []DayValues, currency string) {
	if len(rates) == 0 {
		fmt.Println("No rates found for:", currency)
		return

	}

	rate := rates[len(rates)-1]

	fmt.Println("")
	fmt.Println("____________________________________")
	fmt.Println("|#######====================#######|")
	fmt.Println("|#  BRL         ", currency, "             #|")
	fmt.Println("|#         ====                   #|")
	fmt.Println("|#               Buy:  ", PadNumberWithZero(rate.CotacaoCompra), "   #|")
	fmt.Println("|#  1,00         Sell: ", PadNumberWithZero(rate.CotacaoVenda), "   #|")
	fmt.Println("|#                                #|")
	fmt.Println("|##==============================##|")
	fmt.Println("|    Rate Date  ", rate.DataHoraCotacao[:10], "       |")
	fmt.Println(" ---------------------------------- ")

}
