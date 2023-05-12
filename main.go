package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
)

type Book struct {
	Name           string
	Desc           string
	Picture        string
	LinkOnOzon     string
	Characteristic BookCharacteristic
}

type BookCharacteristic struct {
	PagesCount   string
	Cover        string
	Age          string
	Series       string
	Types        string
	Size         string
	VendorCode   string
	Weight       string
	Illustration string
}

func main() {

	b := &Book{}

	c := colly.NewCollector(
		colly.AllowedDomains("rosman.ru"),
	)

	c.OnHTML(".item-line-first", func(e *colly.HTMLElement) {
		b.Name = e.ChildText("h1")
	})
	c.OnHTML(".item-line-left", func(e *colly.HTMLElement) {
		b.Desc = e.ChildText(".detailtext p")
		b.Picture = e.ChildAttr(".item-pic img", "src")
	})
	c.OnHTML(".marketplace-block", func(e *colly.HTMLElement) {
		b.LinkOnOzon = e.ChildAttr(".ozon", "href")
	})
	c.OnHTML(".tab-content #chars", func(e *colly.HTMLElement) {
		e.ForEach("dl", func(i int, el *colly.HTMLElement) {
			b.Characteristic.PagesCount = el.ChildText("dd:nth-child(2)")
			b.Characteristic.Cover = el.ChildText("dd:nth-child(6)")
			b.Characteristic.Age = el.ChildText("dd:nth-child(10)")
			b.Characteristic.Series = el.ChildText("dd:nth-child(14) a")
			b.Characteristic.Types = el.ChildText("dd:nth-child(18)")
			b.Characteristic.Size = el.ChildText("dd:nth-child(22)")
			b.Characteristic.VendorCode = el.ChildText("dd:nth-child(26)")
			b.Characteristic.Weight = el.ChildText("dd:nth-child(30)")
			b.Characteristic.Illustration = el.ChildText("dd:nth-child(34)")
		})
	})
	c.Visit("https://rosman.ru/catalog/item/luchshie-skazki-na-noch/")

	data, err := json.Marshal(b)
	if err != nil {
		fmt.Println("JSON marshaling failed: %s", err)
	}

	//fmt.Println(b)
	fmt.Println(string(data))
}
