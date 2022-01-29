package model

import "github.com/ardikapras/owldetect/internal/pkg"

type (
	AnalyzeReqBody struct {
		InputText string `json:"input_text"`
		RefText   string `json:"ref_text"`
	}

	Match struct {
		Input     MatchDetails `json:"input"`
		Reference MatchDetails `json:"ref"`
	}

	MatchDetails struct {
		Text     string `json:"text"`
		StartIdx int    `json:"start_idx"`
		EndIdx   int    `json:"end_idx"`
	}
)

func (rb AnalyzeReqBody) Validate() error {
	if len(rb.InputText) == 0 {
		return pkg.NewErrBadRequest("missing `input_text`")
	}
	if len(rb.RefText) == 0 {
		return pkg.NewErrBadRequest("missing `ref_text`")
	}
	if len(rb.InputText) > len(rb.RefText) {
		return pkg.NewErrBadRequest("`ref_text` must be longer than `input_text`")
	}
	return nil
}
