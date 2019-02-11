package main

import (
	"log"
	"math"
	"strconv"
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

func GenerateRepaymentPlan(duration float64, nominalRate float64, iOutPrincipal float64, start string) []RepaymentItem {
	annuity := calculateAnnuity(duration, nominalRate, iOutPrincipal)
	rOutPrincipal, principal := 0.0, 0.0
	var repaymentItems []RepaymentItem

	for date := range generateDates(start, int(duration)) {
		if date != start {
			iOutPrincipal -= principal
		}
		interest := (nominalRate * DaysInMonth * iOutPrincipal) / DaysInYear / 100.00
		principal = annuity - interest
		rOutPrincipal = iOutPrincipal - principal
		if rOutPrincipal < 0 {
			annuity += rOutPrincipal
			principal = annuity - interest
			rOutPrincipal = 0
		}
		repaymentItems = append(repaymentItems, RepaymentItem{
			RoundOff(annuity, 1, 2),
			date,
			RoundOff(iOutPrincipal, 1, 2),
			RoundOff(interest, 1, 2),
			RoundOff(principal, 1, 2),
			RoundOff(rOutPrincipal, 1, 2),
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

func RoundOff(val float64, roundOn float64, decimalPlaces int) (roundedValue float64) {
	var round float64
	pow := math.Pow(10, float64(decimalPlaces))
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

func stringToFloat(input string) float64 {
	if s, err := strconv.ParseFloat(input, 64); err == nil {
		return s
	}
	return 0
}

func calculateAnnuity(duration float64, nominalRate float64, principal float64) float64 {
	m := (principal * nominalRate / 100.00 / MonthsInYear) / (1 - math.Pow(1+nominalRate/100.00/MonthsInYear, -duration))
	return RoundOff(m, .5, 2)
}
