package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func R201(c echo.Context, data interface{}, key string) error {
	if key == "" {
		key = CommonCreated
	}
	localData := GetByKey(key)
	if localData.Code == -1 {
		localData.Messenge = key
	}
	return sendResponse(c, http.StatusCreated, true, data, localData.Messenge, localData.Code)
}
