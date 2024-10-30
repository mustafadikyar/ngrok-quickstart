package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	ClientIP  string    `json:"client_ip"`
	Protocol  string    `json:"protocol"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, "Hoş geldiniz! Bu sayfa Ngrok üzerinden erişildi.")
	})

	http.HandleFunc("/api/info", func(w http.ResponseWriter, r *http.Request) {
		response := Response{
			Message:   "API başarıyla çalışıyor",
			Timestamp: time.Now(),
			ClientIP:  r.RemoteAddr,
			Protocol:  r.Proto,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("Server 3000 portunda başlatılıyor...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
