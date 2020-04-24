package transaction

import (
	"encoding/json"
	"github.com/eduardocfalcao/dio-expert-session-finance/model/transaction"
	"github.com/eduardocfalcao/dio-expert-session-finance/util"
	"io/ioutil"
	"net/http"
	"fmt"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	salaryReceived := util.StringToTime("2020-04-05T11:45:26")

	var transactions = transaction.Transactions {
		transaction.Transaction{
			Title:     "Salário",
			Amount:    1200.00,
			Type:      0,
			CreatedAt: salaryReceived,
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var res = transaction.Transactions{}
	var body, _ = ioutil.ReadAll(r.Body)

	_ = json.Unmarshal(body, &res)

	fmt.Print(res)
}
