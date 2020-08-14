package v1

import (
	"github.com/gin-gonic/gin"
	"html_api/pkg/app"
	"html_api/pkg/e"
	"html_api/service/auth_service"
	"html_api/service/invest_service"
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
	appG.Response(http.StatusOK, e.SUCCESS, investList)
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
