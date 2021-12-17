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
        },
        {
            "year": 3,
            "interest": 312375,
            "principal": 350000,
            "total": 659750
        },
        {
            "year": 4,
            "interest": 311500,
            "principal": 350000,
            "total": 658000
        },
        {
            "year": 5,
            "interest": 310625,
            "principal": 350000,
            "total": 656250
        },
        {
            "year": 6,
            "interest": 309750,
            "principal": 350000,
            "total": 654500
        },
        {
            "year": 7,
            "interest": 308875,
            "principal": 350000,
            "total": 652750
        },
        {
            "year": 8,
            "interest": 308000,
            "principal": 350000,
            "total": 651000
        },
        {
            "year": 9,
            "interest": 307125,
            "principal": 350000,
            "total": 649250
        },
        {
            "year": 10,
            "interest": 306250,
            "principal": 350000,
            "total": 647500
        },
        {
            "year": 11,
            "interest": 305375,
            "principal": 350000,
            "total": 645750
        },
        {
            "year": 12,
            "interest": 304500,
            "principal": 350000,
            "total": 644000
        },
        {
            "year": 13,
            "interest": 303625,
            "principal": 350000,
            "total": 642250
        },
        {
            "year": 14,
            "interest": 302750,
            "principal": 350000,
            "total": 640500
        },
        {
            "year": 15,
            "interest": 301875,
            "principal": 350000,
            "total": 638750
        },
        {
            "year": 16,
            "interest": 301000,
            "principal": 350000,
            "total": 637000
        },
        {
            "year": 17,
            "interest": 300125,
            "principal": 350000,
            "total": 635250
        },
        {
            "year": 18,
            "interest": 299250,
            "principal": 350000,
            "total": 633500
        },
        {
            "year": 19,
            "interest": 298375,
            "principal": 350000,
            "total": 631750
        },
        {
            "year": 20,
            "interest": 297500,
            "principal": 350000,
            "total": 630000
        },
        {
            "year": 21,
            "interest": 296625,
            "principal": 350000,
            "total": 628250
        },
        {
            "year": 22,
            "interest": 295750,
            "principal": 350000,
            "total": 626500
        },
        {
            "year": 23,
            "interest": 294875,
            "principal": 350000,
            "total": 624750
        },
        {
            "year": 24,
            "interest": 294000,
            "principal": 350000,
            "total": 623000
        },
        {
            "year": 25,
            "interest": 293125,
            "principal": 350000,
            "total": 621250
        },
        {
            "year": 26,
            "interest": 292250,
            "principal": 350000,
            "total": 619500
        },
        {
            "year": 27,
            "interest": 291375,
            "principal": 350000,
            "total": 617750
        },
        {
            "year": 28,
            "interest": 290500,
            "principal": 350000,
            "total": 616000
        },
        {
            "year": 29,
            "interest": 289625,
            "principal": 350000,
            "total": 614250
        }
    ]
}
