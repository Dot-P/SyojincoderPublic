package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"sort"

	"github.com/dot_p/syojin/common"
	"github.com/dot_p/syojin/models"
)

// データベースから全ての embedding を取得
func FetchAllEmbeddings(db *sql.DB) []models.EmbeddingInfo {
	var allEmbeddings []models.EmbeddingInfo

	query := `
		SELECT name, category, embedding FROM embedding
		`

	allRows, err := db.Query(query)
	if err != nil {
		log.Fatal("Query execution is denied", err)
	}
	defer allRows.Close()

	for allRows.Next() {
		var name, category, embeddingJSON string
		if err := allRows.Scan(&name, &category, &embeddingJSON); err != nil {
			log.Fatal("Failed to scan row", err)
		}
		var embedding []float64
		if err := json.Unmarshal([]byte(embeddingJSON), &embedding); err != nil {
			log.Fatal("Failed to unmarshal JSON", err)
		}
		allEmbeddings = append(allEmbeddings, models.EmbeddingInfo{Name: name, Category: category, Embedding: embedding})
	}

	return allEmbeddings
}

func FetchSimilarities(db *sql.DB, wrongs []string, allEmbeddings []models.EmbeddingInfo) ([]float64, []string, []string) {

	var dotProductValues []float64
	var formattedStrings []string
	var processedWrongs []string

	tx, err := db.Begin()
	if err != nil {
		log.Fatal("SQL Transaction is denied", err)
	}
	defer tx.Rollback()

	query := `SELECT * FROM embedding WHERE name = ? AND category = ?`

	for _, wrong := range wrongs {
		w_cat, w_name := common.SplitWrong(wrong)

		// SQLクエリを実行して特定の embedding を取得
		rows, err := tx.Query(query, w_name, w_cat)
		if err != nil {
			log.Fatal("Query execution is denied", err)
		}

		for rows.Next() {
			processedWrongs = append(processedWrongs, wrong)
			processedWrongs = append(processedWrongs, wrong)

			var category string
			var name string
			var embeddingJSON string

			if err := rows.Scan(&category, &name, &embeddingJSON); err != nil {
				log.Fatal("Failed to scan row", err)
			}

			var embedding []float64
			if err := json.Unmarshal([]byte(embeddingJSON), &embedding); err != nil {
				log.Fatal("Failed to unmarshal JSON", err)
			}

			type Result struct {
				Index      int
				DotProduct float64
			}

			// 内積を計算して上位2つを選択
			var dotProducts []Result
			for i, embInfo := range allEmbeddings {
				dot := common.DotProduct(embedding, embInfo.Embedding)
				if 1-dot > 0.0001 { // 内積が1でない場合のみ追加
					dotProducts = append(dotProducts, Result{Index: i, DotProduct: dot})
				}
			}

			// 内積の降順でソート
			sort.Slice(dotProducts, func(i, j int) bool {
				return dotProducts[i].DotProduct > dotProducts[j].DotProduct
			})

			// 上位2つを選択
			if len(dotProducts) > 2 {
				dotProducts = dotProducts[:2]
			}

			for _, dp := range dotProducts {
				embInfo := allEmbeddings[dp.Index]
				combinedInfo := common.FormatInfo(embInfo.Category, embInfo.Name)
				dotProductValues = append(dotProductValues, dp.DotProduct)
				formattedStrings = append(formattedStrings, combinedInfo)
			}
		}
	}

	if err := tx.Commit(); err != nil {
		fmt.Println("Error:", err)
	}

	return dotProductValues, formattedStrings, processedWrongs
}
