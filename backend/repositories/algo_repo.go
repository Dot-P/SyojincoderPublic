package repositories

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/dot_p/syojin/models"
)

func FetchAlgoData(db *sql.DB, classifierMapping map[int]string, optionInt int) ([]models.AlgoData, error) {
	var results []models.AlgoData

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

		results = append(results, models.AlgoData{
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
