package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./urls.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table if not exists
	createTable := `
	CREATE TABLE IF NOT EXISTS urls (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		short_code TEXT UNIQUE,
		long_url TEXT
	);
	`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}

	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/", redirectHandler)

	fmt.Println("🚀 URL Shortener running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// ================= HANDLER: SHORTEN =================
func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	type Request struct {
		URL string `json:"url"`
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.URL == "" {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	longURL := strings.TrimSpace(req.URL)

	// Add https:// if missing
	if !strings.HasPrefix(longURL, "http://") && !strings.HasPrefix(longURL, "https://") {
		longURL = "https://" + longURL
	}

	shortCode := generateShortCode()

	// Save to DB
	_, err = db.Exec("INSERT INTO urls (short_code, long_url) VALUES (?, ?)", shortCode, longURL)
	if err != nil {
		http.Error(w, "Failed to save URL", http.StatusInternalServerError)
		return
	}

	shortURL := fmt.Sprintf("http://localhost:8080/%s", shortCode)

	response := map[string]string{
		"original_url": longURL,
		"short_url":    shortURL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// ================= HANDLER: REDIRECT =================
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	code := strings.TrimPrefix(r.URL.Path, "/")
	if code == "" {
		fmt.Fprintln(w, "Welcome to the Go URL Shortener! Use POST /shorten with a JSON body like {\"url\":\"https://example.com\"}")
		return
	}

	var longURL string
	err := db.QueryRow("SELECT long_url FROM urls WHERE short_code = ?", code).Scan(&longURL)
	if err != nil {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}

// ================= HELPER FUNCTION =================
func generateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	code := make([]byte, 6)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}
