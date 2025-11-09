package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {
	var url string
	fmt.Print("Enter a website URL (e.g. google.com or https://google.com): ")
	fmt.Scan(&url)

	// Add "https://" if missing
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "https://" + url
		fmt.Println("Formatted URL:", url)
	}

	fmt.Printf("Checking connectivity to %s ...\n", url)

	client := http.Client{
		Timeout: 5 * time.Second, // Avoid hanging forever
	}

	start := time.Now()
	resp, err := client.Get(url)
	elapsed := time.Since(start)

	if err != nil {
		fmt.Println("❌ Unable to connect to the website.")
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("✅ Connected successfully to %s\n", url)
		fmt.Printf("Response Time: %v\n", elapsed)
	} else {
		fmt.Printf("⚠️ Received status code %d from %s\n", resp.StatusCode, url)
	}
}
