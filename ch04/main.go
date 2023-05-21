// ch04: A github issue viewer
//
// # Examples
//
// ```
// $ go run . golang label:bug is:open json decoder
// 13 issues
// 1       mrsufgi https://api.github.com/repos/kuberhealthy/kuberhealthy/issues/1123
// 2       gsaraf  https://api.github.com/repos/open-telemetry/opentelemetry-go/issues/3063
// 3       jharting        https://api.github.com/repos/getkin/kin-openapi/issues/294
// 4       Ankurk99        https://api.github.com/repos/accuknox/discovery-engine/issues/685
// 5       sdminonne       https://api.github.com/repos/AmadeusITGroup/workflow-controller/issues/89
// 6       elitvinov       https://api.github.com/repos/mittwald/kubernetes-replicator/issues/216
// 7       liuwqiang       https://api.github.com/repos/TykTechnologies/tyk/issues/4066
// 8       bboreham        https://api.github.com/repos/weaveworks-experiments/kspan/issues/14
// 9       renovate[bot]   https://api.github.com/repos/moul/converter/issues/63
// 10      xlfjn   https://api.github.com/repos/webp-sh/webp_server_go/issues/75
// 11      TibebeJS        https://api.github.com/repos/xelaj/mtproto/issues/49
// 12      pcolazurdo      https://api.github.com/repos/aws/aws-sdk-go-v2/issues/1665
// 13      cstrahan        https://api.github.com/repos/bazelbuild/rules_go/issues/2654
// ```
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

func main() {
	result, err := searchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues\n", result.TotalCount)
	for i, issue := range result.Issues {
		fmt.Printf("%d\t%s\t%s\n", i+1, issue.User.Login, issue.URL)
	}
}

const issuesURL = "https://api.github.com/search/issues"

type IssuesResult struct {
	TotalCount int      `json:"total_count"`
	Issues     []*Issue `json:"items"`
}

type Issue struct {
	URL  string
	User *User
}

type User struct {
	Login string
}

func searchIssues(terms []string) (*IssuesResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(issuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
