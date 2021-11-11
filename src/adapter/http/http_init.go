package http

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/xXHachimanXx/Financial-Planning-System.git/src/adapter/http/actuator"
	"github.com/xXHachimanXx/Financial-Planning-System.git/src/adapter/http/transaction"
)

func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateATransaction)
	http.HandleFunc("/health", actuator.Health)
	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
