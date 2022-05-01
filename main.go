package main

import (
	"fmt"
	"github.com/alecthomas/kong"
	"github.com/uelei/go_dolar_get/client"
	lib "github.com/uelei/go_dolar_get/lib"
	"log"
	"time"
)

type Context struct {
	Debug bool
}

type RangeCmd struct {
	Start string `arg:""  name:"start_date" help:"date to start retrieve." type:"string"`
	End   string `arg:""  name:"end_date" help:"date to end retrieve." type:"string"`
}

func (r *RangeCmd) Run(ctx *Context) error {

	datestart, err := time.Parse("2006-01-02", r.Start)

	lib.Check(err)

	dateend, err := time.Parse("2006-01-02", r.End)

	lib.Check(err)

	log.Printf("Starting getting data from %s to %s \n", datestart, dateend)
	results := client.MakeRangeRequest(datestart, dateend)

	lib.WriteResultsToCSV(results, fmt.Sprintf("usd_brl_%s-%s.csv",
		datestart.Format("2006-01-02"), dateend.Format("2006-01-02")))

	log.Println("Done")
	return nil
}

var cli struct {
	Debug bool `help:"Enable debug mode."`

	Range RangeCmd `cmd:"" help:"Range of dates to query"`
}

func main() {
	ctx := kong.Parse(&cli)
	// Call the Run() method of the selected parsed command.
	err := ctx.Run(&Context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
}
