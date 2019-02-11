package main

import (
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	MonthsInYear   = 12.0
	DaysInMonth    = 30.0
	DaysInYear     = 360.0
	DateFormat     = "01.01.2006"
	LongDateFormat = "2006-01-02T15:04:05Z"
)

type LoanItem struct {
	LoanAmount  string  `json:"loanAmount"`
	NominalRate string  `json:"nominalRate"`
	Duration    float64 `json:"duration"`
	StartDate   string  `json:"startDate"`
}

type RepaymentItem struct {
	BorrowerPaymentAmount         float64 `json:"borrowerPaymentAmount"`
	Date                          string  `json:"date"`
	InitialOutstandingPrincipal   float64 `json:"initialOutstandingPrincipal"`
	Interest                      float64 `json:"interest"`
	Principal                     float64 `json:"principal"`
	RemainingOutstandingPrincipal float64 `json:"remainingOutstandingPrincipal"`
}

func GenerateRepaymentPlan(duration float64, rate float64, iOutPrincipal float64, start string) []RepaymentItem {
	pv := getAnnuity(duration, rate, iOutPrincipal)
	rOutPrincipal, principal := 0.0, 0.0
	var repaymentItems []RepaymentItem

	for date := range generateDates(start, int(duration)) {
		if date != start {
			iOutPrincipal -= principal
		}
		roundedValue := Round(((rate*DaysInMonth*iOutPrincipal)/DaysInYear)/100, 1, 2)
		principal = pv - roundedValue
		rOutPrincipal = iOutPrincipal - principal
		if rOutPrincipal < 0 {
			pv += rOutPrincipal
			principal = pv - roundedValue
			rOutPrincipal = 0
		}
		repaymentItems = append(repaymentItems, RepaymentItem{
			pv,
			date,
			iOutPrincipal,
			roundedValue,
			removeTrailingZero(principal),
			removeTrailingZero(rOutPrincipal),
		})
	}

	return repaymentItems
}

func generateDates(t string, n int) chan string {
	ch := make(chan string)

	date, err := time.Parse(DateFormat, t)
	if err != nil {
		err = nil
		date, err = time.Parse(LongDateFormat, t)
		if err != nil {
			log.Fatal(err)
			return ch
		}
	}
	year, month, day := date.Date()
	nextMonth, nextYear := int(month), int(year)

	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			date, err = time.Parse("1.1.2006", strconv.Itoa(day)+"."+
				strconv.Itoa(nextMonth)+"."+
				strconv.Itoa(nextYear))
			if err != nil {
				log.Fatal(err)
				continue
			}
			ch <- date.Format(LongDateFormat)

			nextMonth += 1

			startMonth := 1
			if nextMonth > MonthsInYear {
				nextMonth = startMonth
				nextYear += 1
			}
		}
	}()
	return ch
}

func Round(val float64, roundOn float64, places int) (roundedValue float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	roundedValue = round / pow

	return roundedValue
}

func floatToString(inputNum float64) string {
	return strconv.FormatFloat(inputNum, 'f', 6, 64)
}

func stringToFloat(input string) float64 {
	if s, err := strconv.ParseFloat(input, 64); err == nil {
		return s
	}
	return 0
}

func removeTrailingZero(value float64) float64 { // Remove Tailing zero
	fv := floatToString(value)
	if strings.Contains(fv, ".") {
		if fv[len(fv)-1:] == "0" || fv[len(fv)-1:] == "." {
			if s, err := strconv.ParseFloat(fv[:len(fv)-1], 64); err == nil {
				return s
			}
		}
	}
	return stringToFloat(fv)
}

func getAnnuity(duration float64, rate float64, principal float64) float64 {
	m := (principal * rate / 100 / MonthsInYear) / (1 - math.Pow(1+rate/100/MonthsInYear, -duration))
	return Round(m, .5, 2)
}

