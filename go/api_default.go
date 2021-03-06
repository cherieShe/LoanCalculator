/*
 * Loan Calculator
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
 package swagger

 import (
	 "net/http"
	 "encoding/json"
	 "io/ioutil"
	 "fmt"
	 "math"
 )
 
 func  calculate(calculateloanBody CalculateloanBody) (loanRepayments LoanRepayments) {
	 
	 
	 
	 loanRepayments.TotalInterestPayable = loanRepayments.MonthlyRepayments*12*calculateloanBody.LoanTerm
 
	 
	 // loanRepayments.MonthlyRepayments= int32(float64(calculateloanBody.LoanAmount)*calculateloanBody.InterestRate/12)
	 
	 if calculateloanBody.LoanType == "interest" {
		 for i:=int32(0); i < calculateloanBody.LoanTerm; i++ {
			 var loanRepaymentsAmountOwing LoanRepaymentsAmountOwing
			 
		 if calculateloanBody.PaymentFrequency=="monthly"{
			 
			 loanRepaymentsAmountOwing.Principal = calculateloanBody.LoanAmount 
			 loanRepaymentsAmountOwing.Year = i
			 loanRepayments.MonthlyRepayments= int32(float64(calculateloanBody.LoanAmount)*calculateloanBody.InterestRate/12)
			 loanRepayments.TotalInterestPayable = loanRepayments.MonthlyRepayments*12*calculateloanBody.LoanTerm
			 loanRepaymentsAmountOwing.Interest = loanRepayments.MonthlyRepayments*12*calculateloanBody.LoanTerm-i*loanRepayments.MonthlyRepayments
			 loanRepaymentsAmountOwing.Total = loanRepaymentsAmountOwing.Principal+loanRepaymentsAmountOwing.Interest-i*loanRepayments.MonthlyRepayments
			 loanRepayments.AmountOwing = append(loanRepayments.AmountOwing, loanRepaymentsAmountOwing)
	 
		 }else if calculateloanBody.PaymentFrequency=="fortnightly"{
		 
 
		 loanRepaymentsAmountOwing.Principal = calculateloanBody.LoanAmount 
			 loanRepaymentsAmountOwing.Year = i
			 loanRepayments.MonthlyRepayments= int32(math.Round( ((float64(calculateloanBody.LoanAmount)*calculateloanBody.InterestRate/26)*100)/100) )
			 loanRepayments.TotalInterestPayable = int32(math.Round(float64(calculateloanBody.LoanAmount)*calculateloanBody.InterestRate*float64(calculateloanBody.LoanTerm)))
			 loanRepaymentsAmountOwing.Interest = int32(math.Round(float64(calculateloanBody.LoanAmount)*calculateloanBody.InterestRate*float64(calculateloanBody.LoanTerm)))-i*loanRepayments.MonthlyRepayments
			 loanRepaymentsAmountOwing.Total = loanRepaymentsAmountOwing.Principal+loanRepaymentsAmountOwing.Interest-i*loanRepayments.MonthlyRepayments
			 loanRepayments.AmountOwing = append(loanRepayments.AmountOwing, loanRepaymentsAmountOwing)
		 
		 }else if calculateloanBody.PaymentFrequency=="weekly"{
			 loanRepaymentsAmountOwing.Principal = calculateloanBody.LoanAmount 
			 loanRepaymentsAmountOwing.Year = i
			 loanRepayments.MonthlyRepayments= int32(math.Round(float64(calculateloanBody.LoanAmount)*calculateloanBody.InterestRate/52))
			 loanRepayments.TotalInterestPayable = int32(math.Round(float64(calculateloanBody.LoanAmount)*calculateloanBody.InterestRate*float64(calculateloanBody.LoanTerm)))
			 loanRepaymentsAmountOwing.Interest = int32(math.Round(float64(calculateloanBody.LoanAmount)*calculateloanBody.InterestRate*float64(calculateloanBody.LoanTerm)))-i*loanRepayments.MonthlyRepayments
			 loanRepaymentsAmountOwing.Total = loanRepaymentsAmountOwing.Principal+loanRepaymentsAmountOwing.Interest-i*loanRepayments.MonthlyRepayments
			 loanRepayments.AmountOwing = append(loanRepayments.AmountOwing, loanRepaymentsAmountOwing)
		 }
		 
		 }
	 }
 
	 
	 
	 return loanRepayments
 }
 
 func CalculateLoan(w http.ResponseWriter, r *http.Request) {
	 w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	 w.WriteHeader(http.StatusOK)
	 // j,_:=json.Marshal("hello")
 
	 //     w.Write(j)
 
	 requestBodyBytes, _ := ioutil.ReadAll(r.Body)
	 var newPerson CalculateloanBody
	//  json.Unmarshal(requestBodyBytes, &newPerson)// 
	 err1 := json.Unmarshal(requestBodyBytes, &newPerson)
	fmt.Println(err1) 
 //1.convert request body byte to CalculateloanBody type 
 //2.and assign properties values to newPerson
 
	 
	 
	 
 
	//  var res LoanRepayments
	 res := calculate(newPerson)
	 jresponse,_:=json.Marshal(&res) //convert CalculateloanBody type to string for sending back
 
	 err2, _ := w.Write(jresponse)
	   fmt.Println(err2) 
 
 }
 