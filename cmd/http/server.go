package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func startServer(handler handlers) error {
	fatalChan := make(chan error)
	// define port, we need to set it as env for Heroku deployment
	port := os.Getenv("PORT")
	if port == "" {
		port = "9056"
	}

	// run server
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/analysis", handler.analyzeHandler.HandleAnalyze)
	log.Printf("server is listening on :%v", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("unable to run server due: %v", err)
	}
	// Check For Interupt Signal
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	select {
	case <-term:
		log.Fatalf("signal terminate detected")
	case err := <-fatalChan:
		log.Fatalf("application failed to run because %s", err.Error())
	}

	return nil
}
