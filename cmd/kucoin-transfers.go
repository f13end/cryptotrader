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

package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/crankykernel/ctrader/cmd/kucoin"
)

var kucoinTransfersCmd = &cobra.Command{
	Use:   "transfers [coin...]",
	Short: "Print deposits and withdrawals",
	Long: `Print deposits and withdrawals for one or more coins

Example: List BTC and LTC transfers:

    ctrader kucoin transfers BTC LTC

Example: List transfers for all coins:

    ctrader kucoin transfers

Warning: Listing transfers for all coins can take a while and may hit
         API limits.
`,
	Run: func(cmd *cobra.Command, args []string) {
		kucoin.Transfers(args)
	},
}

func init() {
	kucoinCmd.AddCommand(kucoinTransfersCmd)
}