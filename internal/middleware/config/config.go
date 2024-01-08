package config

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ly1999-hub/go-go/internal/constant"
)

// CORSConfig ...
func CORSConfig() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions, http.MethodPut, http.MethodPatch, http.MethodDelete},
		AllowHeaders: []string{
			echo.HeaderOrigin, echo.HeaderContentLength, echo.HeaderContentType,
			echo.HeaderAuthorization, constant.HeaderAcceptLanguage, constant.HeaderAppVersion,
			constant.HeaderDeviceID, constant.HeaderUserAgent, constant.HeaderModel,
			constant.HeaderManufacturer, constant.HeaderFCMToken,
		},
		AllowCredentials: false,
		MaxAge:           600,
	})
}

// RateLimiterConfig ..
func RateLimiterConfig() echo.MiddlewareFunc {
	return middleware.RateLimiterWithConfig(middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 10, Burst: 30, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	})
}

// LoggerWithConfig ...
func LoggerWithConfig() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
	})
}
