package v1

import (
	"github.com/gin-gonic/gin"
	"html_api/pkg/app"
	"html_api/pkg/e"
	"html_api/pkg/util"
	"html_api/service/register_service"
	"net/http"
)

type RegisterForm struct {
	Email      string `form:"email" json:"email" valid:"Required; Email; MaxSize(100)"`
	ReEmail    string `form:"re_email" json:"re_email" valid:"Required; Email; MaxSize(100)"`
	Password   string `form:"password" json:"password" valid:"Required; MinSize(8); MaxSize(20)"`
	RePassword string `form:"re_password" json:"re_password" valid:"Required; MinSize(8); MaxSize(20)"`
}

func Register(c *gin.Context) {
	var (
		appG = app.Gin{c}
		form RegisterForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, form)
		return
	}

	register := register_service.Register{
		Email:    form.Email,
		Password: form.Password,
	}
	//检测邮箱是否被注册
	member_id := register.CheckEmail()
	if member_id > 0 {
		appG.Response(http.StatusOK, e.ERROR_AUTH_REGISTER_EMAIL, nil)
		return
	}
	member_id = register.Add()
	if  member_id <= 0 {
		appG.Response(http.StatusOK, e.ERROR_ADD_BORROWAPPLY, nil)
		return
	}
	//注册成功 生成token
	token, err := util.GenerateToken(form.Email, form.Password, member_id)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}
