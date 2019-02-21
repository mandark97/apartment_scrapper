package main

import (
	"fmt"
	"regexp"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

type Apartment struct {
	OfferedBy          string
	Surface            int
	Partitioning       string
	YearOfConstruction string
	Floor              string
	NoRooms            int
	Description        string
	Images             []string
	Price              int
}

const startingPoint = "https://www.olx.ro/"

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.olx.ro"),
		colly.Debugger(&debug.LogDebugger{}),
		colly.Async(true),
		colly.URLFilters(
			regexp.MustCompile("http://httpbin\\.org/(|e.+)$"),
			regexp.MustCompile("http://httpbin\\.org/h.+"),
		),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 10})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)

	})
	fmt.Println("starting")
	c.Visit(startingPoint)

}
