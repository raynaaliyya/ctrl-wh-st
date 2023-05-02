package router

import (
	"database/sql"
	"go_inven_ctrl/controllers"
	"go_inven_ctrl/repository"
	"go_inven_ctrl/usecase"

	"github.com/gin-gonic/gin"
)

func InitRouterTrxReportWh(router *gin.Engine, db *sql.DB) {
	// Dependencies Warehouse Transaction
	trxWhRepo := repository.NewTrxWhRepo(db)
	trxWhUsecase := usecase.NewTrxWhUsecase(trxWhRepo)
	trxWhController := controllers.NewTrxWhController(trxWhUsecase)

	// Dependencies Warehouse Report
	reportTrxWhRepo := repository.NewReportTrxWhRepo(db)
	reportTrxWhUsecase := usecase.NewReportTrxWhUsecase(reportTrxWhRepo)
	reportTrxWhController := controllers.NewReportTrxWhController(reportTrxWhUsecase)

	// Warehouse Routes
	trxInWhRouter := router.Group("/trxwarehouse")
	trxInWhRouter.POST("/transaction", trxWhController.EnrollInsertTrxWh)
	trxInWhRouter.GET("", reportTrxWhController.FindAllReportTrxWh)
	trxInWhRouter.GET("/reportwhid", reportTrxWhController.FindByIdReportTrxWh)
	trxInWhRouter.GET("/reportwhdate", reportTrxWhController.FindByDateReportTrxWh)
}
