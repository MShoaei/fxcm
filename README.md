# FXDL

`fxdl` downloads a maximum of 200,000 candles from [https://forexsb.com](https://forexsb.com)

## Install

To install the library, run:

`$ go get github.com/MShoaei/fxdl`

## Example

```go
package main

import "github.com/MShoaei/fxdl"

func main() {
	// create a client and pass the time frame of your choice
	// valid time frames are: M1, M5, M15, M30
	c := NewHistoryClient(M5)
	// you can set a proxy if you need
	c.SetProxy("socks5://127.0.0.1:12123")
	// *Client.Do fetches the data of the symbol in previously provided time frame
	// and returns a slice of candles. the length of slice should be 200,000 
	candles, err := c.Do("EURUSD")
}
```