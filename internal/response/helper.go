package response

import "github.com/labstack/echo/v4"

func sendResponse(c echo.Context, httpCode int, success bool, data interface{}, message string, code int) error {
	if data == nil {
		data = echo.Map{}
	}
	return c.JSON(httpCode, echo.Map{
		"success": success,
		"data":    data,
		"message": message,
		"code":    code,
	})
}
