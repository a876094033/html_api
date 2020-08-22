package v1

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"html_api/pkg/app"
	"html_api/pkg/e"
	"html_api/pkg/setting"
	"html_api/pkg/util"
	"html_api/service/article_service"
	"net/http"
)

func GetArticles(c *gin.Context) {
	appG := app.Gin{c}
	var pageSize int
	if pageSize = com.StrTo(c.Query("page_size")).MustInt(); pageSize == 0 {
		pageSize = setting.AppSetting.PageSize
	}

	article := article_service.Article{
		PageNum:  util.GetPage(c),
		PageSize: pageSize,
	}

	count, err := article.Count()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_ARTICLE_TOTAL_FAILED, nil)
		return
	}
	articleAll, err := article.GetArticles()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_ARTICLE_LIST_FAILED, nil)
		return
	}

	data := make(map[string]interface{})
	data["total"] = count
	data["list"] = articleAll
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

type article struct {
	ID int `form:"id"`
}

func GetArticle(c *gin.Context) {
	appG := app.Gin{c}
	var form article
	_, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
	}
	article := article_service.Article{
		ID: form.ID,
	}
	articleinfo, err := article.GetArticle()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_ARTICLE_INFO_FAILED, nil)
	}
	appG.Response(http.StatusOK, e.SUCCESS, articleinfo)
}
