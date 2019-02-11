package main

import (
	"encoding/json"
	"io"
	"net/http"
	"encoding/csv"
	"fmt"
	"os"
)

const (
	ExportTypeCSV  = "csv"
	ExportTypeJSON = "json"
)

func jsonWriter(w http.ResponseWriter, r *http.Request, items []RepaymentItem, exportType string) {
	if cb := r.FormValue("callback"); cb != "" {
		w.Header().Set("Content-Type", "application/javascript")
		io.WriteString(w, cb)
		w.Write([]byte("("))
		b, err := json.Marshal(items)
		if err == nil {
			w.Write(b)
		}
		io.WriteString(w, ");")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if exportType == ExportTypeCSV {
		exportResponseToCsv(items)
		json.NewEncoder(w).Encode(map[string]string{"status": fmt.Sprint(http.StatusCreated), "message": "Repayment plan successfully generated to lendico_generated_plan.csv"})
		return
	}

	json.NewEncoder(w).Encode(items)
}

func exportResponseToCsv(items []RepaymentItem) {
	f, err := os.Create("./lendico_generated_plan.csv")
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()
	csvWriter := csv.NewWriter(f)
	csvWriter.Write([]string{"Borrower Payment Amount", "Date", "Initial Outstanding Principal", "Interest", "Principal", "Remaining Outstanding Principal"})
	for _, obj := range items {
		var record []string
		record = append(record, fmt.Sprintf("%.2f", obj.BorrowerPaymentAmount))
		record = append(record, obj.Date)
		record = append(record, fmt.Sprintf("%.2f", obj.InitialOutstandingPrincipal))
		record = append(record, fmt.Sprintf("%.2f", obj.Interest))
		record = append(record, fmt.Sprintf("%.2f", obj.Principal))
		record = append(record, fmt.Sprintf("%.2f", obj.RemainingOutstandingPrincipal))
		csvWriter.Write(record)
	}

	csvWriter.Flush()
}

func GeneratePlanHandler(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var loanItem LoanItem
	err := decoder.Decode(&loanItem)

	exportType := request.URL.Query().Get("format")

	if err != nil {
		panic(err)
	}

	response := GenerateRepaymentPlan(
		loanItem.Duration,
		stringToFloat(loanItem.NominalRate),
		stringToFloat(loanItem.LoanAmount),
		loanItem.StartDate)

	jsonWriter(writer, request, response, exportType)
}
