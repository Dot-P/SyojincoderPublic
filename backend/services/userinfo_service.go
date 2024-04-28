package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/dot_p/syojin/models"
	"github.com/dot_p/syojin/repositories"
)

// 提出数と間違った問題番号を特定
func (s *MyAppService) GetSubmissionService(name string) ([]int, []string) {
	// 特定URLへのアクセス
	now := time.Now()
	nineWeeksAgo := now.AddDate(0, 0, -7*9)
	unixSeconds := strconv.FormatInt(nineWeeksAgo.Unix(), 10)
	url := "https://kenkoooo.com/atcoder/atcoder-api/v3/user/submissions?user=" + name + "&from_second=" + unixSeconds

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, nil
	}

	var submissions []models.Submission
	if err := json.Unmarshal(body, &submissions); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, nil
	}

	var submissionList []int = make([]int, 9)

	for i := 0; i < 9; i++ {
		submissionCount := 0
		UpxWeeksago := now.AddDate(0, 0, -7*i).Unix()
		DownxWeeksago := now.AddDate(0, 0, -7*(i+1)).Unix()
		for _, submisson := range submissions {
			if DownxWeeksago < submisson.EpochSecond && submisson.EpochSecond <= UpxWeeksago && submisson.Result == "AC" {
				submissionCount += 1
			}
		}
		submissionList[i] = submissionCount
	}

	var wrongList []string

	for _, submisson := range submissions {
		if submisson.Result != "AC" {
			wrongList = append(wrongList, submisson.ProblemID)
		}
	}

	// 重複を除く
	unique := make(map[string]bool)
	result := []string{}

	for _, value := range wrongList {
		if _, ok := unique[value]; !ok {
			unique[value] = true
			result = append(result, value)
		}
	}
	wrongList = result

	return submissionList, wrongList
}

/*============================================================================================*/

// コンテストの成績の情報を処理
func (s *MyAppService) GetPerformanceService(name string) ([]int, []int, []string) {

	url := "https://atcoder.jp/users/" + name + "/history/json"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, nil, nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, nil, nil
	}

	var performances []models.Performance
	if err := json.Unmarshal(body, &performances); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil, nil, nil
	}

	// 末尾から7個の要素を取り出す
	numItems := 7
	startIdx := len(performances) - numItems
	if startIdx < 0 {
		startIdx = 0
	}
	lastPerformances := performances[startIdx:]

	// デコードされたデータの使用
	var PerformanceList []int = make([]int, len(lastPerformances))
	var RateList []int = make([]int, len(lastPerformances))
	var ContestName []string = make([]string, len(lastPerformances))

	for i, perf := range lastPerformances {
		PerformanceList[i] = perf.Performance
		RateList[i] = perf.NewRating
		ContestName[i] = perf.ContestScreenName[:6]
	}

	return PerformanceList, RateList, ContestName
}

/*=======================================================================*/

// Embeddeingによって、類似度を計算する
func (s *MyAppService) CalEmbeddingService(wrongs []string) ([]float64, []string, []string) {

	var dotProductValues []float64
	var SimilarProblems []string
	var TargetWrongs []string

	var allEmbeddings []models.EmbeddingInfo = repositories.FetchAllEmbeddings(s.db)

	dotProductValues, SimilarProblems, TargetWrongs = repositories.FetchSimilarities(s.db, wrongs, allEmbeddings)

	return dotProductValues, SimilarProblems, TargetWrongs
}
