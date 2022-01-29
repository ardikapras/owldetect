package main

import (
	"github.com/ardikapras/owldetect/internal/app/usecase/analyze/model"
	"log"
	"strings"
)

func main() {
	// define handlers
	//http.Handle("/", http.FileServer(http.Dir("./static")))
	//http.HandleFunc("/analysis", func(w http.ResponseWriter, r *http.Request) {
	//	// check http method
	//	if r.Method != http.MethodPost {
	//		WriteAPIResp(w, NewErrorResp(NewErrMethodNotAllowed()))
	//		return
	//	}
	//	// parse request body
	//	var reqBody analyzeReqBody
	//	err := json.NewDecoder(r.Body).Decode(&reqBody)
	//	if err != nil {
	//		WriteAPIResp(w, NewErrorResp(NewErrBadRequest(err.Error())))
	//		return
	//	}
	//	// validate request body
	//	err = reqBody.Validate()
	//	if err != nil {
	//		WriteAPIResp(w, NewErrorResp(err))
	//		return
	//	}
	//	// do analysis
	//	matches := doAnalysis(reqBody.InputText, reqBody.RefText)
	//	// output success response
	//	WriteAPIResp(w, NewSuccessResp(map[string]interface{}{
	//		"matches": matches,
	//	}))
	//})

	// initialize all dependencies and start web server
	err := startApp()
	if err != nil {
		log.Printf("failed to start app: %v", err)
		return
	}
}

func doAnalysis(input, ref string) []model.Match {
	idx := strings.Index(ref, input)
	if idx == -1 {
		return nil
	}
	return []model.Match{
		{
			Input: model.MatchDetails{
				Text:     input,
				StartIdx: 0,
				EndIdx:   len(input) - 1,
			},
			Reference: model.MatchDetails{
				Text:     ref[idx : idx+len(input)],
				StartIdx: idx,
				EndIdx:   idx + len(input) - 1,
			},
		},
	}
}
