package analyze

import (
	"context"
	"github.com/ardikapras/owldetect/internal/app/usecase/analyze/model"
)

//IAnalyze ...
type IAnalyze interface {
	DoAnalyze(ctx context.Context, reqBody model.AnalyzeReqBody) ([]model.Match, error)
}
