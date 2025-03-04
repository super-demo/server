package usecases

import (
	"server/internal/core/models"
	"server/internal/core/repositories"
)

type SiteTreeUsecase interface {
	CreateSiteTree(siteTree *models.SiteTree, requesterUserId int) (*models.SiteTree, error)
	GetListSiteTreeBySiteId(siteId int, requesterUserId int) ([]models.SiteTree, error)
	UpdateSiteTree(siteTree *models.SiteTree, requesterUserId int) (*models.SiteTree, error)
	DeleteSiteTree(siteTree *models.SiteTree, requesterUserId int) error
}

type siteTreeUsecase struct {
	siteTreeRepo repositories.SiteTreeRepository
	siteRepo     repositories.SiteRepository
	siteUserRepo repositories.SiteUserRepository
	siteLogRepo  repositories.SiteLogRepository
}

func NewSiteTreeUsecase(siteTreeRepo repositories.SiteTreeRepository, siteRepo repositories.SiteRepository, siteUserRepo repositories.SiteUserRepository, siteLogRepo repositories.SiteLogRepository) SiteTreeUsecase {
	return &siteTreeUsecase{
		siteTreeRepo: siteTreeRepo,
		siteRepo:     siteRepo,
		siteUserRepo: siteUserRepo,
		siteLogRepo:  siteLogRepo,
	}
}

func (u *siteTreeUsecase) CreateSiteTree(siteTree *models.SiteTree, requesterUserId int) (*models.SiteTree, error) {
	txSiteTreeRepo, err := u.siteTreeRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txSiteTreeRepo.Rollback()
		}
	}()

	siteTree.CreatedBy = requesterUserId
	siteTree.UpdatedBy = requesterUserId
	newSiteTree, err := txSiteTreeRepo.CreateSiteTree(siteTree)
	if err != nil {
		txSiteTreeRepo.Rollback()
		return nil, err
	}

	txSiteTreeRepo.Commit()
	return newSiteTree, nil
}

func (u *siteTreeUsecase) GetListSiteTreeBySiteId(siteId int, requesterUserId int) ([]models.SiteTree, error) {
	siteTrees, err := u.siteTreeRepo.GetListSiteTreeBySiteId(siteId)
	if err != nil {
		return nil, err
	}

	return siteTrees, nil
}

func (u *siteTreeUsecase) UpdateSiteTree(siteTree *models.SiteTree, requesterUserId int) (*models.SiteTree, error) {
	txSiteTreeRepo, err := u.siteTreeRepo.BeginLog()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txSiteTreeRepo.Rollback()
		}
	}()

	siteTree.UpdatedBy = requesterUserId
	updatedSiteTree, err := txSiteTreeRepo.UpdateSiteTree(siteTree)
	if err != nil {
		txSiteTreeRepo.Rollback()
		return nil, err
	}

	txSiteTreeRepo.Commit()
	return updatedSiteTree, nil
}

func (u *siteTreeUsecase) DeleteSiteTree(siteTree *models.SiteTree, requesterUserId int) error {
	txSiteTreeRepo, err := u.siteTreeRepo.BeginLog()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			txSiteTreeRepo.Rollback()
		}
	}()

	err = txSiteTreeRepo.DeleteSiteTree(siteTree)
	if err != nil {
		txSiteTreeRepo.Rollback()
		return err
	}

	txSiteTreeRepo.Commit()
	return nil
}
