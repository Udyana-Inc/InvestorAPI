package models

type InvestorRequest struct {
	interestRate     float64
	beginningBalance float64
	yearsHeld        int16
	taxRate          float64
	afterTaxes       bool
}
