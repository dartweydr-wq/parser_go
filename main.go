package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gocolly/colly"
)

type Book struct {
	Name         string
	Desc         string
	Picture      string
	LinkOnOzon   string
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
	flagName := flag.String("url", "", "")
	c := colly.NewCollector()
	flag.Parse()

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
			b.PagesCount = el.ChildText("dd:nth-child(2)")
			b.Cover = el.ChildText("dd:nth-child(6)")
			b.Age = el.ChildText("dd:nth-child(10)")
			b.Series = el.ChildText("dd:nth-child(14) a")
			b.Types = el.ChildText("dd:nth-child(18)")
			b.Size = el.ChildText("dd:nth-child(22)")
			b.VendorCode = el.ChildText("dd:nth-child(26)")
			b.Weight = el.ChildText("dd:nth-child(30)")
			b.Illustration = el.ChildText("dd:nth-child(34)")
		})
	})
	c.Visit(*flagName)

	data, err := json.Marshal(b)
	if err != nil {
		fmt.Println("JSON marshaling failed: %s", err)
	}

	fmt.Println(string(data))
}
