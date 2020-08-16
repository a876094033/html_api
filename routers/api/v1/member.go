package v1

import (
	"github.com/gin-gonic/gin"
	"html_api/pkg/app"
	"html_api/pkg/e"
	"html_api/service/auth_service"
	"html_api/service/invest_service"
	"html_api/service/repay_service"
	"math"
	"net/http"
)

func Member(c *gin.Context) {
	var (
		appG = app.Gin{c}
	)
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

	appG.Response(http.StatusOK, e.SUCCESS, memberInfo)
}

type InvestList struct {
	PageSize int `form:"page_size"`
	Page     int `form:"page"`
}

type investAll struct {
	ID              int
	BorrowName      string
	BorrowStatus    int
	InterestRate    float64
	InvestAmount    float64
	Term            int
	TermType        int
	RepayTime       string
	RepayAmount     float64
	RepayAmountWait float64
	RepayTerm       int
}

func MemberInvest(c *gin.Context) {
	var (
		appG = app.Gin{c}
		form InvestList
	)
	memberId, _ := c.Get("member_id")
	_, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(http.StatusOK, e.ERROR_MEMBER_INVEST_LIST, nil)
		return
	}
	investService := invest_service.InvestInfo{
		MemberId: memberId.(int),
	}
	investList, err := investService.GetMemberInvest(form.Page, form.PageSize)
	if err != nil {
		if err != nil {
			appG.Response(http.StatusOK, e.ERROR_MEMBER_INVEST_LIST, nil)
			return
		}
	}
	investAll := make([]investAll, form.PageSize)
	if len(investList) > 0 {
		for i, v := range investList {
			investAll[i].BorrowStatus = v.BorrowStatus
			investAll[i].BorrowName = v.BorrowName
			investAll[i].Term = v.Term
			investAll[i].ID = v.ID
			investAll[i].InterestRate = v.InterestRate
			investAll[i].InvestAmount = math.Ceil(v.InvestAmount)
			investAll[i].TermType = v.TermType
			//获取借款列表
			repay := repay_service.BorrowRepay{
				BorrowId: v.BorrowId,
			}
			repays, err := repay.GetBorrowRepays()
			var capital float64
			if err == nil {
				for _, va := range repays {
					investAll[i].RepayAmount += va.Capital + va.Interest
					if va.RepayStatus == 0 {
						capital += va.Capital
						investAll[i].RepayAmountWait += va.Capital + va.Interest
						if va.Period > investAll[i].RepayTerm && investAll[i].RepayTerm <= 0{
							investAll[i].RepayTerm = va.Period
							investAll[i].RepayTime = Substr(va.RepayTime, 0, 10)
						}
					}

				}
			}
			investAll[i].RepayAmountWait = math.Ceil(v.InvestAmount/capital*investAll[i].RepayAmountWait)
			investAll[i].RepayAmount = math.Ceil(v.InvestAmount/capital*investAll[i].RepayAmount)

		}
	}
	appG.Response(http.StatusOK, e.SUCCESS, investAll)
}

func GetMemberInfo(c *gin.Context) {
	var (
		appG = app.Gin{c}
	)
	memberId, _ := c.Get("member_id")
	//判断用户是否登陆
	auth := auth_service.Auth{
		ID: memberId.(int),
	}
	memberInfo, err := auth.GetMemberInfo()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, memberInfo)
}

func Substr(str string, start, length int) string {
	if length == 0 {
		return ""
	}
	rune_str := []rune(str)
	len_str := len(rune_str)

	if start < 0 {
		start = len_str + start
	}
	if start > len_str {
		start = len_str
	}
	end := start + length
	if end > len_str {
		end = len_str
	}
	if length < 0 {
		end = len_str + length
	}
	if start > end {
		start, end = end, start
	}
	return string(rune_str[start:end])
}