package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "io/ioutil"
)

type Ticker struct {
	ProductCode     string  `json:"product_code"`
	Timestamp       string  `json:"timestamp"`
	TickID          int     `json:"tick_id"`
	BestBid         float64 `json:"best_bid"`
	BestAsk         float64 `json:"best_ask"`
	BestBidSize     float64 `json:"best_bid_size"`
	BestAskSize     float64 `json:"best_ask_size"`
	TotalBidDepth   float64 `json:"total_bid_depth"`
	TotalAskDepth   float64 `json:"total_ask_depth"`
	Ltp             float64 `json:"ltp"`
	Volume          float64 `json:"volume"`
	VolumeByProduct float64 `json:"volume_by_product"`
}

func main() {
    url := "https://api.bitflyer.jp/v1/getticker"
    resp, _ := http.Get(url)
    defer resp.Body.Close()
    byteArray, _ := ioutil.ReadAll(resp.Body)

    jsonBytes := ([]byte)(byteArray)
    data := new(Ticker)

    if err := json.Unmarshal(jsonBytes, data); err != nil {
      fmt.Println("JSON Unmarshal error:", err)
      return
    }
    fmt.Println(data.ProductCode)
}
