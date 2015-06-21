package main

import (
	"regexp"
	"strings"
)

type Histogram map[string]int

func Scrape(text string) Histogram {
	flat := strings.Replace(text, "\n", "", -1)
	chop := strings.Split(flat, " ")

	hist := make(Histogram)
	for _, val := range chop {
		val = strings.ToLower(val)
		if hist[val] > 0 {
			hist[val] += 1
		} else {
			hist[val] = 1
		}
	}
	return hist
}

func Clean(source string) string {
	r1 := regexp.MustCompile("[^0-9A-Za-z]")
	source = r1.ReplaceAllString(source, " ")
	r2 := regexp.MustCompile("[ ]+")
	source = r2.ReplaceAllString(source, " ")
	return source
}

func (hist Histogram) GreaterThan(n int) Histogram {
	results := make(Histogram)
	for key, val := range hist {
		if val > n {
			results[key] = val
		}
	}
	return results
}

func (hist Histogram) LessThan(n int) Histogram {
	results := make(Histogram)
	for key, val := range hist {
		if val < n {
			results[key] = val
		}
	}
	return results
}

func (hist Histogram) Equals(n int) Histogram {
	return hist.GreaterThan(n - 1).LessThan(n + 1)
}
