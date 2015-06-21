package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	b, _ := ioutil.ReadFile("test.txt")
	text := Clean(string(b))
	hist := Scrape(text).Equals(10)

	fmt.Println(hist)
}
