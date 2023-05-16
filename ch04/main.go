// Grab a github issues
// https://developer.github.com/v3/search/#search-issues
package main

import (
	"fmt"
	"time"
)

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	var res IssuesSearchResult
	fmt.Println(res)
}
