package main

import (
	"flag"
	"fmt"
	"github.com/tcnksm/go-httpstat"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func latency(hostname string) {
	// Create a new HTTP request
	url := fmt.Sprintf("http://%s", hostname)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Create a httpstat powered context
	var result httpstat.Result
	ctx := httpstat.WithHTTPStat(req.Context(), &result)
	req = req.WithContext(ctx)
	// Send request by default HTTP client
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
	log.Printf("DNS lookup: %d ms", int(result.DNSLookup/time.Millisecond))
}

func main() {
	speed := flag.Int("speed", 1, "Speed per second to send the messages at")
	number := flag.Int("number", 1, "Amount of times to call url")
	hostname := flag.String("hostname", "", "Hostname to hit")

	flag.Parse()
	i := 0
	for i < *number {
		go latency(*hostname)
		i++
		time.Sleep(time.Second / time.Duration(*speed))
	}
}
