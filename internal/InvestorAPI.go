package investor

import (
	. "github.com/Udyana-Inc/InvestorAPI/internal/models"
	"github.com/jinzhu/copier"
)

func CompoundInterest(req models.InvestorRequest) (res []*models.InvestorResponse, err error) {
	endingBalance := req.beginningBalance
	interestRate := (1 + (req.interestRate / 100))
	// calculate the compound ending balance after taxes or regardless of taxes
	for years := 0; years <= req.yearsHeld; years++ {
		// copy over similar values from request to response
		err = copier.Copy(res, &req)
		if err != nil {
			return nil, err
		}
		if req.afterTax == false {
			//non taxable account
			endingBalance = interestRate * endingBalance
			res[years].endingBalance = endingBalance

		} else if req.afterTax == true {
			taxRate := 1 - (req.taxRate / 100)
			//taxable account
			endingBalance = interestRate * (endingBalance) * taxRate
			res[years].endingBalance = endingBalance

		}

	}

	return

}
