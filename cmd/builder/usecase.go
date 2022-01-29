package builder

import (
	analyzeuc "github.com/ardikapras/owldetect/internal/app/usecase/analyze"
	analyzeuci "github.com/ardikapras/owldetect/internal/app/usecase/analyze/impl"
	"log"
)

var analyzeUsecaseCache analyzeuc.IAnalyze

// BuildAnalyzeUsecase
func BuildAnalyzeUsecase() analyzeuc.IAnalyze {
	if analyzeUsecaseCache != nil {
		return analyzeUsecaseCache
	}

	uc := analyzeuci.New()

	analyzeUsecaseCache = uc

	log.Print("init usecase analyze")
	return analyzeUsecaseCache
}
