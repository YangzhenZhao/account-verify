package verify

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/YangzhenZhao/account-verify/verify/consts"
	"github.com/YangzhenZhao/goquotes/quotes/stock"
)

type Side int

const (
	BUY Side = iota
	SELL
)

func (side *Side) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	case "buy":
		*side = BUY
	case "sell":
		*side = SELL
	}

	return nil
}

type CodeVerifyPrice struct {
	Close    int64
	PreClose int64
}

type Order struct {
	Code      string
	Side      Side
	TotalFee  int64
	DeaVol    int32
	DealValue int64
}

func VerifyToday(
	beginTotalAssets int64,
	nowPositions map[string]int32,
	orders []Order,
	nowTotalAssets int64,
) {
	codesVerifyPrice := make(map[string]CodeVerifyPrice)
	codesMap := make(map[string]bool)
	codes := []string{}
	for code := range nowPositions {
		codes = append(codes, code)
		codesMap[code] = true
	}
	for _, order := range orders {
		if _, ok := codesMap[order.Code]; !ok {
			codesMap[order.Code] = true
			codes = append(codes, order.Code)
		}
	}
	quote := stock.SinaQuote{}
	for _, code := range codes {
		tick, err := quote.Tick(code)
		if err != nil {
			log.Fatalln(err)
			return
		}
		close, preClose := int64(tick.CurrentPrice*1000000), int64(tick.PreClose*1000000)
		codesVerifyPrice[code] = CodeVerifyPrice{
			Close:    close,
			PreClose: preClose,
		}
	}
	Verify(beginTotalAssets, nowPositions, orders, nowTotalAssets, codesVerifyPrice)
}

func Verify(
	preTotalAssets int64,
	nowPositions map[string]int32,
	orders []Order,
	nowTotalAssets int64,
	codesVerifyPrice map[string]CodeVerifyPrice,
) {
	totalProfit := nowTotalAssets - preTotalAssets
	var noTradePositionsProfit int64 = 0
	var buyOrdersProfit int64 = 0
	var sellOrdersProfit int64 = 0

	noTradePositions := make(map[string]int32)
	for code, vol := range nowPositions {
		noTradePositions[code] = vol
	}
	for _, order := range orders {
		if order.Side == BUY {
			noTradePositions[order.Code] -= order.DeaVol
			currentPrice := codesVerifyPrice[order.Code].Close
			theOrderProfit := currentPrice*int64(order.DeaVol) - (order.DealValue + order.TotalFee)
			buyOrdersProfit += theOrderProfit
		} else {
			preClose := codesVerifyPrice[order.Code].PreClose
			theOrderProfit := order.DealValue - order.TotalFee - preClose*int64(order.DeaVol)
			sellOrdersProfit += theOrderProfit
		}
	}

	for code, vol := range noTradePositions {
		if vol < 0 {
			log.Fatalln("vol = ", vol)
		}
		priceMsg := codesVerifyPrice[code]
		currentPrice, preClose := priceMsg.Close, priceMsg.PreClose
		noTradePositionsProfit += int64(vol) * (currentPrice - preClose)
	}

	fmt.Printf("Actually profit = %d\n", totalProfit)
	fmt.Printf("noTradePositionsProfit = %d\n", noTradePositionsProfit)
	fmt.Printf("buyOrdersProfit = %d\n", buyOrdersProfit)
	fmt.Printf("sellOrdersProfit = %d\n", sellOrdersProfit)
	diff := totalProfit - noTradePositionsProfit - buyOrdersProfit - sellOrdersProfit
	fmt.Printf("Diff = %d, DiffYuan = %f\n", diff, float64(diff)/consts.MONEY_UNIT)
}
