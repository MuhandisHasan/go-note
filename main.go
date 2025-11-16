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

	// note := Note{Title: "First Note", Content: "Hello this is first note"}

	// if err := SaveNotes(&note); err != nil {
	// 	log.Fatal(err)
	// }

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello World")
	// })

	http.HandleFunc("/notes", func(w http.ResponseWriter, r *http.Request) {

		notes, _ := GetNotes()

		resp := Response{
			Status: "success",
			Code:   200,
			Data:   notes,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)

	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
