package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/nyaruka/phonenumbers"
)

var (
	loc   = flag.String("loc", "US", "two letter country code")
	phone = flag.String("phone", "16502530000", "phone number")
)

var names = [...]string{
	"FIXED_LINE",
	"MOBILE",
	"FIXED_LINE_OR_MOBILE",
	"TOLL_FREE",
	"PREMIUM_RATE",
	"SHARED_COST",
	"VOIP",
	"PERSONAL_NUMBER",
	"PAGER",
	"UAN",
	"VOICEMAIL",
	"UNKNOWN",
}

func main() {
	flag.Parse()

	number, err := phonenumbers.Parse(*phone, *loc)
	if err != nil {
		log.Fatal(err, number)
	}
	numberType := phonenumbers.GetNumberType(number)
	fmt.Println(names[numberType])
}
