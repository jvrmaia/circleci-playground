package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type DistanciaPayload struct {
	Points []Point `json:"points"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	var payload DistanciaPayload
	err = json.Unmarshal(body, &payload)

	if err != nil {
		log.Printf("Error parsing payload: %v", err)
		http.Error(w, "can't parse body!", http.StatusBadRequest)
		return
	}

	if len(payload.Points) != 2 {
		log.Printf("Two points required: %v", payload)
		http.Error(w, "Two points required!", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "distancia: %f\n", Distancia(payload.Points[0], payload.Points[1]))
}

func main() {
	http.HandleFunc("/dist", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
