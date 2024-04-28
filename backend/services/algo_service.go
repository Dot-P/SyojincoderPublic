package services

import (
	"github.com/dot_p/syojin/models"
	"github.com/dot_p/syojin/repositories"
)

func (s *MyAppService) AlgoSearchServicer(optionInt int) ([]models.AlgoData, error) {

	var results []models.AlgoData

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

	results, err := repositories.FetchAlgoData(s.db, classifierMapping, optionInt)

	return results, err
}
