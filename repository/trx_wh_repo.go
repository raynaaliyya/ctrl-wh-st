package repository

import (
	"database/sql"
	"fmt"
	"go_inven_ctrl/entity"
	"log"
	"time"
)

type TrxWhRepo interface {
	EnrollInsertTrxWh(EnrollTrxWh *entity.TrxWh) string
}

type trxWhRepo struct {
	db *sql.DB
}

func (trx *trxWhRepo) EnrollInsertTrxWh(EnrollTrxWh *entity.TrxWh) string {
	tx, err := trx.db.Begin()
	if err != nil {
		return "Transaction rollback"
	}

	trx.InsertTrxWh(EnrollTrxWh, tx)
	name, lastStock := trx.GetProductName(EnrollTrxWh.ProductWhId, tx)

	action := trx.GetAct(EnrollTrxWh.ProductWhId, tx)
	if action == "input" {
		addStock := trx.SumTrxProductWh(EnrollTrxWh.ProductWhId, tx)
		fmt.Println(addStock)
		trx.UpdateProductWh(addStock, EnrollTrxWh.ProductWhId, tx)
	} else if action == "output" {
		minStock := trx.MinTrxProductWh(EnrollTrxWh.ProductWhId, tx)
		fmt.Println(minStock)
		trx.UpdateProductWh(minStock, EnrollTrxWh.ProductWhId, tx)
	}

	trx.InsertReportWh(EnrollTrxWh, name, lastStock, tx)
	trx.DeleteTrxWh(EnrollTrxWh.ProductWhId, tx)

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return "product not found"
	} else {
		return "transaction in warehouse committed"
	}
}

func (trx *trxWhRepo) InsertTrxWh(enrollTrxWh *entity.TrxWh, tx *sql.Tx) {
	query := "INSERT INTO trx_wh (id, product_wh_id, product_name, stock, act, created_at) VALUES($1, $2, $3, $4, $5, $6)"

	_, err := tx.Exec(query, &enrollTrxWh.ID, &enrollTrxWh.ProductWhId, &enrollTrxWh.ProductName, &enrollTrxWh.Stock, &enrollTrxWh.Act, &enrollTrxWh.CreatedAt)

	Validate(err, "Insert", tx)

	fmt.Println(enrollTrxWh.Act)
}

func (trx *trxWhRepo) InsertReportWh(enrollTrxWh *entity.TrxWh, name string, lastStock int, tx *sql.Tx) {
	query := "INSERT INTO report_trx_wh (product_wh_id, stock, product_name, act, last_stock, created_at) VALUES($1, $2, $3, $4, $5, $6)"

	enrollTrxWh.CreatedAt = time.Now()

	_, err := tx.Exec(query, enrollTrxWh.ProductWhId, enrollTrxWh.Stock, name, enrollTrxWh.Act, lastStock, enrollTrxWh.CreatedAt)

	Validate(err, "Insert", tx)
}

func (trx *trxWhRepo) GetProductName(productWhId string, tx *sql.Tx) (string, int) {
	query := "SELECT product_wh.product_name, product_wh.stock from trx_wh JOIN product_wh on trx_wh.product_wh_id = product_wh.id WHERE product_wh_id = $1"

	var name string
	var stock int

	err := tx.QueryRow(query, productWhId).Scan(&name, &stock)

	fmt.Println(name)

	Validate(err, "Select", tx)

	return name, stock
}

func (trx *trxWhRepo) GetAct(productWhId string, tx *sql.Tx) string {
	query := "SELECT act FROM trx_wh WHERE product_wh_id = $1"

	var act string
	err := tx.QueryRow(query, productWhId).Scan(&act)

	Validate(err, "Select", tx)

	return act
}

func (trx *trxWhRepo) SumTrxProductWh(productWhId string, tx *sql.Tx) int {
	query := "SELECT SUM(product_wh.stock + trx_wh.stock) FROM trx_wh JOIN product_wh ON trx_wh.product_wh_id = product_wh.id WHERE trx_wh.product_wh_id = $1"

	stock := 0
	err := tx.QueryRow(query, productWhId).Scan(&stock)

	Validate(err, "Select", tx)

	return stock
}

func (trx *trxWhRepo) MinTrxProductWh(productWhId string, tx *sql.Tx) int {
	query := "SELECT SUM(product_wh.stock - trx_wh.stock) FROM trx_wh JOIN product_wh ON trx_wh.product_wh_id = product_wh.id WHERE trx_wh.product_wh_id = $1"

	stock := 0
	err := tx.QueryRow(query, productWhId).Scan(&stock)

	Validate(err, "Select", tx)

	return stock
}

func (trx *trxWhRepo) UpdateProductWh(stock int, productWhId string, tx *sql.Tx) {
	query := "UPDATE product_wh SET stock = $1 WHERE id = $2"
	_, err := tx.Exec(query, stock, productWhId)

	Validate(err, "Update", tx)
}

func (trx *trxWhRepo) DeleteTrxWh(productWhId string, tx *sql.Tx) {
	query := "DELETE FROM trx_wh WHERE product_wh_id = $1"
	_, err := tx.Exec(query, productWhId)

	Validate(err, "Delete", tx)
}

func Validate(err error, message string, tx *sql.Tx) {
	if err != nil {
		tx.Rollback()
		fmt.Println(err, "Transaction Rollback")
	} else {
		fmt.Println("successfully " + message + " data")
	}
}

func NewTrxWhRepo(db *sql.DB) TrxWhRepo {
	repo := new(trxWhRepo)

	repo.db = db

	return repo
}
