package main

import (
	"flag"
	"fmt"
	"github.com/scusi/epc"
	"log"
)

var debug bool

//var out		string
var format string

var bic string
var iban string
var name string
var subject string
var ammount float64
var version int
var encoding int

func init() {
	flag.BoolVar(&debug, "debug", false, "prints debug messages, when true")
	flag.Float64Var(&ammount, "a", 123.42, "ammount to transfer")
	flag.StringVar(&bic, "b", "COBADEFFXXX", "BIC of the recipient")
	flag.StringVar(&iban, "i", "DE56120400000012262200", "IBAN of the recipient")
	flag.StringVar(&name, "n", "Florian Walther", "Name of the recipient")
	flag.StringVar(&subject, "s", "Test Überweisung", "subject of the transfer")
	flag.IntVar(&version, "v", 2, "version of EPC, can be 1 or 2")
	flag.IntVar(&encoding, "e", 1, "encoding used, 1=UTF-8 2=ISO8859-1")
	//flag.StringVar(&out,      "outfile", "-", "file to write output to")
	flag.StringVar(&format, "format", "text", "format to output, 'text', or 'png'")
}

func main() {
	flag.Parse()
	e := epc.NewWithBIC(
		bic,
		name,
		iban,
		subject,
		ammount,
	)
	if debug {
		log.Printf("%s", e)
	}
	switch format {
	case "text":
		fmt.Printf("%s", e)
	case "png":
		qr, err := e.MarshalQR()
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%s", qr)
	}

}
