package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/charts", charts)
	http.HandleFunc("/", static)

	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func static(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static"+r.URL.Path)
}

func charts(w http.ResponseWriter, r *http.Request) {
	tracks := getHot10()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tracks)
}
