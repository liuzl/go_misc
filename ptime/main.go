package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/tkuchiki/parsetime"
	"github.com/ttacon/libphonenumber"
	"log"
)

var (
	loc = flag.String("loc", "US", "two letter country code")
	str = flag.String("str", "", "time to parse")
)

func main() {
	flag.Parse()
	cc := libphonenumber.GetCountryCodeForRegion(*loc)
	tz, ok := libphonenumber.CountryCodeToTimeZones[cc]
	if ok {
		for _, t := range tz {
			fmt.Println(t)
			p, err := parsetime.NewParseTime(t)
			if err != nil {
				log.Fatal(err)
			}
			t, err := p.Parse(*str)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(t)
			j, _ := json.Marshal(t)
			fmt.Println(string(j))
		}
	}
}
