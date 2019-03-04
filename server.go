package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// TransactionRequest defines the request body of transaction request.
type TransactionRequest struct {
	ID       string
	PID      string
	Birthday string
	Type     int64 // 0=unknown, 1=health insurance provided, 2=self-provided
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/form.html")
	t.Execute(w, nil)
}

func failHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/fail.html")
	t.Execute(w, nil)
}

func transactionHandler(w http.ResponseWriter, r *http.Request) {
	fID := r.FormValue("id")
	fPID := r.FormValue("pid")
	fBirthday := r.FormValue("birthday")
	fType, err := strconv.ParseInt(r.FormValue("type"), 10, 64)
	if fID == "" || fPID == "" || fBirthday == "" || fType == 0 || err != nil {
		failHandler(w, r)
		return
	}
	fmt.Printf("%v", r)
	t, _ := template.ParseFiles("templates/result.html")
	t.Execute(w, &TransactionRequest{
		ID:       fID,
		PID:      fPID,
		Type:     fType,
		Birthday: fBirthday,
	})
}

func main() {
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/transaction", transactionHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
