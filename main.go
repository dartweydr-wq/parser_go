package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

type books struct {
	nameBook        string
	descBook        string
	pictureBook     string
	linkOnOzonBook  string
	characteristics []interface{}
}

type characteristicsBooks struct {
	countPagesBook int
	coverBook      string
	ageBook        string
	seriesBook     string
	typeBook       string
	sizeBook       string
	vendorCode     string
	weight         string
	illustration   string
}

func main() {

	b := books{}

	c := colly.NewCollector(
		colly.AllowedDomains("rosman.ru"),
	)

	// Find and print all links
	c.OnHTML(".item-line-first", func(e *colly.HTMLElement) {
		b.nameBook = e.ChildText("h1")
	})
	c.OnHTML(".item-line-left", func(e *colly.HTMLElement) {
		b.descBook = e.ChildText(".detailtext p")
		b.pictureBook = e.ChildAttr(".item-pic img", "src")
	})
	c.OnHTML(".marketplace-block", func(e *colly.HTMLElement) {
		b.linkOnOzonBook = e.ChildAttr(".ozon", "href")
	})
	c.OnHTML(".tab-content #chars", func(e *colly.HTMLElement) {
		e.ForEach("dd", func(_ int, el *colly.HTMLElement) {
			b.characteristics = append(b.characteristics, el.Text)
		})
	})
	c.Visit("https://rosman.ru/catalog/item/luchshie-skazki-na-noch/")

	/*data, err := json.Marshal(&b)
	if err != nil {
		fmt.Println("JSON marshaling failed: %s", err)
	}*/

	fmt.Println(b)
	//fmt.Println(data)
}
