// A github issue searcher
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const issuesURL = "https://api.github.com/search/issues"

func main() {
	q := url.QueryEscape(strings.Join(os.Args[1:], " "))
	resp, err := http.Get(issuesURL + "?q=" + q)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("HTTP status code: %d", resp.StatusCode)
	}
	var issues Issues
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", issues.TotalCount)
}

type Issues struct {
	TotalCount int      `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number int
}
