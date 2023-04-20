package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raynaaliyya/ctrl-wh-st/models"
	"github.com/raynaaliyya/ctrl-wh-st/usecase"
)

type ProductWhController struct {
	usecase usecase.ProductWhUsecase
}

func NewProductWhController(u usecase.ProductWhUsecase) *ProductWhController {
	controller := ProductWhController{
		usecase: u,
	}

	return &controller
}

func (c *ProductWhController) FindProducts(ctx *gin.Context) {
	res := c.usecase.FindProducts()

	ctx.JSON(http.StatusOK, res)
}

func (c *ProductWhController) FindProductById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid product ID")
		return
	}

	res := c.usecase.FindProductById(id)
	ctx.JSON(http.StatusOK, res)
}

func (c *ProductWhController) FindProductByName(ctx *gin.Context) {
	name := ctx.Param("name")

	res, err := c.usecase.FindProductByName(name)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid product Name")
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (c *ProductWhController) Input(ctx *gin.Context) {
	var product models.ProductWh

	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data. Please check the request body and try again."})
		return
	}

	res, err := c.usecase.Input(&product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (c *ProductWhController) Edit(ctx *gin.Context) {
	var product models.ProductWh

	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}

	res := c.usecase.Edit(&product)
	ctx.JSON(http.StatusOK, res)
}

func (c *ProductWhController) Output(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid user ID")
		return
	}

	res := c.usecase.Output(id)
	ctx.JSON(http.StatusOK, res)
}
