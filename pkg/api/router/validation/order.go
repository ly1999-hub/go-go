package validation

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"github.com/ly1999-hub/go-go/internal/util/log"
)

type Order struct{}

func (v Order) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload model.OrderRequest
		payload.UserOrder = c.Get("user_login").(model.User).ID.Hex()
		if err := c.Bind(&payload); err != nil {
			log.Error("Error bind Order", log.LogData{
				"payload": payload,
				"error":   err.Error(),
			})
			return response.R400(c, nil, "")
		}
		c.Set("order_request", payload)
		return next(c)
	}
}
