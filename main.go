package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

const startingPoint = "https://www.olx.ro/imobiliare/apartamente-garsoniere-de-vanzare/1-camera/bucuresti-ilfov-judet/"

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.olx.ro"),
		colly.URLFilters(
			regexp.MustCompile(startingPoint),
			regexp.MustCompile("https://www.olx.ro/imobiliare/apartamente-garsoniere-de-vanzare/1-camera/bucuresti-ilfov-judet/\\?page=(.)"),
		),
	)

	apartmentCollector := colly.NewCollector(
		colly.AllowedDomains("www.olx.ro"),
		//colly.Async(true),
		colly.URLFilters(
			regexp.MustCompile("https://www.olx.ro/oferta/([a-zA-Z-\\d.#]+)"),
		),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		apartmentCollector.Visit(e.Request.AbsoluteURL(link))
	})
	//c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 10})

	apartmentCollector.OnHTML(".offerbody", collectApartment)

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

func collectApartment(e *colly.HTMLElement) {
	apartment := Apartment{
		Title: e.ChildText("div.offer-titlebox > h1"),
		Price: e.ChildText(".price-label strong"),
		Link:  e.Request.URL.String(),
	}
	dateAdded, err := parseDate(e.ChildText("div.offer-titlebox em"))
	if err != nil {
		fmt.Println("Couldn't get announcement date")
	}

	apartment.DateAdded = dateAdded

	e.ForEach(".item", func(i int, element *colly.HTMLElement) {
		switch element.ChildText("tr th") {
		case "Oferit de":
			apartment.OfferedBy = element.ChildText(".value strong a")
		case "Compartimentare":
			apartment.Partitioning = element.ChildText(".value strong a")
		case "Suprafata utila":
			apartment.Surface, _ = strconv.Atoi(strings.Split(element.ChildText(".value strong"), " ")[0])
		case "Etaj":
			apartment.Floor = element.ChildText(".value strong a")
		case "An constructie":
			apartment.YearOfConstruction = element.ChildText(".value strong a")
		}
	})

	apartment.Description = e.ChildText("#textContent")
	var images []string
	e.ForEach(".img-item img", func(i int, element *colly.HTMLElement) {
		images = append(images, element.Attr("src"))
	})
	apartment.Images = images

	fmt.Println(apartment)
}
