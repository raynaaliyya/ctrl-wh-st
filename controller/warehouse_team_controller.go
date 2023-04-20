package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/raynaaliyya/ctrl-wh-st/models"
	"github.com/raynaaliyya/ctrl-wh-st/usecase"
)

type WarehouseTeamController struct {
	usecase usecase.WarehouseTeamUsecase
}

func NewWarehouseTeamController(u usecase.WarehouseTeamUsecase) *WarehouseTeamController {
	return &WarehouseTeamController{
		usecase: u,
	}
}

func (c *WarehouseTeamController) FindEmployees(ctx *gin.Context) {
	res := c.usecase.FindEmployees()

	ctx.JSON(http.StatusOK, res)
}

func (c *WarehouseTeamController) FindEmployeeById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid employee ID")
		return // return kosong untuk memberhentikan fungsi
	}

	res := c.usecase.FindEmployeeById(id)
	ctx.JSON(http.StatusOK, res)
}

func (c *WarehouseTeamController) Login(ctx *gin.Context) {
	var employee models.EmployeeReq
	if err := ctx.ShouldBindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := c.usecase.Login(&employee)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.SetCookie("jwt", u.AccessToken, 60*60*24, "/", "localhost", false, true)
	ctx.JSON(http.StatusOK, u)
}

func (c *WarehouseTeamController) Register(ctx *gin.Context) {
	var newEmployee models.WarehouseTeam

	// request body
	if err := ctx.ShouldBindJSON(&newEmployee); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}

	res, err := c.usecase.Register(&newEmployee)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

func (c *WarehouseTeamController) Edit(ctx *gin.Context) {
	var employee models.WarehouseTeam

	if err := ctx.BindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data. Please check the request body and try again."})
		return
	}

	res := c.usecase.Edit(&employee)
	ctx.JSON(http.StatusOK, res)
}

func (c *WarehouseTeamController) Unreg(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid employee ID")
		return
	}

	res := c.usecase.Unreg(id)
	ctx.JSON(http.StatusOK, res)
}
