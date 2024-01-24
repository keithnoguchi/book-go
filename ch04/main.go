// Github issue lister
package main

import (
	//"encoding/json"
	"fmt"
	//"log"
	//"net/http"
	"net/url"
	"os"
	"strings"
)

const githubURL = "https://api.github.com/search/issues"

func main() {
	q := url.QueryEscape(strings.Join(os.Args[1:], " "))
	fmt.Printf("%s\n", q)
}

type Issues struct {
	TotalCount int      `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number int
}
