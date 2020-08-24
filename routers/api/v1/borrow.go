package v1

import (
	"github.com/Unknwon/com"
	_ "github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"html_api/models"
	"html_api/pkg/app"
	"html_api/pkg/e"
	"html_api/pkg/setting"
	"html_api/pkg/util"
	"html_api/service/borrow_service"
	"net/http"
)

func GetBorrows(c *gin.Context) {
	appG := app.Gin{c}
	var pageSize int
	if pageSize = com.StrTo(c.Query("page_size")).MustInt(); pageSize == 0 {
		pageSize = setting.AppSetting.PageSize
	}
	borrowStatus := com.StrTo(c.Query("borrow_status")).MustInt()

	//valid := validation.Validation{}
	borrowService := borrow_service.Borrow{
		PageNum:  util.GetPage(c),
		PageSize: pageSize,
	}
	if borrowStatus > 1 {
		borrowService.BorrowStatus = borrowStatus
	}
	total, err := borrowService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_BORROW_FAIL, nil)
		return
	}

	borrows, err := borrowService.GetBorrows()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_BORROW_FAIL, nil)
		return
	}

	//var borrowAll []models.BorrowAll
	borrowAll := make([]models.BorrowAll, len(borrows))
	for i, v := range borrows {
		borrowAll[i].ID = v.ID
		borrowAll[i].Term = v.Term
		borrowAll[i].BorrowStatus = v.BorrowStatus
		borrowAll[i].BorrowName = v.BorrowName
		borrowAll[i].RepayType = v.RepayType
		borrowAll[i].Amount = v.Amount
		borrowAll[i].AmountLimit = v.AmountLimit
		borrowAll[i].Purpose = v.Purpose
		borrowAll[i].RepayName = v.RepayName
		borrowAll[i].Diya = v.Diya
		borrowAll[i].TermType = v.TermType
		borrowAll[i].InterestRate = v.InterestRate
		borrowAll[i].InvestCount = models.GetInvestCount(v.ID)
		borrowAll[i].InvestSum = v.Amount - v.AmountLimit
		borrowAll[i].BorrowImg = v.BorrowImg
	}
	data := make(map[string]interface{})
	data["total"] = total
	data["list"] = borrowAll
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

type BorrowIdForm struct {
	BorrowID int `form:"borrow_id" json:"borrow_id"`
}

func GetBorrow(c *gin.Context) {
	var (
		appG     = app.Gin{c}
		borrowId BorrowIdForm
	)

	httpCode, errCode := app.BindAndValid(c, &borrowId)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, borrowId.BorrowID)
		return
	}
	if borrowId.BorrowID <= 0 {
		appG.Response(http.StatusInternalServerError, e.INVALID_PARAMS, borrowId.BorrowID)
		return
	}
	borrow_service := borrow_service.Borrow{
		ID: borrowId.BorrowID,
	}
	borrow, err := borrow_service.GetBorrow()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_BORROW_DETAIL_FAIL, borrowId.BorrowID)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, borrow)
}
