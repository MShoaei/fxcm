package fxcm

import (
	"compress/gzip"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var millenium = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

type period time.Duration

const (
	M1  = period(1 * time.Minute)
	M5  = period(5 * time.Minute)
	M15 = period(15 * time.Minute)
	M30 = period(30 * time.Minute)
)

func NewHistoryClient(p period) *Client {
	return &Client{
		p:  time.Duration(p),
		hc: &http.Client{},
	}
}

func (c *Client) SetProxy(proxy string) {
	proxyURL, _ := url.Parse(proxy)

	c.hc.Transport = &http.Transport{
		Proxy:           http.ProxyURL(proxyURL),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
}

type data struct {
	Time   []int     `json:"time"`
	Open   []float64 `json:"open"`
	High   []float64 `json:"high"`
	Low    []float64 `json:"low"`
	Close  []float64 `json:"close"`
	Volume []int     `json:"volume"`
}

func (c *Client) Do(symbol string) ([]*Candle, error) {
	u := fmt.Sprintf("https://data.forexsb.com/data/%s%s.gz", symbol, c.p)
	resp, err := c.hc.Get(u)
	if err != nil {
		return nil, err
	}

	r, err := gzip.NewReader(resp.Body)
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	d := &data{}

	if err = json.Unmarshal(b, d); err != nil {
		return nil, err
	}

	candles := make([]*Candle, 0, 200_000)

	for i := 0; i < len(d.Time); i++ {
		c := Candle{
			Start:      millenium.Add(time.Duration(d.Time[i]) * 60 * time.Second),
			End:        millenium.Add(time.Duration(d.Time[i])*60*time.Second + c.p - 1),
			OpenPrice:  d.Open[i],
			ClosePrice: d.Close[i],
			MaxPrice:   d.High[i],
			MinPrice:   d.Low[i],
			Volume:     d.Volume[i],
		}
		candles = append(candles, &c)
	}
	return candles, nil
}
