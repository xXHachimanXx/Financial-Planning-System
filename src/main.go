package main

import (
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/", getTransactions)

	http.ListenAndServe(":8080", nil)
}

type Transaction struct {
	Title     string
	Amount    float32
	Type      int
	CreatedAt time.Time
}

type Transaction []Transaction

func getTransactions(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var timeLayout = "2006-01-01T15:22:22"
	time.Parse(timeLayout, "2014-01-02T15:13:22")

	var transactions = Transactions{
		Transaction {
			Title: "Salário",
			Amount: 1200.0,
			Type: 0,
			CreatedAt: ,
		}
	}

}
