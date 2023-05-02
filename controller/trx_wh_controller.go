package controllers

import (
	"go_inven_ctrl/entity"
	"go_inven_ctrl/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrxWhController struct {
	usecase usecase.TrxWhUsecase
}

func (c *TrxWhController) EnrollInsertTrxWh(ctx *gin.Context) {
	var newTrxWh entity.TrxWh

	if err := ctx.BindJSON(&newTrxWh); err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	res := c.usecase.EnrollInsertTrxWh(&newTrxWh)

	ctx.JSON(http.StatusOK, res)
}

func NewTrxWhController(u usecase.TrxWhUsecase) *TrxWhController {
	controller := TrxWhController{
		usecase: u,
	}

	return &controller
}
