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

import "fmt"

type symbolInfo struct {
	tickSize float64
}

type ExchangeInfoService struct {
	Symbols map[string]symbolInfo
}

func NewExchangeInfoService() *ExchangeInfoService {
	return &ExchangeInfoService{
		Symbols: make(map[string]symbolInfo),
	}
}

func (s *ExchangeInfoService) Update() error {
	exchangeInfo, err := GetExchangeInfo()
	if err != nil {
		return err
	}
	for _, symbol := range exchangeInfo.Symbols {
		symbolInfo := symbolInfo{}
		for _, filter := range symbol.Filters {
			if filter.FilterType == "PRICE_FILTER" {
				symbolInfo.tickSize = filter.TickSize
			}
		}
		s.Symbols[symbol.Symbol] = symbolInfo
	}
	return nil
}

func (s *ExchangeInfoService) GetTickSize(symbol string) (float64, error) {
	symbolInfo, ok := s.Symbols[symbol]
	if !ok {
		return 0, fmt.Errorf("symbol not found")
	}
	return symbolInfo.tickSize, nil
}
