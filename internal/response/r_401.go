package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func R401(c echo.Context, data interface{}, key string) error {
	if key == "" {
		key = CommonUnauthorized
	}

	localeData := GetByKey(key)
	if localeData.Code == -1 {
		localeData.Messenge = key
	}
	return sendResponse(c, http.StatusUnauthorized, false, data, localeData.Messenge, localeData.Code)
}
