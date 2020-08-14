package v1

import (
	"github.com/gin-gonic/gin"
	"html_api/pkg/app"
	"html_api/pkg/e"
	"html_api/service/auth_service"
	"html_api/service/borrow_service"
	"html_api/service/invest_service"
	"net/http"
)

type InvestForm struct {
	BorrowId int     `form:"borrow_id"`
	Amount   float64 `form:"amount"`
	//MemberId int     `form:"member_id" valid:"Min(1)"`
}

func Invest(c *gin.Context) {
	var (
		appG = app.Gin{c}
		form InvestForm
	)

	_, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(http.StatusOK, errCode, form)
		return
	}
	memberId, _ := c.Get("member_id")
	//判断用户是否登陆
	auth := auth_service.Auth{
		ID: memberId.(int),
	}
	memberInfo, err := auth.GetMember()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	//判断用户余额
	if memberInfo.AmountUsed < form.Amount {
		appG.Response(http.StatusOK, e.ERROR_MEMBER_AMOUNT_USED_LESS, nil)
		return
	}
	/*判断标的所剩余余额和状态*/
	//获取标的信息
	borrow := borrow_service.Borrow{
		ID: form.BorrowId,
	}
	borrowInfo, err := borrow.GetBorrow()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_BORROW_DETAIL_FAIL, nil)
		return
	}
	//判断标的所剩余可投金额
	if borrowInfo.Amount-borrowInfo.AmountLimit < form.Amount {
		appG.Response(http.StatusOK, e.ERROR_MEMBER_BORROW_AMOUNT_LIMIT, nil)
		return
	}
	/*进行投资*/
	invest := invest_service.InvestInfo{
		BorrowId: form.BorrowId,
		MemberId: memberId.(int),
		Amount:   form.Amount,
	}
	if err := invest.Add(); err != nil {
		appG.Response(http.StatusOK, e.ERROR_MEMBER_INVEST_FAILED, nil)
		return
	}
	//更新用户余额
	authInfo := auth_service.Auth{
		ID:           memberId.(int),
		AmountAll:    memberInfo.AmountAll - form.Amount,
		AmountFrozen: memberInfo.AmountFrozen + form.Amount,
		AmountUsed:   memberInfo.AmountUsed - form.Amount,
	}
	if err := authInfo.UpdateMember(); err != nil {
		appG.Response(http.StatusOK, e.ERROR_MEMBER_INVEST_FAILED, nil)
		return
	}
	//更新标的剩余可投资金额
	b := borrow_service.Borrow{
		ID: form.BorrowId,
		AmountLimit: form.Amount,
	}
	if err := b.Update(); err != nil {
		appG.Response(http.StatusOK, e.ERROR_MEMBER_INVEST_FAILED, nil)
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
