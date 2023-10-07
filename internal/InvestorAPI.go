package investor

import (
	"fmt"
	"time"

	"github.com/Udyana-Inc/InvestorAPI/internal/models"
	"github.com/jinzhu/copier"
)

func CompoundInterest(req models.InvestorRequest) (res []*models.InvestorResponse, err error) {
	endingBalance := req.BeginningBalance
	interestRate := (1 + (req.InterestRate / 100))
	taxRate := 1 - (req.TaxRate / 100)
	res = make([]*models.InvestorResponse, req.YearsHeld+1)
	// calculate the compound ending balance after taxes or regardless of taxes
	for years := 1; years <= req.YearsHeld; years++ {
		res[years] = &models.InvestorResponse{}
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
			//taxable account
			endingBalance = interestRate * (endingBalance) * taxRate
			res[years].EndingBalance = endingBalance
			res[years].CurrentYear = time.Now().AddDate(years, 0, 0)

		}

	}

	return

}

func greaterThanDesiredBalance(req models.InvestorRequest) (res []*models.InvestorResponse, err error) {
	desiredAmount := req.DesiredAmount
	req.YearsHeld = 1
	res, err = CompoundInterest(req)
	if err != nil {
		fmt.Println(err)
	}
	for res[len(res)].EndingBalance < desiredAmount {
		req.YearsHeld = +1
		res, err = CompoundInterest(req)
	}
	return
}
