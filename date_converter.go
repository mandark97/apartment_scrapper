package main

import (
	"fmt"
	"strings"
	"time"
)

var months = map[string]string{
	"ianuarie":   "January",
	"februarie":  "February",
	"martie":     "March",
	"aprilie":    "April",
	"mai":        "May",
	"iunie":      "June",
	"iulie":      "July",
	"august":     "August",
	"septembrie": "September",
	"octombrie":  "October",
	"noiembrie":  "November",
	"decembrie":  "December",
}

func parseDate(dateString string) (time.Time, error) {
	arr := strings.Split(dateString, ",")
	layout := "02 January 2006"
	for _, s := range arr {
		if strings.Contains(s, "2019") {
			return time.Parse(layout, replaceRoDate(s))
		}
	}

	return time.Now(), fmt.Errorf("couldn't parse date")
}

func replaceRoDate(roDate string) string {
	roDate = strings.ToLower(strings.Trim(roDate, " "))
	for rmonth, month := range months {
		if strings.Contains(roDate, rmonth) {
			return strings.Replace(roDate, rmonth, month, 1)
		}
	}

	return roDate
}
