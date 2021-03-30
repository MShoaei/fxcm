package fxdl

import "time"

type Candle struct {
	Start      time.Time
	End        time.Time
	OpenPrice  float64
	ClosePrice float64
	MaxPrice   float64
	MinPrice   float64
	Volume     int
}
