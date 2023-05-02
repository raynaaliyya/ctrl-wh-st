package usecase

import "go_inven_ctrl/repository"

type ReportTrxWhUsecase interface {
	FindAllReportTrxWh() any
	FindByIdReportTrxWh(id string) any
	FindByDateReportTrxWh(date string) any
}

type reportTrxWhUsecase struct {
	reportTrxWhRepo repository.ReportTrxWhRepo
}

func (u *reportTrxWhUsecase) FindAllReportTrxWh() any {
	return u.reportTrxWhRepo.GetAllReportTrxWh()
}

func (u *reportTrxWhUsecase) FindByIdReportTrxWh(id string) any {
	return u.reportTrxWhRepo.GetByIdReportTrxWh(id)
}

func (u *reportTrxWhUsecase) FindByDateReportTrxWh(date string) any {
	return u.reportTrxWhRepo.GetByDateReportTrxWh(date)
}

func NewReportTrxWhUsecase(repoTrxWh repository.ReportTrxWhRepo) ReportTrxWhUsecase {
	return &reportTrxWhUsecase{
		reportTrxWhRepo: repoTrxWh,
	}
}
