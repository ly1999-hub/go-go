package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"github.com/ly1999-hub/go-go/internal/util"
	daoAdmin "github.com/ly1999-hub/go-go/pkg/admin/dao"
	"github.com/thoas/go-funk"
	"gopkg.in/mgo.v2/bson"
)

func CheckPermission(permission string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var (
				daoRole = daoAdmin.Role{}
			)
			customer := c.Get("customer").(model.Customer)
			if !customer.Active {
				return response.R401(c, nil, "")
			}
			if customer.Role != "" {
				role := daoRole.FindOne(util.GetRequestContext(c), bson.M{"code": customer.Role})
				if funk.Contains(role.Permissions, permission) {
					return next(c)
				}
			}
			return response.R401(c, nil, "")
		}
	}
}
