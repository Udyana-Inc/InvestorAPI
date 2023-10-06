package main

import (
	"fmt"

	investor "github.com/Udyana-Inc/InvestorAPI/internal"
	"github.com/Udyana-Inc/InvestorAPI/internal/models"
)

func main() {
	request := models.InvestorRequest{InterestRate: 8, BeginningBalance: 180000.00, YearsHeld: 10, TaxRate: 10, AfterTaxes: false}
	fmt.Println("main app")
	fmt.Println(investor.CompoundInterest(request))
}
