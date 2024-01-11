package middleware

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/constant"
	"github.com/ly1999-hub/go-go/internal/response"
	"github.com/ly1999-hub/go-go/internal/util"
	"github.com/ly1999-hub/go-go/internal/util/log"
	"github.com/ly1999-hub/go-go/pkg/admin/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RequireLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			dao = dao.Admin{}
		)
		adminId := getIDAdminFromContext(c)
		if adminId.IsZero() {
			log.Error("Not Admin", log.LogData{})
			return response.R401(c, nil, "")
		}
		admin := dao.FindByID(util.GetRequestContext(c), adminId)
		if admin.ID.IsZero() || !admin.Active {

			return response.R401(c, nil, "")
		}
		c.Set("admin", admin)
		return next(c)
	}
}

func getIDAdminFromContext(c echo.Context) (id primitive.ObjectID) {
	authHeader := c.Request().Header.Get(constant.HeaderAuthorization)
	if authHeader == "" {
		return
	}
	tokenString := authHeader[7:]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		fmt.Println("Failed to parse token:", err)
		return
	}

	data, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return
	}

	if ok {
		idString := data["_id"].(string)
		id, _ = primitive.ObjectIDFromHex(idString)
	}

	return
}
