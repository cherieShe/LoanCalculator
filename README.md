# LoanCalculator
Google app engine url for this repo:
https://vertical-task-332900.ts.r.appspot.com/calculate-loan



The following input parameters were used in postman to check if the app is running
{
 "loan_amount": 350000,
 "loan_type": "interest",
 "payment_frequency": "monthly",
 "interest_rate": 0.03,
 "loan_term": 30
}

The ouput from the postman:
{
  "monthly_repayments": 875,
  "total_interest_payable": 315000,
  
  "amount_owing": [
        {
            "interest": 315000,
            "principal": 350000,
            "total": 665000
        },
        
        {
            "year": 1,
            "interest": 314125,
            "principal": 350000,
            "total": 663250
        },
        {
            "year": 2,
            "interest": 313250,
            "principal": 350000,
            "total": 661500
        }...
        
