package main

import (
	"encoding/base64"
	"fmt"
	"github.com/scusi/epc"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

func ibanIsValid(iban string) bool {
	ibanRE := regexp.MustCompile(`^[A-Z]{2}[0-9A-Z]{2,32}$`)
	return ibanRE.MatchString(iban)
}

func nameIsValid(name string) bool {
	nameRE := regexp.MustCompile(`^[0-9A-Za-z\s\?\-\:\(\)\.\,\+\'üäöÜÄÖß]{2,70}$`)
	return nameRE.MatchString(name)
}

func subjectIsValid(subject string) bool {
	subjectRE := regexp.MustCompile(`^[0-9A-Za-z\s\?\-\:\(\)\.\,\+\'üäöÜÄÖß]{1,140}$`)
	return subjectRE.MatchString(subject)
}

func BICIsValid(bic string) bool {
	bicRE := regexp.MustCompile(`(?m)([a-zA-Z]{4})([a-zA-Z]{2})(([2-9a-zA-Z]{1})([0-9a-np-zA-NP-Z]{1}))((([0-9a-wy-zA-WY-Z]{1})([0-9a-zA-Z]{2}))|([xX]{3})|)`)
	return bicRE.MatchString(bic)
}

func urlparam2pD(r *http.Request) (pageData map[string]string) {
	pageData = make(map[string]string)
	values := r.URL.Query()
	for k, v := range values {
		pageData[k] = v[0]
		if debug {
			log.Printf("read parameter %s : %s", k, v)
		}
	}
	up := pD2URLparam(pageData)
	pageData["epcurl"] = up
	return
}

func pD2URLparam(pageData map[string]string) (up string) {
	// encode URL values
	uv := url.Values{}
	for k, v := range pageData {
		uv.Add(k, v)
	}
	up = uv.Encode()
	return
}

func pD2epc(pageData map[string]string) (e *epc.EPC, err error) {
	amount := 0.0
	if len(pageData["epcamount"]) > 0 {
		amount, err = strconv.ParseFloat(pageData["epcamount"], 64)
		if err != nil {
			return
		}
	}
	if ibanIsValid(pageData["epciban"]) != true {
		err = fmt.Errorf("IBAN is invalid")
		return
	}
	if nameIsValid(pageData["epcname"]) != true {
		err = fmt.Errorf("Name is invalid")
		return
	}
	if subjectIsValid(pageData["epcsubject"]) != true {
		err = fmt.Errorf("Subject is invalid")
		return
	}
	if len(pageData["epcbic"]) > 0 {
		if BICIsValid(pageData["epcbic"]) != true {
			err = fmt.Errorf("BIC is invalid")
			return
		}
		e, err = epc.NewWithBIC(
			pageData["epcbic"],
			pageData["epcname"],
			pageData["epciban"],
			pageData["epcsubject"],
			amount,
		)
	} else {
		e, err = epc.New(
			pageData["epcname"],
			pageData["epciban"],
			pageData["epcsubject"],
			amount,
		)
	}
	return
}

func epc2b64QR(e *epc.EPC) (b64QR string, err error) {
	qr, err := e.MarshalQR()
	if err != nil {
		return
	}
	b64QR = base64.StdEncoding.EncodeToString(qr)
	return
}
