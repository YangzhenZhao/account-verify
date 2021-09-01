// // 20200710
// package main

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// 	"os"

// 	"github.com/YangzhenZhao/account-verify/verify"
// )

// func getOrders() []verify.Order {
// 	file, err := os.Open("orders.json")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
// 	content, err := ioutil.ReadAll(file)

// 	var orders []verify.Order
// 	if err := json.Unmarshal(content, &orders); err != nil {
// 		log.Fatal(err)
// 	}

// 	return orders
// }

// type Position struct {
// 	Code string
// 	Vol  int32
// }

// func getPositions() map[string]int32 {
// 	file, err := os.Open("positions.json")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
// 	content, err := ioutil.ReadAll(file)

// 	var positions []Position
// 	if err := json.Unmarshal(content, &positions); err != nil {
// 		log.Fatal(err)
// 	}

// 	res := make(map[string]int32)
// 	for _, item := range positions {
// 		res[item.Code] = item.Vol
// 	}
// 	return res
// }

// func getPrice() map[string]verify.CodeVerifyPrice {
// 	file, err := os.Open("price.json")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
// 	content, err := ioutil.ReadAll(file)

// 	var res map[string]verify.CodeVerifyPrice
// 	if err := json.Unmarshal(content, &res); err != nil {
// 		log.Fatal(err)
// 	}

// 	return res
// }

// func main() {
// 	nowPositions := getPositions()
// 	orders := getOrders()
// 	// fmt.Println(nowPositions["002376"])
// 	// fmt.Printf("%+v\n", orders[0])
// 	// codesMap := make(map[string]bool)
// 	// codes := []string{}
// 	// for code := range nowPositions {
// 	// 	codes = append(codes, code)
// 	// 	codesMap[code] = true
// 	// }
// 	// for _, order := range orders {
// 	// 	if _, ok := codesMap[order.Code]; !ok {
// 	// 		codesMap[order.Code] = true
// 	// 		codes = append(codes, order.Code)
// 	// 	}
// 	// }
// 	// fmt.Print("[")
// 	// for _, code := range codes {
// 	// 	fmt.Print(`"`, code, `"`, ",")
// 	// }
// 	// fmt.Println("]")

// 	codesVerifyPrice := getPrice()
// 	// fmt.Println(codesVerifyPrice["000001"])
// 	// 初始资金 1000000，收盘资金 999992.27
// 	verify.Verify(1000000000000, nowPositions, orders, 999992270000, codesVerifyPrice)
// }
