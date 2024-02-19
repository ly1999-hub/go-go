package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"github.com/ly1999-hub/go-go/internal/util"
	"github.com/ly1999-hub/go-go/pkg/admin/dao"
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
				return response.R401(c, nil, response.CommonNotActive)
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

func CheckUserPermission(permission string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var (
				daoRole = dao.Role{}
			)
			user := c.Get("user_login").(model.User)
			if !user.Active {
				return response.R401(c, nil, response.CommonNotActive)
			}
			if user.Role != "" {
				role := daoRole.FindOne(util.GetRequestContext(c), bson.M{"code": user.Role})
				if funk.Contains(role.Permissions, permission) {
					return next(c)
				}
			}
			return response.R401(c, nil, response.CommonUnauthorized)
		}
	}
}
