package usecase

import (
	"go_inven_ctrl/entity"
	"go_inven_ctrl/repository"
)

type TrxWhUsecase interface {
	EnrollInsertTrxWh(EnrollTrxWh *entity.TrxWh) string
}

type trxWhUSecase struct {
	trxWhRepo repository.TrxWhRepo
}

func (tx *trxWhUSecase) EnrollInsertTrxWh(EnrollTrxWh *entity.TrxWh) string {
	return tx.trxWhRepo.EnrollInsertTrxWh(EnrollTrxWh)
}

func NewTrxWhUsecase(trxWh repository.TrxWhRepo) TrxWhUsecase {
	return &trxWhUSecase{
		trxWhRepo: trxWh,
	}
}
