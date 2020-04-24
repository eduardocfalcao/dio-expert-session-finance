package http

import (
	"net/http"
	"github.com/eduardocfalcao/dio-expert-session-finance/adapter/http/transaction"
	"github.com/eduardocfalcao/dio-expert-session-finance/adapter/http/actuator"
)

func Init() {
	http.HandleFunc("/transactions", transaction.GetTransactions)
	http.HandleFunc("/transactions/create", transaction.CreateTransaction)

	http.HandleFunc("/health", actuator.Healthcheck)

	http.ListenAndServe(":8080", nil)
}
