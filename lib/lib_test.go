package lib_test

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/uelei/go_dolar_get/lib"
)

func Test_Check(t *testing.T) {
	t.Run("check error lib", func(t *testing.T) {
		// If the function panics, recover() will
		// return a non nil value.
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("function should panic")
			}
		}()

		e := errors.New("error")
		lib.Check(e)
	})

	t.Run("Check no error", func(t *testing.T) {
		lib.Check(nil)
	})
}

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

func TextFormatDateToBC(t *testing.T) {
	date := time.Date(2022, 1, 2, 10, 20, 11, 0, time.UTC)

	result := lib.FormatDateToBC(date)

	if result != "2022-01-02" {
		t.Fatal("error converting date")
	}
}
