package controllers

import (
	"fmt"
	"go_inven_ctrl/entity"
	"go_inven_ctrl/usecase"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
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
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "wh" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	res := c.usecase.FindEmployees()

	ctx.JSON(http.StatusOK, gin.H{
		"data":       res,
		"login with": email,
	})
}

func (c *WarehouseTeamController) FindEmployeeById(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	role := claims["role"].(string)
	if role != "wh" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	id := ctx.Param("id")

	photo := c.usecase.FindEmployeeById(id)
	filepath := fmt.Sprintf("D:/Documents/coding/final-project-enigma/goventory-control/goventory-control/images/%s", photo)

	file, err := os.Open(filepath)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	fileBytes := make([]byte, fileInfo.Size())
	_, err = file.Read(fileBytes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	ctx.Data(http.StatusOK, "image/jpeg", fileBytes)
}

func (c *WarehouseTeamController) Register(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "wh" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	var newEmployee entity.WarehouseTeam

	// request body
	if err := ctx.ShouldBind(&newEmployee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	file, err := ctx.FormFile("photo")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input Photo"})
		return
	}

	filename := file.Filename
	newEmployee.Photo = filename

	if errFile := ctx.SaveUploadedFile(file, fmt.Sprintf("./images/%s", filename)); errFile != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	res := c.usecase.Register(&newEmployee)

	ctx.JSON(http.StatusCreated, gin.H{
		"data":       res,
		"login with": email,
	})
}

func (c *WarehouseTeamController) Edit(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "wh" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	var employee entity.WarehouseTeam

	if err := ctx.BindJSON(&employee); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data. Please check the request body and try again."})
		return
	}

	res := c.usecase.Edit(&employee)
	ctx.JSON(http.StatusOK, gin.H{
		"data":       res,
		"login with": email,
	})
}

func (c *WarehouseTeamController) Unreg(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "wh" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	id := ctx.Param("id")

	photo := c.usecase.FindEmployeeById(id)
	filePath := fmt.Sprintf("D:/Documents/coding/final-project-enigma/goventory-control/goventory-control/images/%s", photo)

	err := os.Remove(filePath)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	res := c.usecase.Unreg(id)
	ctx.JSON(http.StatusOK, gin.H{
		"data":       res,
		"login with": email,
	})
}
