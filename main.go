package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/uelei/brl-rate-get/pkg"
)

func handle_get(currency string) {

	yesterday := brlrateget.GenerateYesterdayString()

	currentTime := time.Now()

	rates := brlrateget.MakeRangeRequest(yesterday, currentTime, strings.ToUpper(currency), false)

	brlrateget.FormatGetResult(rates, currency)
}

func handle_range(args []string) {

	if len(args) < 3 {

		fmt.Println("You must pass a currency, start_date and a end_date after the command")
		fmt.Println(" brl-rate-get -range usd 2023-10-01 2023-10-31")
		os.Exit(0)

	}

	start_date := brlrateget.ParseStringToDate(args[1])
	end_date := brlrateget.ParseStringToDate(args[2])

	currency := args[0]

	rates := brlrateget.MakeRangeRequest(start_date, end_date, strings.ToUpper(currency), true)
	filename := fmt.Sprintf("%s_brl_%s-%s.csv", currency,
		start_date.Format("2006-01-02"), end_date.Format("2006-01-02"))
	brlrateget.WriteResultsToCSV(rates, filename)
	fmt.Println("File saved ", filename)
}

func main() {

	getCmd := flag.Bool("get", false, "To get yesterday closing currency")
	rangeCmd := flag.Bool("range", false, "To get a range of closing currency rates")

	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("You must pass a currency after the command")
		os.Exit(0)
	}

	switch {
	case *rangeCmd:
		handle_range(flag.Args())
	case *getCmd:
		handle_get(flag.Args()[0])
	default:
		fmt.Println("You must pass a command")
	}

}
