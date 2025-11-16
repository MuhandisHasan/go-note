package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// API response wrapper
type Response struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Data   []Note `json:"data"`
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from", r)
		}
	}()

	http.HandleFunc("/lol/{id}", func(w http.ResponseWriter, r *http.Request) {

		id := r.PathValue("id")

		fmt.Fprintf(w, "Hello World "+id)
	})

	http.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {

			notes, _ := GetNotes()

			resp := Response{
				Status: "success",
				Code:   200,
				Data:   notes,
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(resp)

		} else if r.Method == http.MethodPost {
			var data Note

			err := json.NewDecoder(r.Body).Decode(&data)

			if err != nil {
				http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
				return
			}

			SaveNotes(&data)

		}

	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
