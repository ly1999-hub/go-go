package response

const (
	CommonSuccess       = "common_success"
	CommonNotFound      = "common_not_found"
	CommonExistEmail    = "common_exist_email"
	CommonUnauthorized  = "common_unauthorized"
	CommonBadRequest    = "common_bad_request"
	CommonNotActive     = "common_not_active"
	CommonErrorService  = "common_error_service"
	CommonErrorSendgrid = "common_error_sendgrid"
	CommonErrorPassword = "common_error_password"
)

var common = []Code{
	{
		Key:      CommonSuccess,
		Messenge: "thành công",
		Code:     1,
	},
	{
		Key:      CommonExistEmail,
		Messenge: "đã tồn tại email",
		Code:     50,
	}, {
		Key:      CommonUnauthorized,
		Messenge: "không có quyền thực hiện yêu cầu",
		Code:     401,
	}, {
		Key:      CommonBadRequest,
		Messenge: "Yêu cầu không được thực hiện",
		Code:     400,
	}, {
		Key:      CommonErrorPassword,
		Messenge: "sai mật khẩu",
		Code:     400,
	},
	{
		Key:      CommonNotActive,
		Messenge: "tài khoản bị khóa",
		Code:     401,
	}, {
		Key:      CommonErrorService,
		Messenge: "Lỗi server",
		Code:     500,
	},
	{
		Key:      CommonErrorSendgrid,
		Messenge: "Lỗi khi thực hiện gửi email",
		Code:     500,
	},
}
