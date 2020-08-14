package v1

import (
	"github.com/gin-gonic/gin"
	"html_api/pkg/app"
	"html_api/pkg/e"
	"html_api/service/recharge_service"
	"net/http"
)

type rechargeForm struct {
	Amount     float64 `form:"amount""` // valid:"Required;Min(1000)
	MemberId   int     `form:"member_id"`
	CardNumber string  `form:"card_number"`
	Bank       string  `form:"bank"`
}

func AddRecharge(c *gin.Context) {
	var (
		appG = app.Gin{c}
		form rechargeForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, form)
		return
	}
	member_id, _ := c.Get("member_id")
	recharge := recharge_service.Recharge{
		Amount:   form.Amount,
		MemberId: member_id.(int),
	}
	if err := recharge.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_RECHARGE, nil)
		return
	}
	//绑卡
	card := recharge_service.Card{
		MemberId:   member_id.(int),
		CardNumber: form.CardNumber,
		Bank:       form.Bank,
	}
	if err := card.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_RECHARGE, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
