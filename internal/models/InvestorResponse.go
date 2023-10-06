package models

import (
	"time"
)

type InvestorResponse struct {
	InterestRate     float64
	BeginningBalance float64
	EndingBalance    float64
	YearsHeld        int
	TaxRate          float64
	AfterTaxes       bool
	CurrentYear      time.Time
}
