package analyze

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ardikapras/owldetect/internal/app/usecase/analyze"
	"github.com/ardikapras/owldetect/internal/app/usecase/analyze/model"
	"github.com/ardikapras/owldetect/internal/pkg"
	"net/http"
)

// HTTPHandler -
type HTTPHandler interface {
	HandleAnalyze(w http.ResponseWriter, req *http.Request)
}

// Dependencies contain all dependencies for handler
type dependencies struct {
	analyzeUC analyze.IAnalyze
}

// New constructor
func New(analyzeUC analyze.IAnalyze) HTTPHandler {
	return &dependencies{
		analyzeUC: analyzeUC,
	}
}

func (d dependencies) HandleAnalyze(w http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	// check http method
	if req.Method != http.MethodPost {
		pkg.WriteAPIResp(w, pkg.NewErrorResp(pkg.NewErrMethodNotAllowed()))
		return
	}
	// parse request body
	var reqBody model.AnalyzeReqBody
	err := json.NewDecoder(req.Body).Decode(&reqBody)
	if err != nil {
		pkg.WriteAPIResp(w, pkg.NewErrorResp(pkg.NewErrBadRequest(err.Error())))
		return
	}
	// validate request body
	err = reqBody.Validate()
	if err != nil {
		pkg.WriteAPIResp(w, pkg.NewErrorResp(err))
		return
	}
	// do analysis
	matches, err := d.analyzeUC.DoAnalyze(ctx, reqBody)
	if err != nil {
		pkg.WriteAPIResp(w, pkg.NewErrorResp(err))
		return
	}
	fmt.Printf("matches::%+v\n", matches)
	// output success response
	pkg.WriteAPIResp(w, pkg.NewSuccessResp(map[string]interface{}{
		"matches": matches,
	}))
}
