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
)

var common = []Code{
	{
		Key:      CommonSuccess,
		Messenge: "thành công",
		Code:     1,
	},
	{
		Key:      CommonExistEmail,
		Messenge: "da ton tai email",
		Code:     50,
	}, {
		Key:      CommonUnauthorized,
		Messenge: "không có quyền thực hiện yêu cầu",
		Code:     401,
	}, {
		Key:      CommonBadRequest,
		Messenge: "Yeu cau khong duoc thuc hien",
		Code:     400,
	},
}
