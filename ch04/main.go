// ch04: A github issue viewer
package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {
	result, err := searchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues\n", result.TotalCount)
}

type IssuesResult struct {
	TotalCount int `json:"total_count"`
}

func searchIssues(terms []string) (*IssuesResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	fmt.Println(q)
	return &IssuesResult{}, nil
}
