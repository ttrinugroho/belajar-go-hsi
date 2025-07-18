package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

var (
	ErrInvalid error
)

type ErrResponse struct {
	Error string `json:"error"`
}

type SuccesResponse struct {
	Status string `json:"status"`
}

func logRequest(r *http.Request) {
	log.WithFields(log.Fields{
		"method": r.Method,
		"url":    r.URL.String(),
	}).Info()
}

func isValid(e, a string) error {

	var messages []string
	email := strings.TrimSpace(e)
	age, err := strconv.Atoi(strings.TrimSpace(a))

	if email == "" {
		messages = append(messages, "email kosong")
	}
	if err != nil {
		messages = append(messages, "umur kosong")
		messages = append(messages, "bukan angka")
	} else if age <= 18 {
		messages = append(messages, "umur kurang dari 18 tahun")
	}

	if len(messages) == 0 {
		return nil
	}

	ErrInvalid = errors.New(strings.Join(messages, " atau "))
	return fmt.Errorf("%w", ErrInvalid)
}

func weResponse(w http.ResponseWriter, res any) {
	json.NewEncoder(w).Encode(res)
}

func handler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)

	email := r.URL.Query().Get("email")
	age := r.URL.Query().Get("age")
	err := isValid(email, age)

	w.Header().Add("content-type", "application/json")
	if errors.Is(err, ErrInvalid) {
		w.WriteHeader(http.StatusBadRequest)
		weResponse(w, ErrResponse{err.Error()})
	} else {
		weResponse(w, SuccesResponse{"ok"})
	}
}

func main() {
	http.HandleFunc("/validate", handler)
	fmt.Println("server http:localhost:8080")
	http.ListenAndServe(":8080", nil)
}
