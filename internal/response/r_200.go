package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func R200(c echo.Context, data interface{}, key string) error {
	if key == "" {
		key = CommonSuccess
	}
	localeData := GetByKey(key)

	if localeData.Code == -1 {
		localeData.Messenge = key
	}
	return sendResponse(c, http.StatusOK, true, data, localeData.Messenge, localeData.Code)
}
