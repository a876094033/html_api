package v1

import (
	"github.com/gin-gonic/gin"
	"html_api/pkg/app"
	"html_api/pkg/e"
	"html_api/service/borrow_service"
	"html_api/service/repay_service"
	"net/http"
)

func Repay(c *gin.Context) {
	var (
		appG = app.Gin{c}
	)
	//生成还款
	//1查询借款满标
	borrow := borrow_service.Borrow{
		AmountLimit:  0,
		BorrowStatus: 2,
	}
	borrows, err := borrow.GetBorrowsRepay()
	if err != nil || len(borrows) <= 0 {
		appG.Response(http.StatusOK, e.ERROR_BORROW_REPAY_CREATE_NONE, "77777777777")
	}
	for _, v := range borrows {
		calculation := &app.Calculation{
			Amount:    v.Amount,
			Term:      v.Term,
			TermType:  v.TermType,
			RepayType: v.RepayType,
			Interest:  v.InterestRate / 100,
		}
		//2计算还款本息
		interest, _ := calculation.CalcBorrowInterest()
		if len(interest) > 0 {
			for _, i := range interest {
				//3 生成还款本息
				repay := repay_service.BorrowRepay{
					BorrowId:      v.ID,
					Period:        i.Period,
					Capital:       i.Capital,
					Interest:      i.Interest,
					RepayStatus:   0,
					RepayTime:     i.RepayTime,
				}
				if err := repay.AddBorrowRepay(); err != nil {
					appG.Response(http.StatusOK, e.SUCCESS, err)
				}
			}
		}
	}

	//4更新标的状态

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
