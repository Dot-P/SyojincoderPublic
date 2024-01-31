package controllers

import (
    "fmt"
    "strconv"
    "strings"
)

// AlgoData 構造体
type AlgoData struct {
    ProblemID    string `json:"problem_id"`
    Difficulties string `json:"difficulties"`
    Classifier   string `json:"classifier"`
    URL          string `json:"url"`
}

func expAlgo(optionInt int) ([]AlgoData, error) {
    // データベース接続
    db := connectDB()
    defer db.Close()

    // Classifierの分類をマッピングする
    classifierMapping := map[int]string{
        1:  "全探索",
        2:  "累積和及びimos法",
        3:  "二分探索",
        4:  "座標圧縮及びランレングス圧縮",
        5:  "動的計画法",
        6:  "しゃくとり法",
        7:  "BFS",
        8:  "Union-Find",
        9:  "ワーシャルフロイド法",
        10: "ユークリッドの互除法",
        11: "エラトステネスの篩及び整数問題",
        12: "セグ木",
        13: "論理演算",
        14: "包除原理",
        15: "DFS",
    }

    var results []AlgoData

    // SQLクエリを実行
    query := "SELECT category, difficulties, qnum FROM algo WHERE classifier = ?"
    rows, err := db.Query(query, optionInt)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    for rows.Next() {
        var category, difficulties string
        var qnum int
        if err := rows.Scan(&category, &difficulties, &qnum); err != nil {
            return nil, err
        }

        // AlgoData構造体にデータを格納
        problemID := strings.ToUpper(category) + strconv.Itoa(qnum)
        url := fmt.Sprintf("https://atcoder.jp/contests/%s/tasks/%s_%s", strings.ToLower(category)+strconv.Itoa(qnum), strings.ToLower(category)+strconv.Itoa(qnum), strings.ToLower(difficulties))

        classifierStr, ok := classifierMapping[optionInt]
        if !ok {
            classifierStr = "不明"
        }

        results = append(results, AlgoData{
            ProblemID:    problemID,
            Difficulties: difficulties,
            Classifier:   classifierStr,
            URL:          url,
        })
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return results, nil
}
