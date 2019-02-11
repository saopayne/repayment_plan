#### Steps to run the application:

1. Run "go get -u github.com/gorilla/mux"
2. run "sh build.sh"
3. curl POST -d {"loanAmount": "5000","nominalRate": "5.0","duration": 24, "startDate": "2018-01-01T00:00:01Z"} http://localhost:8080/generate-plan?format=json

   curl POST -d {"loanAmount": "5000","nominalRate": "5.0","duration": 24, "startDate": "2018-01-01T00:00:01Z"} http://localhost:8080/generate-plan?format=csv
   
   
#### Task Description
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