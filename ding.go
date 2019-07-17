package senddingding

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

func sendDingdingMarkdown(robotUrl string, title string, text string) error {

	params := map[string]interface{}{
		"msgtype": "markdown",
		"markdown": map[string]interface{}{
			"title": title,
			"text":  text,
		},
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	httpClient := &http.Client{Transport: tr}

	var req *http.Request
	var err error
	rawParams, _ := json.Marshal(params)
	req, err = http.NewRequest("POST", robotUrl, bytes.NewReader(rawParams))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	response, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	if response != nil && response.Body != nil {
		defer response.Body.Close()
	}

	if response != nil && response.StatusCode != http.StatusOK {
		return nil
	} else {
		return fmt.Errorf("err response , status code %s", response.StatusCode)
	}
}
