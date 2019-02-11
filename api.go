package main

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	ExportTypeCSV  = "csv"
	ExportTypeJSON = "json"
)

func jsonWriter(w http.ResponseWriter, r *http.Request, d []RepaymentItem, exportType string) {
	if cb := r.FormValue("callback"); cb != "" {
		w.Header().Set("Content-Type", "application/javascript")
		io.WriteString(w, cb)
		w.Write([]byte("("))
		b, err := json.Marshal(d)
		if err == nil {
			w.Write(b)
		}
		io.WriteString(w, ");")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(d)
}

func GeneratePlanHandler(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var loanItem LoanItem
	err := decoder.Decode(&loanItem)

	if err != nil {
		panic(err)
	}

	response := GenerateRepaymentPlan(
		loanItem.Duration,
		stringToFloat(loanItem.NominalRate),
		stringToFloat(loanItem.LoanAmount),
		loanItem.StartDate)

	jsonWriter(writer, request, response, ExportTypeJSON)
}
