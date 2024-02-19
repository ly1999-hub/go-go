package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func R404(c echo.Context, data interface{}, key string) error {
	if key == "" {
		key = CommonNotFound
	}
	localData := GetByKey(key)
	if localData.Code == -1 {
		localData.Messenge = key
	}
	return sendResponse(c, http.StatusNotFound, false, data, localData.Messenge, localData.Code)
}
