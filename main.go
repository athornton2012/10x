package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/athornton2012/10x/pkg/parser"
	"github.com/athornton2012/10x/pkg/query"
)

type QueryHandler struct {
	Data []map[string]string
}

func (h QueryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	err := query.ValidateParams(h.Data, q)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		result, err := query.QueryData(q, h.Data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		res, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write(res)
	}
}

func SetupServer(filepath string) (*http.Server, error) {
	data, err := parser.ParseCSV(filepath)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("No data in CSV")
	}

	server := http.Server{Addr: ":8080"}
	http.Handle("/query", QueryHandler{
		Data: data,
	})

	return &server, nil
}

func main() {
	filepath := os.Args[1]

	server, err := SetupServer(filepath)
	if err != nil {
		panic(err)
	}

	go func () {
		if err := server.ListenAndServe(); err != nil {
		  log.Fatalf("Error HTTP Listen And Serve: %v", err)
		}
	} ()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error during HTTP Shutdown: %v", err)
	}
}
