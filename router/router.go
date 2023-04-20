package router

import (
	"github.com/gin-gonic/gin"
	"github.com/raynaaliyya/ctrl-wh-st/controller"
)

var r *gin.Engine

func InitRouterEmployee(employeeController *controller.WarehouseTeamController) {
	// gin router
	r = gin.Default()

	r.POST("/register", employeeController.Register)
	r.POST("/auth/login", employeeController.Login)

	// group users
	userRouter := r.Group("/warehouse-team/employees")
	userRouter.Use(authMiddleware())

	// other routes
	userRouter.GET("/:id/profile", profile)

	// routes (GET, POST, PUT, DELETE)
	userRouter.GET("", employeeController.FindEmployees)
	userRouter.GET("/:id", employeeController.FindEmployeeById)
	userRouter.PUT("", employeeController.Edit)
	userRouter.DELETE("/:id", employeeController.Unreg)
}

func InitRouterProduct(productController *controller.ProductWhController) {
	// gin router
	r = gin.Default()

	r.POST("/input", productController.Input)

	// group users
	userRouter := r.Group("/warehouse/products")

	// routes (GET, POST, PUT, DELETE)
	userRouter.GET("", productController.FindProducts)
	userRouter.GET("/:id", productController.FindProductById)
	userRouter.GET("/:name", productController.FindProductByName)
	userRouter.PUT("", productController.Edit)
	userRouter.DELETE("/:id", productController.Output)
}

func RunServer(addr string) error {
	return r.Run(addr)
}
