// handlers/prime_handler.go
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func numberPrime(num int) bool {
	if num < 2 {
		return false
	}
	for a := 2; a*a <= num; a++ {
		if num%a == 0 {
			return false
		}
	}
	return true
}

func PrimeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Bad request method", http.StatusMethodNotAllowed)
		return
	}

	var requestData struct {
		Number int `json:"number"`
	}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestData)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	isPrimeNumber := numberPrime(requestData.Number)

	responseData := struct {
		Number int    `json:"number"`
		Prime  bool   `json:"prime"`
	}{
		Number: requestData.Number,
		Prime:  isPrimeNumber,
	}

	//Retorno
	w.Header().Set("Content-Type", "application/json")
	if isPrimeNumber {
		json.NewEncoder(w).Encode(responseData)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, toJSON(responseData))
	}
}

func toJSON(data interface{}) string {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return string(jsonData)
}
