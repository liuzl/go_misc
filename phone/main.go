package main

import (
	"flag"
	"fmt"
	"github.com/ttacon/libphonenumber"
	"log"
)

var (
	loc = flag.String("loc", "US", "two letter country code")
)

func main() {
	flag.Parse()
	//number, err := libphonenumber.Parse("16502530000", "US")

	number, err := libphonenumber.ParseAndKeepRawInput("16502530000", "US")
	if err != nil {
		log.Fatal(err, number)
	}
	cc := libphonenumber.GetCountryCodeForRegion(*loc)
	tz, ok := libphonenumber.CountryCodeToTimeZones[cc]
	if ok {
		fmt.Println(tz)
	}
}
