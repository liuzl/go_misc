package main

import (
	"flag"
	"github.com/GeertJohan/go.rice"
	"github.com/golang/glog"
	"github.com/liuzl/goutil/rest"
	"github.com/nyaruka/phonenumbers"
	"net/http"
	"net/http/httputil"
)

var (
	serverAddr = flag.String("addr", ":8080", "bind address")
)

type NumberInfo struct {
	*phonenumbers.PhoneNumber
	IsPossibleNumber bool                         `json:"is_possible_number"`
	IsValidNumber    bool                         `json:"is_valid_number"`
	NumberType       phonenumbers.PhoneNumberType `json:"number_type"`
	E164             string                       `json:"e164"`
	RegionCode       string                       `json:"region_code"`
	TimeZones        []string                     `json:"time_zones"`
	Carrier          string                       `json:"carrier"`
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		rest.MustEncode(w, struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{Status: "error", Message: err.Error()})
		return
	}
	w.Write(dump)
}

func PhoneHandler(w http.ResponseWriter, r *http.Request) {
	glog.Infof("addr=%s method=%s host=%s uri=%s",
		r.RemoteAddr, r.Method, r.Host, r.RequestURI)
	r.ParseForm()
	phone := r.FormValue("phone")
	cc := r.FormValue("cc")
	if cc == "" {
		cc = "ZZ"
	}
	number, err := phonenumbers.ParseAndKeepRawInput(phone, cc)
	if err != nil {
		rest.MustEncode(w, struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{Status: "error", Message: err.Error()})
	} else {
		info := &NumberInfo{PhoneNumber: number}
		info.IsPossibleNumber = phonenumbers.IsPossibleNumber(number)
		info.IsValidNumber = phonenumbers.IsValidNumber(number)
		info.NumberType = phonenumbers.GetNumberType(number)
		info.E164 = phonenumbers.Format(number, phonenumbers.E164)
		info.RegionCode = phonenumbers.GetRegionCodeForNumber(number)
		info.TimeZones, _ = phonenumbers.GetTimezonesForPrefix(info.E164)
		info.Carrier = number.GetPreferredDomesticCarrierCode()
		rest.MustEncode(w, info)
	}
}

func main() {
	flag.Parse()
	defer glog.Flush()
	defer glog.Info("server exit")
	http.HandleFunc("/echo/", EchoHandler)
	http.HandleFunc("/api/", PhoneHandler)
	http.Handle("/", http.FileServer(rice.MustFindBox("ui").HTTPBox()))
	glog.Info("server listen on", *serverAddr)
	glog.Error(http.ListenAndServe(*serverAddr, nil))
}
