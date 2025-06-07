package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

type Name struct {
	ID   int    `json:"id"`
	Name string `json:"Name"`
}

type Response struct {
	Status int  `json:"status"`
	Data   Name `json:"data"`
}

func homePostPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received at /post")

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		return
	}

	var res Name
	err := json.NewDecoder(r.Body).Decode(&res)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		fmt.Println("JSON decode error:", err)
		return
	}

	fmt.Println("The response:", res)

	// Create structured response
	response := Response{
		Status: 200,
		Data:   res,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
func handleGetPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received at /get")

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}

	getId := 2 // You can later get this from query parameters

	// Sample data
	res := []Name{
		{ID: 1, Name: "Varun"},
		{ID: 2, Name: "Ram"},
	}

	// Filtered logic example (optional)
	var filtered []Name
	for _, item := range res {
		if item.ID == getId {
			filtered = append(filtered, item)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(filtered) // or just `res` if you want to return all
}

func handle_html(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received at /get_html")

	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}

	// Path to the HTML file
	htmlPath := filepath.Join("static", "index.html")

	http.ServeFile(w, r, htmlPath)
}

func RunServer() {
	fmt.Println("Started a Go Server on port 8080")
	http.HandleFunc("/post", homePostPageHandler)
	http.HandleFunc("/get", handleGetPageHandler)
	http.HandleFunc("/get_html", handle_html)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
