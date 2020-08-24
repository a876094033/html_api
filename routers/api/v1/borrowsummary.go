package v1

import (
	"github.com/gin-gonic/gin"
	"html_api/pkg/app"
	"html_api/pkg/e"
	"html_api/service/borrow_summary"
	"net/http"
)

type borrowSummaryForm struct {
	BorrowID int `form:"borrow_id"`
}

func GetBorrowSummary(c *gin.Context) {
	var (
		appG = app.Gin{c}
		form borrowSummaryForm
	)
	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, form)
		return
	}

	summary := borrow_summary.BorrowSummary{
		BorrowID: form.BorrowID,
	}
	borrowSummary, err := summary.GetBorrowSummary()
	if err != nil {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
	}

	appG.Response(http.StatusOK, e.SUCCESS, borrowSummary)
}
