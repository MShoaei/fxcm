package fxcm

import (
	gosocketio "github.com/graarh/golang-socketio"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	socketio   *gosocketio.Client
	APIToken   string
}
