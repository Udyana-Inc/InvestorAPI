package investor

import (
	"time"

	"github.com/Udyana-Inc/InvestorAPI/internal/models"
	"github.com/jinzhu/copier"
)

func CompoundInterest(req models.InvestorRequest) (res []*models.InvestorResponse, err error) {
	endingBalance := req.BeginningBalance
	interestRate := (1 + (req.InterestRate / 100))

	// calculate the compound ending balance after taxes or regardless of taxes
	for years := 1; years <= req.YearsHeld; years++ {
		// copy over similar values from request to response
		err = copier.Copy(res[years], &req)
		if err != nil {
			return nil, err
		}
		if req.AfterTaxes == false {
			//non taxable account
			endingBalance = interestRate * endingBalance
			res[years].EndingBalance = endingBalance
			res[years].CurrentYear = time.Now().AddDate(years, 0, 0)

		} else if req.AfterTaxes == true {
			taxRate := 1 - (req.TaxRate / 100)
			//taxable account
			endingBalance = interestRate * (endingBalance) * taxRate
			res[years].EndingBalance = endingBalance

		}

	}

	return

}
