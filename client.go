package fxdl

import (
	"net/http"
	"time"
)

type Client struct {
	p  time.Duration
	hc *http.Client
}
