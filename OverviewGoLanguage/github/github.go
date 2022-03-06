package github

import (
	"encoding/json"
	"fmt"
	"github/atakanteko/golangplayground/helpers"
	"io/ioutil"
	"net/http"
)

func SearchIssues() (*helpers.IssueResult, error) {
	resp, err := http.Get("https://api.github.com/search/issues" + "?q=" + "java")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		err := resp.Body.Close()
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("search query failed:%s", resp.Status)
	}
	var issues helpers.IssueResult
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &issues)
	if err != nil {
		return nil, err
	}
	return &issues, nil
}
