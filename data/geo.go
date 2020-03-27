package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//_geo_country_region ..
type _geo_country_region struct {
	ID       string `json:"id"`
	Iso2Code string `json:"iso2Code"`
	Value    string `json:"value"`
}
type _geo_country_core struct {
	ID          string `json:"id"`
	Iso2Code    string `json:"iso2Code"`
	Name        string `json:"name"`
	CapitalCity string `json:"capitalCity"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
}

type _geo_country struct {
	_geo_country_core
	Region      _geo_country_region `json: "region"`
	Adminregion _geo_country_region `json::"adminregion"`
	IncomeLevel _geo_country_region `json:"incomeLevel"`
	LendingType _geo_country_region `json:"lendingType"`
}

func main() {
	countries := getCountryList()
	fmt.Println(*countries[0])
	cntCores := []_geo_country_core{}
	for _, v := range countries {
		cntCores = append(cntCores, v._geo_country_core)
	}
	b, _ := json.Marshal(cntCores)

	d1 := []byte(string(b))
	ioutil.WriteFile("country.out.json", d1, 0644)
}

func getCountryList() []*_geo_country {
	b, err := ioutil.ReadFile("country.in.json")
	if err != nil {
		fmt.Print(err)
	}
	var countries []*_geo_country
	err = json.Unmarshal(b, &countries)

	return countries
}
