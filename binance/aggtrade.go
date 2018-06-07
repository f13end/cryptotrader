// The MIT License (MIT)
//
// Copyright (c) 2018 Cranky Kernel
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use, copy,
// modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS
// BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
// ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package binance

import (
	"time"
)

type StreamAggTrade struct {
	EventType       string  `json:"e"`
	EventTimeMillis int64   `json:"E"`
	Symbol          string  `json:"s"`
	TradeID         int64   `json:"a"`
	Price           float64 `json:"p,string"`
	Quantity        float64 `json:"q,string"`
	FirstTradeID    int64   `json:"f"`
	LastTradeID     int64   `json:"l"`
	TradeTimeMillis int64   `json:"T"`
	BuyerMaker      bool    `json:"m"`
	Ignored         bool    `json:"M"`
}

func (t *StreamAggTrade) QuoteQuantity() float64 {
	return t.Quantity * t.Price
}

func (t *StreamAggTrade) Timestamp() time.Time {
	return time.Unix(0, t.TradeTimeMillis*int64(time.Millisecond))
}
