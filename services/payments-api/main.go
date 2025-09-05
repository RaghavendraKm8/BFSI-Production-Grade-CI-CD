package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Payment struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func getPayment(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	p := Payment{ID: id, Amount: 100.0, Status: "SUCCESS"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func main() {
	http.HandleFunc("/healthz", healthz)
	http.HandleFunc("/v1/payments", getPayment)
	log.Println("Starting payments-api on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
