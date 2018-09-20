package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/time", TheTime).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// These names (curTime, Time, theTime, time) would be confusing in a more complex app.
// Don't follow my example here. :)
type theTime struct {
	Time string `json:"time,omitempty"`
}

var curTime []theTime

// TheTime returns the time
func TheTime(w http.ResponseWriter, r *http.Request) {
	curTime = append(curTime, theTime{Time: time.Now().String()})
	json.NewEncoder(w).Encode(curTime)
}
