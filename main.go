package main

import (
	"log"

	db "github.com/raynaaliyya/ctrl-wh-st/config"
	"github.com/raynaaliyya/ctrl-wh-st/controller"
	"github.com/raynaaliyya/ctrl-wh-st/repository"
	"github.com/raynaaliyya/ctrl-wh-st/router"
	"github.com/raynaaliyya/ctrl-wh-st/usecase"
)

func main() {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could not initialize db connection %s", err)
	}

	// Dependencies Warehouse Team
	warehouseTeamRepo := repository.NewWarehouseTeamRepo(db.GetDB())
	warehouseTeamUsecase := usecase.NewWarehouseTeamUsecase(warehouseTeamRepo)
	warehouseTeamController := controller.NewWarehouseTeamController(warehouseTeamUsecase)

	// Dependencies Product Warehouse
	// productWhRepo := repository.NewProductWhRepo(db.GetDB())
	// productWhUsecase := usecase.NewProductWhUsecase(productWhRepo)
	// productWhController := controller.NewProductWhController(productWhUsecase)

	router.InitRouterEmployee(warehouseTeamController)
	// router.InitRouterProduct(productWhController)
	router.RunServer(":8080")

}
