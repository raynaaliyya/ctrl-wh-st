package repository

import (
	"database/sql"
	"go_inven_ctrl/entity"
	"log"
)

type ReportTrxWhRepo interface {
	GetAllReportTrxWh() any
	GetByIdReportTrxWh(id string) any
	GetByDateReportTrxWh(date string) any
}

type reportTrxWhRepo struct {
	db *sql.DB
}

func (r *reportTrxWhRepo) GetAllReportTrxWh() any {
	var reportsTrxWh []entity.ReportWh

	query := "SELECT product_wh_id, stock, product_name, act, last_stock, created_at FROM report_trx_wh ORDER BY act ASC, created_at ASC"
	rows, err := r.db.Query(query)

	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var reportTrxWh entity.ReportWh

		if err := rows.Scan(&reportTrxWh.ProductWhId, &reportTrxWh.Stock, &reportTrxWh.ProductName, &reportTrxWh.Act, &reportTrxWh.LastStock, &reportTrxWh.CreatedAt); err != nil {
			log.Println(err)
		}

		reportsTrxWh = append(reportsTrxWh, reportTrxWh)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	if len(reportsTrxWh) == 0 {
		return "no data"
	}

	return reportsTrxWh

}

func (r *reportTrxWhRepo) GetByIdReportTrxWh(id string) any {
	var reportsTrxWh []entity.ReportWh

	query := "SELECT product_wh_id, stock, product_name, act, last_stock, created_at FROM report_trx_wh WHERE product_wh_id = $1 ORDER BY act ASC, created_at ASC"

	rows, err := r.db.Query(query, id)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var reportTrxWh entity.ReportWh

		if err := rows.Scan(&reportTrxWh.ProductWhId, &reportTrxWh.Stock, &reportTrxWh.ProductName, &reportTrxWh.Act, &reportTrxWh.LastStock, &reportTrxWh.CreatedAt); err != nil {
			log.Println(err)
		}

		reportsTrxWh = append(reportsTrxWh, reportTrxWh)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	if len(reportsTrxWh) == 0 {
		return "no data or product id not found"
	}

	return reportsTrxWh

}

func (r *reportTrxWhRepo) GetByDateReportTrxWh(date string) any {
	var reportsTrxWh []entity.ReportWh

	query := "SELECT product_wh_id, stock, product_name, act, last_stock, created_at FROM report_trx_wh WHERE created_at = $1 ORDER BY act ASC, created_at ASC"

	rows, err := r.db.Query(query, date)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var reportTrxWh entity.ReportWh

		if err := rows.Scan(&reportTrxWh.ProductWhId, &reportTrxWh.Stock, &reportTrxWh.ProductName, &reportTrxWh.Act, &reportTrxWh.LastStock, &reportTrxWh.CreatedAt); err != nil {
			log.Println(err)
		}

		reportsTrxWh = append(reportsTrxWh, reportTrxWh)
	}

	if err := rows.Err(); err != nil {
		log.Println(err)
	}

	if len(reportsTrxWh) == 0 {
		return "no data or date product not found"
	}

	return reportsTrxWh
}

func NewReportTrxWhRepo(db *sql.DB) ReportTrxWhRepo {
	repo := new(reportTrxWhRepo)

	repo.db = db

	return repo
}
