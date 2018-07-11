package fcoin

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Client struct {
	// http
	key    string
	secret string
	client *http.Client

	// websocket
	WS *websocket.Conn
}

func Authorize(key, secret string, localtime int64) (*Client, error) {
	c := &Client{
		client: &http.Client{},
	}

	// "auth" by checking server time
	rsp, err := c.ServerTime()
	if err != nil {
		return nil, err
	}

	diff := rsp.Data - localtime
	if diff > MaxTimeDiffMs || diff < -MaxTimeDiffMs {
		return nil, fmt.Errorf("Inconsistent time from server")
	}

	c.key = key
	c.secret = secret
	return c, nil
}

// Internal helper method for final request
// NOTE: non-zero status code does not cause error
func (c *Client) request(endpoint string, method string, signReq bool, args interface{}, ret interface{}) error {
	argStr := ""
	uri := BaseUrl + endpoint
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return err
	}

	if args != nil {
		// parse args into:
		// - body: json dump
		reader, err := encode(args)
		if err != nil {
			return err
		}
		// - argStr: for signature
		argValues := structToMap(args)
		argStr = argValues.Encode()

		// Append args to either body or uri, depends on request method
		if method == "POST" {
			req.Body = ioutil.NopCloser(reader)
			req.Header.Add("Content-type", "application/json")
		} else if method == "GET" {
			req.URL.RawQuery = argStr
		}
	}

	// As of now, 3 public API does not need signing
	if signReq {
		t := time.Now().Unix() * 1000
		timeStr := strconv.FormatInt(t, 10)
		signature := Sign(method, uri, timeStr, argStr, c.secret)

		req.Header.Add("FC-ACCESS-KEY", c.key)
		req.Header.Add("FC-ACCESS-SIGNATURE", signature)
		req.Header.Add("FC-ACCESS-TIMESTAMP", timeStr)
	}

	// Try unmarshal to rsp struct
	rsp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("Failed requesting %v: %v", uri, err)
	}
	defer rsp.Body.Close()

	if err := json.NewDecoder(rsp.Body).Decode(ret); err != nil {
		return fmt.Errorf("Invalid response format %v", err)
	}
	return nil
}
