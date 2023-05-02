package router

import (
	"database/sql"
	"go_inven_ctrl/controllers"
	"go_inven_ctrl/repository"
	"go_inven_ctrl/usecase"

	jwt_controller "github.com/Uchel/auth-final/controller"
	jwt_middleware "github.com/Uchel/auth-final/middleware"
	jwt_repository "github.com/Uchel/auth-final/repository"
	jwt_usecase "github.com/Uchel/auth-final/usecase"

	"github.com/gin-gonic/gin"
)

func InitRouterEmployee(router *gin.Engine, db *sql.DB) {
	// auth jwt login ExampleAdmin
	jwtAdminWhRepo := jwt_repository.NewAdminWhLoginRepo(db)
	jwtAdminWhUsecase := jwt_usecase.NewAdminWhUsecase(jwtAdminWhRepo)
	jwtAdminWhCtrl := jwt_controller.NewAdminLoginController(jwtAdminWhUsecase, 20)

	// Dependencies Warehouse Team
	warehouseTeamRepo := repository.NewWarehouseTeamRepo(db)
	warehouseTeamUsecase := usecase.NewWarehouseTeamUsecase(warehouseTeamRepo)
	warehouseTeamController := controllers.NewWarehouseTeamController(warehouseTeamUsecase)

	// Login session
	router.POST("/auth/login_wh", jwtAdminWhCtrl.LoginAdmin)

	// Route group after authentication
	adminWhRouter := router.Group("/warehouse-team/employees")
	adminWhRouter.Use(jwt_middleware.AuthMiddleware())

	// routes (GET, POST, PUT, DELETE)
	adminWhRouter.POST("/admin_wh", warehouseTeamController.Register)
	adminWhRouter.GET("", warehouseTeamController.FindEmployees)
	adminWhRouter.GET("/:id", warehouseTeamController.FindEmployeeById)
	adminWhRouter.PUT("/update", warehouseTeamController.Edit)
	adminWhRouter.DELETE("/admin_wh/:id", warehouseTeamController.Unreg)

}
