package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
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
	capitalCity string `json:"capitalCity"`
	longitude   string `json:"longitude"`
	latitude    string `json:"latitude"`
}

type _geo_country struct {
	_geo_country_core
	Region      _geo_country_region `json: "region"`
	Adminregion _geo_country_region `json::"adminregion"`
	IncomeLevel _geo_country_region `json:"incomeLevel"`
	LendingType _geo_country_region `json:"lendingType"`
}

type _c_map struct {
	Letter string               `json:"letter"`
	Data   []*_geo_country_core `json:"data"`
}

func main() {
	countries := getCountryList()
	fmt.Println(*countries[0])
	cntCores := []_geo_country_core{}
	m := make(map[string][]*_geo_country_core)
	for _, v := range countries {
		cntCores = append(cntCores, v._geo_country_core)
		hc := []rune(v._geo_country_core.Name)[0]
		//if _, ok := m[string(hc)]; ok {
		m[string(hc)] = append(m[string(hc)], &v._geo_country_core)
		//}
	}

	/* b, _ := json.Marshal(cntCores)
	d1 := []byte(string(b))
	ioutil.WriteFile("country.out.json", d1, 0644) */

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var cmapList []*_c_map
	for _, k := range keys {
		list := m[k]
		cmap := &_c_map{Letter: k, Data: list}
		cmapList = append(cmapList, cmap)
	}

	b1, _ := json.Marshal(cmapList)
	d2 := []byte(string(b1))
	ioutil.WriteFile("pageformat.country.json", d2, 0644)

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
