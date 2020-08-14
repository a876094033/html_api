package v1

import (
	//"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"html_api/pkg/app"
	"html_api/pkg/e"
	"html_api/service/borrow_apply"
	"net/http"
)

type AddBorrowApplyForm struct {
	Name          string  `form:"name" valid:"Required"`
	Phone         string  `form:"phone" valid:"Required"`
	Sex           int     `form:"sex" valid:"Range(1,2)"`
	Birth         string  `form:"birth" valid:"Required"`
	Email         string  `form:"email" valid:"Required"`
	ApplyType     int     `form:"apply_type" valid:"Range(1,2)"`
	PropertyType  int     `form:"property_type" valid:"Required;Min(1);Max(10)"`
	SeniorAmount  float64 `form:"senior_amount" valid:"Required"`
	ApplyAmount   float64 `form:"apply_amount" valid:"Required"`
	Period        int     `form:"period" valid:"Required;Min(1)"`
	CaseNumber    string  `form:"case_number" valid:"Required"`
	Postcode      string  `form:"postcode" valid:"Required"`
	Address       string  `form:"address" valid:"Required"`
	AddressDetail string  `form:"address_detail"`
}

//申请借款
func AddBorrowApply(c *gin.Context) {
	var (
		appG = app.Gin{c}
		form AddBorrowApplyForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)

	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, form)
		return
	}

	borrowapply := borrow_apply.BorrowApply{
		Name:          form.Name,
		Phone:         form.Phone,
		Birth:         form.Birth,
		Sex:           form.Sex,
		Email:         form.Email,
		ApplyType:     form.ApplyType,
		PropertyType:  form.PropertyType,
		SeniorAmount:  form.SeniorAmount,
		ApplyAmount:   form.ApplyAmount,
		Period:        form.Period,
		CaseNumber:    form.CaseNumber,
		Postcode:      form.Postcode,
		Address:       form.Address,
		AddressDetail: form.AddressDetail,
	}

	if err := borrowapply.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_BORROWAPPLY, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)

}
