package utils

import (
	"errors"
	"net/http"
	"self-payroll/helper"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("token")
		if token == "" {
			return helper.ResponseErrorJson(c, http.StatusUnauthorized, errors.New("token doko?"))
		}

		claims, err := DecodeJWT(token)
		if err != nil {
			return helper.ResponseErrorJson(c, http.StatusUnauthorized, errors.New("invalid token"))
		}

		c.Set("positionID", claims["positionID"])
		c.Set("userID", claims["id"])
		c.Set("isAdmin", claims["isAdmin"])
		return next(c)
	}
}

func CheckIsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !c.Get("isAdmin").(bool) {
			return helper.ResponseErrorJson(c, http.StatusUnauthorized, errors.New("unauthorized"))
		}
		return next(c)
	}
}
