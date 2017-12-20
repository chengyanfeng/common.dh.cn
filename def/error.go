package def

const (
	GENERAL_ERR   int = 400
	NO_SHARE_ERR  int = 401
	PWD_REQUIRED  int = 402
	IP_LIMIT_ERR  int = 403
	NOT_LOGIN     int = 404
	DATA_LINE_ERR int = 405

	ERROR_OK      = 0     // 请求成功
	ERROR_FALSE   = -1    // 系统繁忙,服务器暂不可用，建议稍候重试
	ERROR_WAITING = 1     // 正在处理，请稍等
	ERROR_PARAM   = 40001 // 参数值错误或非法，请检查参数值是否有效
)
