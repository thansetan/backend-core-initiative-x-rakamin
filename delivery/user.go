package delivery

import (
	"fmt"
	"net/http"
	"reflect"
	"self-payroll/helper"
	"self-payroll/model"
	"self-payroll/request"
	"self-payroll/utils"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

type userDelivery struct {
	userUsecase model.UserUsecase
}

type UserDelivery interface {
	Mount(group *echo.Group)
}

func NewUserDelivery(userUsecase model.UserUsecase) UserDelivery {
	return &userDelivery{userUsecase: userUsecase}
}

func (p *userDelivery) Mount(group *echo.Group) {
	group.GET("", p.FetchUserHandler, utils.AuthMiddleware, utils.CheckIsAdmin)
	group.POST("", p.StoreUserHandler)
	group.GET("/my", p.DetailUserHandler, utils.AuthMiddleware)
	group.DELETE("/my", p.DeleteUserHandler, utils.AuthMiddleware)
	group.PATCH("/my", p.EditUserHandler, utils.AuthMiddleware)
	group.POST("/withdraw", p.WithdrawHandler, utils.AuthMiddleware)
	group.POST("/admin/register", p.AdminRegisterHandler)
	group.POST("/login", p.UserLoginHandler)
}

func (p *userDelivery) FetchUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	limit := c.QueryParam("limit")
	offset := c.QueryParam("skip")

	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)

	userList, err := p.userUsecase.FetchUser(ctx, limitInt, offsetInt)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return helper.ResponseSuccessJson(c, "success", userList)

}

func (p *userDelivery) StoreUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())

	}

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(c, "Error validation", errVal)
	}

	user, err := p.userUsecase.StoreUser(ctx, &req)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}

	return helper.ResponseSuccessJson(c, "success", user)
}

func (p *userDelivery) DetailUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Get("userID").(float64)

	user, err := p.userUsecase.GetByID(ctx, int(id))
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}

	return helper.ResponseSuccessJson(c, "", user)

}

func (p *userDelivery) DeleteUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Get("userID").(float64)
	fmt.Println("id nya : ", id, "tipenya: ", reflect.TypeOf(id))

	err := p.userUsecase.DestroyUser(ctx, int(id))
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
	}

	return helper.ResponseSuccessJson(c, "", "")

}

func (p *userDelivery) EditUserHandler(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.UpdateRequest

	if err := c.Bind(&req); err != nil {
		return helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())

	}

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(c, "Error validation", errVal)
	}

	id := c.Get("userID").(float64)

	user, err := p.userUsecase.EditUser(ctx, int(id), &req)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
	}

	return helper.ResponseSuccessJson(c, "Success edit", user)
}

func (p *userDelivery) WithdrawHandler(c echo.Context) error {
	userID := c.Get("userID").(float64)
	ctx := c.Request().Context()
	var req request.WithdrawRequest
	if err := c.Bind(&req); err != nil {
		return helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())
	}
	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(c, "Error validation", errVal)
	}
	fmt.Println("userIDnya ===>>>", int(userID))
	err := p.userUsecase.WithdrawSalary(ctx, int(userID), &req)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
	}

	return helper.ResponseSuccessJson(c, "Success withdraw salary", "")

}

func (p *userDelivery) AdminRegisterHandler(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.AdminRequest
	if err := c.Bind(&req); err != nil {
		return helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())
	}

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(c, "Error validation", errVal)
	}
	admin, err := p.userUsecase.AdminRegister(ctx, &req)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnauthorized, err)
	}
	return helper.ResponseSuccessJson(c, "success", admin)
}

func (p *userDelivery) UserLoginHandler(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.LoginRequest
	if err := c.Bind(&req); err != nil {
		return helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())
	}

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(c, "Error validation", errVal)
	}
	user, err := p.userUsecase.Login(ctx, &req)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnauthorized, err)
	}

	return helper.ResponseSuccessJson(c, "success", user)
}
