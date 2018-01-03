// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

var fetchCount = flag.Int("count", 10, "number of random sites to fetch")
var randSeed = flag.Int64("seed", time.Now().Unix(), "random seed")

func main() {
	flag.Parse()

	log.Printf("Fetching top 1m sites")

	sites, err := FetchTopSites()
	if err != nil {
		log.Fatalf("%s", err)
	}

	log.Printf("Done")

	rand.Seed(*randSeed)

	start := time.Now()
	ch := make(chan string)

	log.Printf("fetching %d URLs", *fetchCount)

	for i := 0; i < *fetchCount; i++ {
		url := sites[rand.Intn(len(sites))]
		go fetch(url, ch) // start a goroutine
	}

	for i := 0; i < *fetchCount; i++ {
		fmt.Println(<-ch) // receive from channel ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// FetchTopSites returns the top 1M dataset from Alexa.
func FetchTopSites() ([]string, error) {
	resp, err := http.Get("http://s3.amazonaws.com/alexa-static/top-1m.csv.zip")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch top1m dataset: %s", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %s", err)
	}

	zipReader, err := zip.NewReader(bytes.NewReader(body), (int64)(len(body)))
	if err != nil {
		return nil, fmt.Errorf("failed to unzip body: %s", err)
	}

	for _, f := range zipReader.File {
		if f.Name != "top-1m.csv" {
			return nil, fmt.Errorf("unexpected zip file entry '%s'", f.Name)
		}

		top1m, err := f.Open()
		if err != nil {
			return nil, fmt.Errorf("failed to extract '%s': %s", f.Name, err)
		}

		csvReader := csv.NewReader(top1m)

		records := make([]string, 0, 1000000)
		for {
			// Read CSV records, expecting 2 fields, e.g. "1,google.com"
			r, err := csvReader.Read()
			switch err {
			case nil:
				records = append(records, r[1])
			case io.EOF:
				return records, nil

			default:
				return nil, fmt.Errorf("bad CSV record: %s", err)
			}

		}
	}

	return nil, errors.New("empty zip file")
}

func fetch(str string, ch chan<- string) {
	url, err := url.Parse(str)
	if err != nil {
		ch <- fmt.Sprintf("failed to parse URL '%s': %s", str, err)
		return
	}

	if url.Scheme == "" {
		url.Scheme = "http"
	}

	start := time.Now()
	resp, err := http.Get(url.String())
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
