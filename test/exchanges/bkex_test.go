package test

// Copyright (c) 2015-2019 Bitontop Technologies Inc.
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

import (
	"fmt"
	"log"
	"testing"

	"github.com/bitontop/gored/coin"
	"github.com/bitontop/gored/exchange"
	"github.com/bitontop/gored/pair"

	"github.com/bitontop/gored/exchange/bkex"
	"github.com/bitontop/gored/test/conf"
	// "../../exchange/bkex"
	// "../conf"
)

/********************Public API********************/
func Test_Bkex(t *testing.T) {
	e := InitBkex()

	pair := pair.GetPairByKey("BTC|ETH")

	// Test_Coins(e)
	// Test_Pairs(e)
	Test_Pair(e, pair)
	// Test_Orderbook(e, pair)
	// Test_ConstraintFetch(e, pair)
	// Test_Constraint(e, pair)

	Test_Balance(e, pair)
	Test_Trading(e, pair, 0.00000001, 100)
	// Test_Trading_Sell(e, pair, 0.00000001, 667)
	// Test_Withdraw(e, pair.Base, 1.5, "1HB5XMLmzFVj8ALj6mfBsbifRoD4miY36v")

	Test_CheckBalance(e, pair.Target, exchange.AssetWallet)
	Test_CheckAllBalance(e, exchange.AssetWallet)
	// Test_DoTransfer(e, pair.Target, "1", exchange.AssetWallet, exchange.SpotWallet)
	Test_DoWithdraw(e, pair.Target, "1", "0x37E0Fc27C6cDB5035B2a3d0682B4E7C05A4e6C46", "tag")
}

func InitBkex() exchange.Exchange {
	coin.Init()
	pair.Init()
	config := &exchange.Config{}
	config.Source = exchange.EXCHANGE_API
	conf.Exchange(exchange.BKEX, config)
	fmt.Printf("%+v\n", config)

	ex := bkex.CreateBkex(config)
	log.Printf("Initial [ %v ] ", ex.GetName())

	config = nil
	return ex
}
