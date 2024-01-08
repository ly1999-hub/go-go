package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/model"
	"github.com/ly1999-hub/go-go/internal/response"
	"github.com/ly1999-hub/go-go/internal/util"
	"github.com/ly1999-hub/go-go/pkg/admin/dao"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

func CheckPermission(permission string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var dao = dao.Role{}
			admin := c.Get("admin").(model.Admin)
			if admin.Root {
				return next(c)
			}
			if admin.Role != "" {
				idRole, _ := primitive.ObjectIDFromHex(admin.Role)
				role := dao.FindOne(util.GetRequestContext(c), bson.M{"_id": idRole})
				fmt.Print(funk.Contains(role.Permissions, permission))
				if funk.Contains(role.Permissions, permission) {
					return next(c)
				}
			}
			return response.R401(c, nil, "")
		}
	}
}
