package main

import (
	"bytes"
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

func NewClient(apiBase string) (*Client, error) {
	if apiBase == "" {
		return nil, errors.New("ApiBase is required to create a Client")
	}

	return &Client{
		Client:  &http.Client{},
		ApiBase: apiBase,
	}, nil
}

func (c *Client) Send(req *http.Request, v interface{}) error {
	var (
		err  error
		resp *http.Response
	)

	// default headers
	req.Header.Set("Content-Type", "text/xml")

	resp, err = c.Client.Do(req)
	c.log(req, resp)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return xml.NewDecoder(resp.Body).Decode(v)
}

func (c *Client) SetLog(log io.Writer) {
	c.Log = log
}

// log will dump request and response to the log file
func (c *Client) log(r *http.Request, resp *http.Response) {
	if c.Log != nil {
		var (
			reqDump  string
			respDump []byte
		)

		if r != nil {
			reqDump = fmt.Sprintf("%s %s. Data: %s", r.Method, r.URL.String(), r.Form.Encode())
		}
		if resp != nil {
			respDump, _ = httputil.DumpResponse(resp, true)
		}

		c.Log.Write([]byte(fmt.Sprintf("Request: %s\nResponse: %s\n", reqDump, string(respDump))))
	}
}

func (c *Client) NewRequest(ctx context.Context, payload interface{}) (*http.Request, error) {
	xmlData, _ := xml.MarshalIndent(payload, "", "  ")

	return http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.ApiBase,
		bytes.NewReader(xmlData),
	)
}
