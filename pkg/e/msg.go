package e

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "페이지에 오류가 생겼습니다.",
	ERROR_ADD_BORROWAPPLY:          "대출신청이 실패했습니다.",
	ERROR_NOT_EXIST_BORROWAPPLY:    "문장이 존재하지 않습니다. ",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token방문권한 검증이 실패했습니다.",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token유효기간이 초과됬습니다.",
	ERROR_AUTH_TOKEN:               "Token생성실패.",
	ERROR_AUTH:                     "Token오류가 생겼습니다.",
	ERROR_AUTH_REGISTER_EMAIL:      "이미 등록된 메일입니다.",
	ERROR_AUTH_REGISTER_FAILED:     "등록실패.",

	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "사진검증이 실패했습니다.",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "사진업로드가 실패했습니다. ",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "사진저장이 실패했습니다.",

	ERROR_COUNT_BORROW_FAIL:      "대출총금액 획득실패.",
	ERROR_GET_BORROW_FAIL:        "대출리스트 획득실패.",
	ERROR_GET_BORROW_DETAIL_FAIL: "대출상세정보 획득실패.",

	ERROR_ADD_RECHARGE:               "충전이 실패했습니다.",
	ERROR_MEMBER_AMOUNT_USED_LESS:    "잔액이 부족합니다.",
	ERROR_MEMBER_BORROW_AMOUNT_LIMIT: "상품의 잔여투자한도가 초과되었습니다.",
	ERROR_MEMBER_INVEST_FAILED:       "투자가 실패했습니다.",
	ERROR_MEMBER_INVEST_LIST:         "투자리스트 획득실패.",

	//脚本
	ERROR_BORROW_REPAY_CREATE_NONE: "마감된 상품의 대출이 실패했습니다.",

	ERROR_ARTICLE_TOTAL_FAILED: "문장총수량 획득실패.",
	ERROR_ARTICLE_LIST_FAILED:  "문장리스트 획득실패.",
	ERROR_ARTICLE_INFO_FAILED:  "문장정보 획득실패.",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
