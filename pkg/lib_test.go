package brlrateget_test

import (
	"os"
	"testing"
	"time"
	lib "github.com/uelei/brl-rate-get/pkg"
)

func TestWriteResultsToCSV(t *testing.T) {
	r := []lib.DayValues{}

	lib.WriteResultsToCSV(r, "test.csv")
	_, err := os.Lstat("test.csv")
	if err != nil {
		t.Fatal("error not saved file ")
	}

	dat, err := os.ReadFile("test.csv")
	if err != nil {
		t.Fatal("Error not ReadFile")
	}

	if string(dat) != "Day,Buy,Sell\n" {
		t.Fatal("file content invalid")
	}
}

func TestTextFormatDateToBC(t *testing.T) {
	date := time.Date(2022, 1, 2, 10, 20, 11, 0, time.UTC)

	result := lib.FormatDateToBC(date)
	if result != "01-02-2022" {
		t.Fatal("error converting date")
	}
}

func TestGenerateYesterdayString(t *testing.T) {

	result := lib.GenerateYesterdayString()

	currentTime := time.Now()

	if result.Day() != currentTime.Day()-1 {
		t.Fatal("error creating yesterday date")
	}
}

func TestParseStringToDate(t *testing.T) {

	date := lib.ParseStringToDate("2023-10-31")

	if date.Day() != 31 {
		t.Fatal("Error converting string to date Day")
	}

	if date.Month() != 10 {
		t.Fatal("Error converting string to date Month")
	}

	if date.Year() != 2023 {
		t.Fatal("Error converting string to date Year")
	}

}
