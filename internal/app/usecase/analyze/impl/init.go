package impl

import (
	"context"
	"github.com/ardikapras/owldetect/internal/app/usecase/analyze"
	"github.com/ardikapras/owldetect/internal/app/usecase/analyze/model"
	"strings"
)

// Dependencies contain all dependencies for usecase
type analyzeUsecase struct{}

// New workflow constructor
func New() analyze.IAnalyze {
	return &analyzeUsecase{}
}

func (a analyzeUsecase) DoAnalyze(ctx context.Context, reqBody model.AnalyzeReqBody) ([]model.Match, error) {
	idx := strings.Index(reqBody.RefText, reqBody.InputText)
	if idx == -1 {
		return nil, nil
	}
	return []model.Match{
		{
			Input: model.MatchDetails{
				Text:     reqBody.InputText,
				StartIdx: 0,
				EndIdx:   len(reqBody.InputText) - 1,
			},
			Reference: model.MatchDetails{
				Text:     reqBody.RefText[idx : idx+len(reqBody.InputText)],
				StartIdx: idx,
				EndIdx:   idx + len(reqBody.InputText) - 1,
			},
		},
	}, nil
}
