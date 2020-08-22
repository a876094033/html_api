package e

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_ADD_BORROWAPPLY:          "提交借款申请失败",
	ERROR_NOT_EXIST_BORROWAPPLY:    "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
	ERROR_AUTH_REGISTER_EMAIL:      "邮箱已经被注册",
	ERROR_AUTH_REGISTER_FAILED:     "注册失败",

	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "图片验证失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "图片上传失败",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "图片保存失败",

	ERROR_COUNT_BORROW_FAIL:      "获取借款总数失败",
	ERROR_GET_BORROW_FAIL:        "获取借款列表失败",
	ERROR_GET_BORROW_DETAIL_FAIL: "获取借款详情失败",

	ERROR_ADD_RECHARGE:               "充值失败",
	ERROR_MEMBER_AMOUNT_USED_LESS:    "用户余额不足",
	ERROR_MEMBER_BORROW_AMOUNT_LIMIT: "超过标的剩余可投资额度",
	ERROR_MEMBER_INVEST_FAILED:       "投资失败",
	ERROR_MEMBER_INVEST_LIST:         "获取投资列表失败",

	//脚本
	ERROR_BORROW_REPAY_CREATE_NONE: "获取满标的借款失败",

	ERROR_ARTICLE_TOTAL_FAILED: "获取文章总数失败",
	ERROR_ARTICLE_LIST_FAILED:  "获取文章列表失败",
	ERROR_ARTICLE_INFO_FAILED:  "获取文章信息失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
