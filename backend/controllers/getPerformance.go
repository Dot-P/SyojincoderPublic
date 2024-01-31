package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getPerformance(name string) ([]int, []int, []string) {

	url := "https://atcoder.jp/users/"+ name + "/history/json"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
	}
	defer resp.Body.Close()

	// レスポンスボディの読み込み
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
	}

	type Performance struct {
		IsRated           bool   `json:"IsRated"`
		Place             int64  `json:"Place"`
		OldRating         int    `json:"OldRating"`
		NewRating         int    `json:"NewRating"`
		Performance       int    `json:"Performance"`
		InnerPerformance  int    `json:"InnerPerformance"`
		ContestScreenName string `json:"ContestScreenName"`
		ContestName       string `json:"ContestName"`
		EndTime           string `json:"EndTime"` 
	}
	
	// JSONデコード
	var performances []Performance
	if err := json.Unmarshal(body, &performances); err != nil {
		fmt.Println("Error decoding JSON:", err)
	}

	// 末尾から7個の要素を取り出す
	numItems := 7
	startIdx := len(performances) - numItems
	if startIdx < 0 {
		startIdx = 0
	}
	lastPerformances := performances[startIdx:]

	// デコードされたデータの使用
	var Performancearr []int = make([]int, len(lastPerformances))
	var Ratearr []int = make([]int, len(lastPerformances))
	var ContestName []string = make([]string, len(lastPerformances))

	for i, perf := range lastPerformances {
		Performancearr[i] = perf.Performance
		Ratearr[i] = perf.NewRating
		ContestName[i] = perf.ContestScreenName[:6]
	}

	return Performancearr, Ratearr, ContestName
}