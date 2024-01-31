package middleware

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/ly1999-hub/go-go/internal/constant"
	"github.com/ly1999-hub/go-go/internal/response"
	"github.com/ly1999-hub/go-go/internal/util"
	"github.com/ly1999-hub/go-go/pkg/api/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RequireLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var dao = dao.Customer{}
		idCustomer := getIDCustomerFromContext(c)
		if idCustomer.IsZero() {
			return response.R401(c, nil, "")
		}
		doc := dao.FindByID(util.GetRequestContext(c), idCustomer)
		if !doc.Active {
			return response.R401(c, nil, "")
		}
		c.Set("customer", doc)
		return next(c)
	}
}

func getIDCustomerFromContext(c echo.Context) (id primitive.ObjectID) {
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
