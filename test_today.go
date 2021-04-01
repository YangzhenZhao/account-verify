package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/YangzhenZhao/account-verify/verify"
)

func getOrders() []verify.Order {
	file, err := os.Open("today_orders.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)

	var orders []verify.Order
	if err := json.Unmarshal(content, &orders); err != nil {
		log.Fatal(err)
	}

	return orders
}

type Position struct {
	Code string
	Vol  int32
}

func getPositions() map[string]int32 {
	file, err := os.Open("today_positions.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)

	var positions []Position
	if err := json.Unmarshal(content, &positions); err != nil {
		log.Fatal(err)
	}

	res := make(map[string]int32)
	for _, item := range positions {
		res[item.Code] = item.Vol
	}
	return res
}

func main() {
	nowPositions := getPositions()
	orders := getOrders()
	// fmt.Println(nowPositions["002376"])
	// fmt.Printf("%+v\n", orders[0])
	// 初始资金 19943587.40，收盘资金 20144639.39
	verify.VerifyToday(19943587400000, nowPositions, orders, 20144639390000)
}
