package services

import (
	"github.com/dot_p/syojin/models"
)

type UserInfoServicer interface {
	GetSubmissionService(name string) ([]int, []string)
	GetPerformanceService(name string) ([]int, []int, []string)
	CalEmbeddingService(wrongs []string) ([]float64, []string, []string)
}

type AlgoSearchServicer interface {
	AlgoSearchServicer(optionInt int) ([]models.AlgoData, error)
}
