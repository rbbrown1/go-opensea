package goopensea

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Ctx        context.Context
	RetryCount int
}

func NewClient(ctx context.Context) *Client {

	client := Client{
		BaseURL: "https://api.opensea.io",
		HTTPClient: &http.Client{
			Timeout: 15 * time.Second,
		},
		Ctx: ctx,
	}
	return &client
}

func (c *Client) Request(method string, url string,
	queryParams interface{}, result interface{}) (res *http.Response, err error) {

	// build the URL query string if one is given
	if queryParams != nil {
		q, _ := query.Values(queryParams)
		url = fmt.Sprintf("%s%s?%s", c.BaseURL, url, q.Encode())
	} else {
		url = fmt.Sprintf("%s%s", c.BaseURL, url)
	}

	// create the request object
	req, err := http.NewRequestWithContext(c.Ctx, method, url, nil)
	if err != nil {
		return res, err
	}

	// execute the query
	res, err = c.HTTPClient.Do(req)
	if err != nil {
		return res, err
	}
	defer res.Body.Close()

	// if not a 200, then return a error
	if res.StatusCode != 200 {
		defer res.Body.Close()

		body, _ := ioutil.ReadAll(res.Body)

		return res, fmt.Errorf("%s", body)
	}

	// if a return struct is given, dump data into it
	if result != nil {
		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(result)
		if err != nil {
			return res, err
		}
	}

	return res, nil
}
