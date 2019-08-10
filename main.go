package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bwynn/silverStrand/logs"
	"github.com/gorilla/mux"
	"github.com/robfig/cron"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/health", logs.LogWrapper(handleHealth))
	r.HandleFunc("/api/job", logs.LogWrapper(handleJob))

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8800",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

// serve health
func handleHealth(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

// fetch weather data -> weatherundergound.com/api/asdfasdj
// add new location - POST
//	- update location details
//	- add new record
// get location
// query location by timerange

// functions
// cron job to fetch weather data
func handleJob(w http.ResponseWriter, r *http.Request) {
	startJob()
	fmt.Println(fmt.Sprint("Results are here"))
}

// create a new table in DB
func startJob() {
	msg := "every minute, or set interval, dispatch our call to our service api"
	c := cron.New()
	c.AddFunc("@every 1min", func() { fmt.Println(msg) })
	c.Start()

	c.Run()
}
