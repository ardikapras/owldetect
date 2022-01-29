package main

import (
	"github.com/ardikapras/owldetect/cmd/builder"
	"github.com/ardikapras/owldetect/internal/app/interface/handler/analyze"
)

type handlers struct {
	analyzeHandler analyze.HTTPHandler
}

func startApp() error {
	// initialize handler
	analyzeHandler := analyze.New(builder.BuildAnalyzeUsecase())

	// web server & handlers
	handler := handlers{
		analyzeHandler: analyzeHandler,
	}

	return startServer(handler)
}
