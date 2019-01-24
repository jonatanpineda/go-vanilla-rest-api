package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Item struct {
	Id    string  `json:id`
	Name  string  `json:name`
	Price float32 `json:price`
}

func main() {
	db := Database{
		"1": Item{
			Id:    "1",
			Name:  "Shoes",
			Price: 50,
		},
		"2": Item{
			Id:    "2",
			Name:  "Socks",
			Price: 5,
		},
	}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe(":8000", mux))
}

type Database map[string]Item

func (db Database) list(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	items := []Item{}
	for _, item := range db {
		items = append(items, item)
	}
	json.NewEncoder(w).Encode(items)
}

func (db Database) price(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	item := req.URL.Query().Get("item")
	result, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	json.NewEncoder(w).Encode(result)
}
