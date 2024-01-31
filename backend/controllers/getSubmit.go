package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"strconv"
)


func removeDuplicates(slice []string) []string {
    seen := make(map[string]bool)
    result := []string{}

    for _, value := range slice {
        if _, ok := seen[value]; !ok {
            seen[value] = true
            result = append(result, value)
        }
    }

    return result
}

func getSubmission(name string) ([]int, []string) {
	now := time.Now()

	nineWeeksAgo := now.AddDate(0, 0, -7*9)

	unixSeconds := strconv.FormatInt(nineWeeksAgo.Unix(), 10)

	url := "https://kenkoooo.com/atcoder/atcoder-api/v3/user/submissions?user=" + name + "&from_second=" + unixSeconds

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

	type Submission struct {
		ID            int64   `json:"id"`
		EpochSecond   int64   `json:"epoch_second"`
		ProblemID     string  `json:"problem_id"`
		ContestID     string  `json:"contest_id"`
		UserID        string  `json:"user_id"`
		Language      string  `json:"language"`
		Point         float64 `json:"point"`
		Length        int     `json:"length"`
		Result        string  `json:"result"`
		ExecutionTime int     `json:"execution_time"`
	}

	// JSONデコード
	var submissions []Submission
	if err := json.Unmarshal(body, &submissions); err != nil {
		fmt.Println("Error decoding JSON:", err)
	}

	// デコードされたデータの使用
	var submissionarr []int = make([]int, 9)

	for xx := 0; xx < 9; xx++{
		submissionCount := 0
		UpxWeeksago := now.AddDate(0, 0, -7*xx).Unix()
		DownxWeeksago := now.AddDate(0, 0, -7*(xx+1)).Unix()
		for _, sub := range submissions {
			if DownxWeeksago < sub.EpochSecond && sub.EpochSecond <= UpxWeeksago && sub.Result=="AC"{
				submissionCount += 1
			}
		}
		submissionarr[xx] = submissionCount
	}

	var wrongarr []string

	for _, sub := range submissions {
		if sub.Result=="WA"{
			wrongarr = append(wrongarr, sub.ProblemID)
	    }
	}

	wrongarr = removeDuplicates(wrongarr)

	return submissionarr, wrongarr
}