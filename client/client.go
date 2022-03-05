package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"time"
)

const (
	timeoutSec = 10
)

type PingResult struct {
	Status string `json:"status"`
}

type Param struct {
	Code         string    `json:"code"`
	Base64Images [4]string `json:"images"`
}

type Result struct {
	Status       int       `json:"status"`
	Stdout       string    `json:"stdout"`
	Stderr       string    `json:"stderr"`
	Base64Images [4]string `json:"images"`
	ElapsedTime  string    `json:"elapsed_time"`
}

func newURL(p string) string {
	u := &url.URL{
		Scheme: "https",
		Host:   "websh.jiro4989.com",
		Path:   path.Join("api", p),
	}
	return u.String()
}

func Ping() (*PingResult, error) {
	url := newURL("ping")
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%d %s", res.StatusCode, res.Status)
	}

	var result PingResult
	d := json.NewDecoder(res.Body)
	err = d.Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %s", err)
	}
	return &result, err
}

func (p *Param) Set(code string) *Param {
	p.Code = code
	return p
}

func newClient() *http.Client {
	return &http.Client{
		Timeout: timeoutSec * time.Second,
	}
}

func Post(p Param) (*Result, error) {
	j, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", newURL("shellgei"), bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}

	res, err := newClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result Result
	d := json.NewDecoder(res.Body)
	if err = d.Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
