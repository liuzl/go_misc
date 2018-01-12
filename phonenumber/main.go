package main

import (
	"flag"
	"github.com/GeertJohan/go.rice"
	"github.com/golang/glog"
	"github.com/liuzl/gocountries"
	"github.com/liuzl/goutil/rest"
	"github.com/liuzl/phonenumbers"
	"net/http"
	"net/http/httputil"
)

var (
	serverAddr = flag.String("addr", ":8080", "bind address")
)

type NumberInfo struct {
	*phonenumbers.PhoneNumber
	IsValidNumber bool                         `json:"is_valid_number"`
	NumberType    phonenumbers.PhoneNumberType `json:"number_type"`
	E164          string                       `json:"e164"`
	RegionCode    string                       `json:"region_code"`
	Location      map[string]string            `json:"location"`
	TimeZones     []string                     `json:"time_zones"`
	Carrier       map[string]string            `json:"carrier"`
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
	number, err := phonenumbers.Parse(phone, cc)
	//number, err := phonenumbers.ParseAndKeepRawInput(phone, cc)
	if err != nil {
		rest.MustEncode(w, struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}{Status: "error", Message: err.Error()})
	} else {
		info := &NumberInfo{PhoneNumber: number}
		info.IsValidNumber = phonenumbers.IsValidNumber(number)
		info.NumberType = phonenumbers.GetNumberType(number)
		info.E164 = phonenumbers.Format(number, phonenumbers.E164)
		info.RegionCode = phonenumbers.GetRegionCodeForNumber(number)
		country := gocountries.FindCountryByAlpha(info.RegionCode)
		lang := ""
		if country != nil && len(country.Languages) > 0 {
			lang = country.Languages[0]
		}
		println(lang)
		info.Location = make(map[string]string)
		info.Location["en"], _ = phonenumbers.GetGeocodingForNumber(number, "en")
		if lang != "" {
			tmp, _ := phonenumbers.GetGeocodingForNumber(number, lang)
			if tmp != "" && tmp != info.Location["en"] {
				info.Location[lang] = tmp
			}
		}
		info.TimeZones, _ = phonenumbers.GetTimezonesForNumber(number)
		info.Carrier = make(map[string]string)
		info.Carrier["en"], _ = phonenumbers.GetCarrierForNumber(number, "en")
		if lang != "" {
			tmp, _ := phonenumbers.GetCarrierForNumber(number, lang)
			if tmp != "" && tmp != info.Carrier["en"] {
				info.Carrier[lang] = tmp
			}
		}
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
