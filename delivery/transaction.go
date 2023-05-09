package delivery

import (
	"self-payroll/helper"
	"self-payroll/model"
	"self-payroll/utils"
	"strconv"

	"github.com/labstack/echo/v4"
)

type transactionDelivery struct {
	transactionUsecase model.TransactionUsecase
}

type TransactionDelivery interface {
	Mount(group *echo.Group)
}

func NewTransactionDelivery(transactionUsecase model.TransactionUsecase) TransactionDelivery {
	return &transactionDelivery{transactionUsecase: transactionUsecase}
}

func (p *transactionDelivery) Mount(group *echo.Group) {
	group.GET("/all", p.FetchTransactionHandler, utils.CheckIsAdmin)
	group.GET("/my", p.FetchTransactionByUserID)
}

func (p *transactionDelivery) FetchTransactionHandler(c echo.Context) error {
	ctx := c.Request().Context()

	limit := c.QueryParam("limit")
	offset := c.QueryParam("skip")

	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)

	transactions, i, err := p.transactionUsecase.Fetch(ctx, limitInt, offsetInt)
	if err != nil {
		return helper.ResponseErrorJson(c, i, err)
	}

	return helper.ResponseSuccessJson(c, "success", transactions)
}

func (p *transactionDelivery) FetchTransactionByUserID(c echo.Context) error {
	ctx := c.Request().Context()
	userID := c.Get("userID").(float64)

	limit := c.QueryParam("limit")
	offset := c.QueryParam("skip")

	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)

	transactions, code, err := p.transactionUsecase.FetchByUserID(ctx, int(userID), limitInt, offsetInt)
	if err != nil {
		return helper.ResponseErrorJson(c, code, err)
	}
	return helper.ResponseSuccessJson(c, "success", transactions)
}
