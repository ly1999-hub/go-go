package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func R403(c echo.Context, data interface{}, key string) error {
	if key == "" {
		key = CommonForbidden
	}
	localData := GetByKey(key)
	if localData.Code == -1 {
		localData.Messenge = key
	}

	return sendResponse(c, http.StatusForbidden, false, data, localData.Messenge, localData.Code)
}
