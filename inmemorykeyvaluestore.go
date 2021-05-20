package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	db     = map[string]interface{}{}
	dbLock sync.Mutex
)

type Entry struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func sendResponse(entry *Entry, w http.ResponseWriter) {
	w.Header().Set("ContentType", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(entry); err != nil {
		log.Printf("error encoding %+v - %s", entry, err)
	}
}

func addKey(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	entry := &Entry{}
	if err := dec.Decode(entry); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	dbLock.Lock()
	defer dbLock.Unlock()
	db[entry.Key] = entry.Value

	sendResponse(entry, w)
}

func getKey(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[4:]
	dbLock.Lock()
	defer dbLock.Unlock()
	value, ok := db[key]
	if !ok {
		http.Error(w, fmt.Sprintf("Key %q not found", key), http.StatusNotFound)
		return
	}
	entry := &Entry{
		Key:   key,
		Value: value,
	}
	sendResponse(entry, w)
}
func main() {
	http.HandleFunc("/db", addKey)
	http.HandleFunc("/db/", getKey)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
