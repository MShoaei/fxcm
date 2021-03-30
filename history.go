package fxcm

import (
	"crypto/tls"
	"encoding/json"
	gosocketio "github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"net/http"
	"net/url"
)

func NewHistoryClient(api string) (*Client, error) {
	var completeURL = "https://tradingstation3-demo.fxcm.com/socket.io/?access_token=" + api
	c, err := gosocketio.Dial(gosocketio.GetUrl(completeURL, 443, true), &transport.WebsocketTransport{PingInterval: 25000, PingTimeout: 60000})
	if err != nil {
		return nil, err
	}

	return &Client{
		APIToken:   api,
		socketio:   c,
		httpClient: &http.Client{},
	}, nil
}

func (c *Client) SetProxy(proxy string) {
	proxyURL, _ := url.Parse(proxy)

	c.httpClient.Transport = &http.Transport{
		Proxy:           http.ProxyURL(proxyURL),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
}

func (c *Client) Do() ([]*Candle, error) {
	req, err := http.NewRequest(http.MethodGet, (&url.URL{}).String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "node-XMLHttpRequest")
	req.Header.Set("Authorization", "Bearer "+c.socketio.Id()+c.APIToken)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	json.NewDecoder(resp.Body).Decode(nil)
	return nil, nil
}
