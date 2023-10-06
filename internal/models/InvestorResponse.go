package models

type InvestorResponse struct {
	interestRate     float64
	beginningBalance float64
	endingBalance    float64
	yearsHeld        int16
	taxRate          float64
	afterTaxes       bool
}
