package main

import (
	"flag"
	"fmt"
	"github.com/aliafshar/ghcontribs"
	"log"
	"time"
)

const SHORT_FORM = "2006-Jan-02"

var after = flag.String("after", "", "Date after which to count. e.g. 2013-Sep-30")

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: ghcontrib [options] username, [username, ...]\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	var d time.Time
	if *after == "" {
		d = time.Time{}
	} else {
		p, err := time.Parse(SHORT_FORM, *after)
		d = p
		if err != nil {
			log.Fatalln(err)
		}
	}
	for _, u := range flag.Args() {
		t, err := ghcontribs.TotalContribsFor(u, d)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(u, t)
	}
}
