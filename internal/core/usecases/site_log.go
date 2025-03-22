package usecases

import (
	"server/internal/core/models"
	"server/internal/core/repositories"
)

type SiteLogUsecase interface {
	GetListSiteLog() ([]models.SiteLog, error)
	GetSiteLogBySiteId(siteId int) ([]models.SiteLog, error)
}

type siteLogUsecase struct {
	siteLogRepo repositories.SiteLogRepository
}

func NewSiteLogUsecase(siteLogRepo repositories.SiteLogRepository) SiteLogUsecase {
	return &siteLogUsecase{
		siteLogRepo: siteLogRepo,
	}
}

func (u *siteLogUsecase) GetListSiteLog() ([]models.SiteLog, error) {
	siteLogs, err := u.siteLogRepo.GetListSiteLog()
	if err != nil {
		return nil, err
	}

	return siteLogs, nil
}

func (u *siteLogUsecase) GetSiteLogBySiteId(siteId int) ([]models.SiteLog, error) {
	siteLogs, err := u.siteLogRepo.GetSiteLogBySiteId(siteId)
	if err != nil {
		return nil, err
	}

	return siteLogs, nil
}
