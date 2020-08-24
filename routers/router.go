package routers

import (
	"html_api/middleware/Cors"
	"net/http"

	"github.com/gin-gonic/gin"

	//_ "html_api/docs"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"html_api/middleware/jwt"
	"html_api/pkg/export"
	"html_api/pkg/qrcode"
	"html_api/pkg/upload"
	"html_api/routers/api"
	"html_api/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(Cors.Cors())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.POST("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)

	//借款列表 pagesize
	r.POST("/borrow", v1.GetBorrows)
	//注册
	r.POST("/register", v1.Register)
	//登陆

	apiv1 := r.Group("/api/v1")
	//获取借款详情
	apiv1.POST("/getborrow", v1.GetBorrow)
	//生成还款脚本
	apiv1.POST("/repaycreate", v1.Repay)
	apiv1.POST("/article", v1.GetArticles)
	apiv1.POST("/getarticle", v1.GetArticle)
	apiv1.POST("/getborrowsummary", v1.GetBorrowSummary)
	apiv1.Use(jwt.JWT())
	{

		//借款申请
		apiv1.POST("/borrowapply", v1.AddBorrowApply)
		//充值申请
		apiv1.POST("/recharge", v1.AddRecharge)
		//投资申请
		apiv1.POST("/invest", v1.Invest)
		//获取用户信息
		apiv1.POST("/member", v1.Member)
		//获取用户投资列表
		apiv1.POST("/memberinvest", v1.MemberInvest)
		//获取用户详细信息
		apiv1.POST("/memberinfo", v1.GetMemberInfo)
	}

	return r
}
