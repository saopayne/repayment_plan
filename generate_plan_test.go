package main

import (
	"fmt"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestStringToFloat(t *testing.T) {
	assertEqual(t, stringToFloat("3444.00"), 3444.00, 	"")
	assertEqual(t, stringToFloat("3444.10"), 3444.1, 	"")
}

func TestCalculateAnnuity(t *testing.T) {
	assertEqual(t, calculateAnnuity(24, 5, 5000), 219.36, "calculate annuity method produces a wrong result")
}

func TestRoundOff(t *testing.T) {
	assertEqual(t, RoundOff(33.6677, .5, 2), 33.67, "RoundOff result is wrong")
	assertEqual(t, RoundOff(33.6677, 1, 2), 33.66, "RoundOff result is not Floor")
}

func TestGenerateDates(t *testing.T) {
	var dateBatchOne []string
	var datesBatchTwo []string
	for date := range generateDates("02.12.2017", 2) {
		dateBatchOne = append(dateBatchOne, date)
	}
	assertEqual(t, dateBatchOne[0], "2017-12-01T00:00:00Z", "The first generated date is correct")
	assertEqual(t, dateBatchOne[1], "2018-01-01T00:00:00Z", "The second generated date is correct")
	for date := range generateDates("2018-01-01T00:00:00Z", 2) {
		datesBatchTwo = append(datesBatchTwo, date)
	}
	assertEqual(t, datesBatchTwo[1], "2018-02-01T00:00:00Z", "")
}

func TestGenerateRepaymentPlan(t *testing.T) {
	repaymentPlans := GenerateRepaymentPlan(24, 5, 5000, "01.01.2018")
	assertEqual(t, repaymentPlans[0], RepaymentItem{219.36, "2018-01-01T00:00:00Z", 5000, 20.83, 198.52, 4801.47}, "")
	assertEqual(t, repaymentPlans[1], RepaymentItem{219.36, "2018-02-01T00:00:00Z", 4801.47, 20.00, 199.35, 4602.11}, "")
	assertEqual(t, repaymentPlans[len(repaymentPlans)-1], RepaymentItem{219.28, "2019-12-01T00:00:00Z", 218.37, 0.9, 218.37, 0}, "")
}