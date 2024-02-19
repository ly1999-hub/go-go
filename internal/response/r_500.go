package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func R500(c echo.Context, data interface{}, key string) error {
	if key == "" {
		key = CommonErrorService
	}
	localData := GetByKey(key)
	if localData.Code == -1 {
		localData.Messenge = key
	}
	return sendResponse(c, http.StatusInternalServerError, false, data, localData.Messenge, localData.Code)
}
