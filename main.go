package main

import (
	"encoding/json"
	"io"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/eufelipemateus/map-retrieve-countries/utils"
)

type Country struct {
	Name string `json:"name" `
	ID   string `json:"id" `
}

func main() {
	f, err := os.Open("./world.svg")

	utils.Check(err)

	defer f.Close()

	dat, err := io.ReadAll(f)
	utils.Check(err)

	r := strings.NewReader(string(dat))

	doc, err := goquery.NewDocumentFromReader(r)
	utils.Check(err)

	paths := doc.Find("path")

	total := paths.Length()

	var countrs []Country

	for i := 0; i < total; i++ {
		attrs := paths.Get(i).Attr
		var id, title string

		for a := range attrs {
			if attrs[a].Key == "id" {
				id = attrs[a].Val
			}
			if attrs[a].Key == "title" {
				title = attrs[a].Val
			}

		}
		newCountry := Country{ID: id, Name: title}

		countrs = append(countrs, newCountry)

	}

	file, _ := json.MarshalIndent(countrs, "", " ")

	_ = os.WriteFile("country_list.json", file, 0644)
}
