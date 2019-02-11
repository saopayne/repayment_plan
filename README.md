### Task Description
In order to inform borrowers about the final repayment schedule, 
we need to have precalculated repayment plans throughout the life time of a loan.
    
To be able to calculate a repayment plan specific input parameters are necessary:
   
   - duration (number of instalments in months)
   - nominal interest rate
   - total loan amount ("total principal amount")
   - Date of Disbursement/Payout
These four parameters need to be input parameters.
    
The goal is to calculate a repayment plan for an annuity loan. Therefore the amount that the
borrower has to payback every month, consisting principal and interest repayments, does not change (the last instalment might be an exception).

The annuity amount has to be derived from three of the input parameters (duration,nominal interest rate, total loan amount) before starting the plan calculation.(use http://financeformulas.net/Annuity_Payment_Formula.html as reference)

### Installation

Clone this repository:

```sh
$ git clone https://github.com/saopayne/repayment_plan.git
```

### Running the application

Start the Go application

```sh
$ cd /repayment_plan
$ sh build.sh
```

### Making a HTTP POST request

Using any HTTP client, make a POST request to the URL:

```
http://localhost:8080/generate-plan?format=csv
```

Specify the payload of the POST request with a sample below:

```
{
  "loanAmount": "5000",
  "nominalRate": "5.0",
  "duration": 24,
  "startDate": "2018-01-01T00:00:01Z"
}
```

Or using the effective CURL:

```sh
curl -d '{"loanAmount": "5000","nominalRate": "5.0","duration": 24, "startDate": "2018-01-01T00:00:01Z"}' -X POST http://localhost:8080/generate-plan?format=json
curl -d '{"loanAmount": "5000","nominalRate": "5.0","duration": 24, "startDate": "2018-01-01T00:00:01Z"}' -X POST http://localhost:8080/generate-plan?format=csv
```


 
##### A sample response:

   
```json
[
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2018-01-01T00:00:00Z",
        "initialOutstandingPrincipal": 5000,
        "interest": 20.83,
        "principal": 198.52,
        "remainingOutstandingPrincipal": 4801.47
    },
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2018-02-01T00:00:00Z",
        "initialOutstandingPrincipal": 4801.47,
        "interest": 20,
        "principal": 199.35,
        "remainingOutstandingPrincipal": 4602.11
    },
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2018-03-01T00:00:00Z",
        "initialOutstandingPrincipal": 4602.11,
        "interest": 19.17,
        "principal": 200.18,
        "remainingOutstandingPrincipal": 4401.93
    },
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2018-04-01T00:00:00Z",
        "initialOutstandingPrincipal": 4401.93,
        "interest": 18.34,
        "principal": 201.01,
        "remainingOutstandingPrincipal": 4200.91
    },
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2018-05-01T00:00:00Z",
        "initialOutstandingPrincipal": 4200.91,
        "interest": 17.5,
        "principal": 201.85,
        "remainingOutstandingPrincipal": 3999.06
    },
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2018-06-01T00:00:00Z",
        "initialOutstandingPrincipal": 3999.06,
        "interest": 16.66,
        "principal": 202.69,
        "remainingOutstandingPrincipal": 3796.36
    },
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2018-07-01T00:00:00Z",
        "initialOutstandingPrincipal": 3796.36,
        "interest": 15.81,
        "principal": 203.54,
        "remainingOutstandingPrincipal": 3592.82
    },
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2018-08-01T00:00:00Z",
        "initialOutstandingPrincipal": 3592.82,
        "interest": 14.97,
        "principal": 204.38,
        "remainingOutstandingPrincipal": 3388.43
    },
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2018-09-01T00:00:00Z",
        "initialOutstandingPrincipal": 3388.43,
        "interest": 14.11,
        "principal": 205.24,
        "remainingOutstandingPrincipal": 3183.18
    },
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2018-10-01T00:00:00Z",
        "initialOutstandingPrincipal": 3183.18,
        "interest": 13.26,
        "principal": 206.09,
        "remainingOutstandingPrincipal": 2977.09
    },
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2018-11-01T00:00:00Z",
        "initialOutstandingPrincipal": 2977.09,
        "interest": 12.4,
        "principal": 206.95,
        "remainingOutstandingPrincipal": 2770.13
    },
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2018-12-01T00:00:00Z",
        "initialOutstandingPrincipal": 2770.13,
        "interest": 11.54,
        "principal": 207.81,
        "remainingOutstandingPrincipal": 2562.31
    },
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2019-01-01T00:00:00Z",
        "initialOutstandingPrincipal": 2562.31,
        "interest": 10.67,
        "principal": 208.68,
        "remainingOutstandingPrincipal": 2353.63
    },
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2019-02-01T00:00:00Z",
        "initialOutstandingPrincipal": 2353.63,
        "interest": 9.8,
        "principal": 209.55,
        "remainingOutstandingPrincipal": 2144.08
    },
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2019-03-01T00:00:00Z",
        "initialOutstandingPrincipal": 2144.08,
        "interest": 8.93,
        "principal": 210.42,
        "remainingOutstandingPrincipal": 1933.65
    },
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2019-04-01T00:00:00Z",
        "initialOutstandingPrincipal": 1933.65,
        "interest": 8.05,
        "principal": 211.3,
        "remainingOutstandingPrincipal": 1722.35
    },
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2019-05-01T00:00:00Z",
        "initialOutstandingPrincipal": 1722.35,
        "interest": 7.17,
        "principal": 212.18,
        "remainingOutstandingPrincipal": 1510.16
    },
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2019-06-01T00:00:00Z",
        "initialOutstandingPrincipal": 1510.16,
        "interest": 6.29,
        "principal": 213.06,
        "remainingOutstandingPrincipal": 1297.1
    },
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2019-07-01T00:00:00Z",
        "initialOutstandingPrincipal": 1297.1,
        "interest": 5.4,
        "principal": 213.95,
        "remainingOutstandingPrincipal": 1083.14
    },
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2019-08-01T00:00:00Z",
        "initialOutstandingPrincipal": 1083.14,
        "interest": 4.51,
        "principal": 214.84,
        "remainingOutstandingPrincipal": 868.3
    },
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2019-09-01T00:00:00Z",
        "initialOutstandingPrincipal": 868.3,
        "interest": 3.61,
        "principal": 215.74,
        "remainingOutstandingPrincipal": 652.55
    },
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2019-10-01T00:00:00Z",
        "initialOutstandingPrincipal": 652.55,
        "interest": 2.71,
        "principal": 216.64,
        "remainingOutstandingPrincipal": 435.91
    },
    {
        "borrowerPaymentAmount": 219.36,
        "date": "2019-11-01T00:00:00Z",
        "initialOutstandingPrincipal": 435.91,
        "interest": 1.81,
        "principal": 217.54,
        "remainingOutstandingPrincipal": 218.37
    },
    {
        "borrowerPaymentAmount": 219.28,
        "date": "2019-12-01T00:00:00Z",
        "initialOutstandingPrincipal": 218.37,
        "interest": 0.9,
        "principal": 218.37,
        "remainingOutstandingPrincipal": 0
    }
]
```

A cherry :), there's an optional query param (?format=csv) that can be used to export as CSV.

  