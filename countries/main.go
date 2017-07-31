package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type NameInfo struct {
	Common   string `json:"common"`
	Official string `json:"official"`
}

type NameNode struct {
	Common   string               `json:"common"`
	Official string               `json:"official"`
	Native   map[string]*NameInfo `json:"native"`
}

type CountryInfo struct {
	Name         NameInfo             `json:"name"`
	Tld          []string             `json:"tld"`
	Cca2         string               `json:"cca2"`
	Ccn3         string               `json:"ccn3"`
	Cca3         string               `json:"cca3"`
	Cioc         string               `json:"cioc"`
	Currency     []string             `json:"currency"`
	CallingCode  []string             `json:"callingCode"`
	Capital      string               `json:"capital"`
	AltSpellings []string             `json:"altSpellings"`
	Region       string               `json:"region"`
	Subregion    string               `json:"subregion"`
	Languages    map[string]string    `json:"languages"`
	Translations map[string]*NameInfo `json:"translations"`
	Latlng       []float64            `json:"latlng"`
	Demonym      string               `json:"demonym"`
	Landlocked   bool                 `json:"landlocked"`
	Borders      []string             `json:"borders"`
	Area         float64              `json:"area"`
}

func main() {
	data, _ := ioutil.ReadFile("countries.json")
	var items []CountryInfo
	err := json.Unmarshal(data, &items)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(items)
	fmt.Println(len(items))
}
