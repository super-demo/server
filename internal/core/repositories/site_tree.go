package repositories

import (
	"log"
	"server/internal/core/models"

	"gorm.io/gorm"
)

type SiteTreeRepository interface {
	BeginLog() (SiteTreeRepository, error)
	Commit() error
	Rollback() error
	CreateSiteTree(siteTree *models.SiteTree) (*models.SiteTree, error)
	GetListSiteTreeBySiteId(siteId int, userId int) ([]models.GetWorkspaceList, error)
	GetListWorkspaceBySiteIdAndPeople(siteId int, userId int) ([]models.GetWorkspaceList, error)
	GetSiteParentsBySiteId(siteId int) ([]models.SiteTree, error)
	UpdateSiteTree(siteTree *models.SiteTree) (*models.SiteTree, error)
	DeleteSiteTree(siteTree *models.SiteTree) error
}

type siteTreeRepository struct {
	db *gorm.DB
	tx *gorm.DB
}

func NewSiteTreeRepository(db *gorm.DB) SiteTreeRepository {
	return &siteTreeRepository{db: db}
}

func (r *siteTreeRepository) BeginLog() (SiteTreeRepository, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &siteTreeRepository{db: r.db, tx: tx}, nil
}

func (r *siteTreeRepository) Commit() error {
	return r.tx.Commit().Error
}

func (r *siteTreeRepository) Rollback() error {
	return r.tx.Rollback().Error
}

func (r *siteTreeRepository) CreateSiteTree(siteTree *models.SiteTree) (*models.SiteTree, error) {
	if err := r.db.Create(siteTree).Error; err != nil {
		return nil, err
	}

	return siteTree, nil
}

func (r *siteTreeRepository) GetListSiteTreeBySiteId(siteId int, userId int) ([]models.GetWorkspaceList, error) {
	var siteList []models.GetWorkspaceList

	query := `
		SELECT
			s.*,
			st.site_parent_id,
			sp.name AS site_parent_name
		FROM sites s
		INNER JOIN site_trees st ON s.site_id = st.site_child_id
		LEFT JOIN sites sp ON st.site_parent_id = sp.site_id
		WHERE st.site_parent_id = $1
		AND s.deleted_at IS NULL
		ORDER BY s.site_id;
	`

	err := r.db.Raw(query, siteId).Scan(&siteList).Error
	if err != nil {
		return nil, err
	}

	return siteList, nil
}

func (r *siteTreeRepository) GetListWorkspaceBySiteIdAndPeople(siteId int, userId int) ([]models.GetWorkspaceList, error) {
	var siteList []models.GetWorkspaceList

	query := `
		SELECT
			s.*,
			st.site_parent_id,
			sp.name AS site_parent_name
		FROM sites s
		INNER JOIN site_trees st ON s.site_id = st.site_child_id
		LEFT JOIN sites sp ON st.site_parent_id = sp.site_id
		INNER JOIN site_peoples su ON su.site_id = s.site_id
		WHERE st.site_parent_id = $1
		AND su.user_id = $2 
		AND s.deleted_at IS NULL
		ORDER BY s.site_id;
	`

	err := r.db.Raw(query, siteId, userId).Scan(&siteList).Error
	if err != nil {
		return nil, err
	}

	return siteList, nil
}

func (r *siteTreeRepository) GetSiteParentsBySiteId(siteId int) ([]models.SiteTree, error) {
	var siteTree []models.SiteTree

	log.Println("siteId", siteId)

	query := `
		WITH RECURSIVE parent_cte AS (
			SELECT site_parent_id
			FROM site_trees
			WHERE site_child_id = $1

			UNION ALL

			SELECT st.site_parent_id
			FROM site_trees st
			JOIN parent_cte p ON st.site_child_id = p.site_parent_id
		)
		SELECT site_parent_id FROM parent_cte;
	`

	err := r.db.Raw(query, siteId).Scan(&siteTree).Error
	if err != nil {
		return nil, err
	}

	log.Println("siteTree", siteTree)

	return siteTree, nil
}

func (r *siteTreeRepository) UpdateSiteTree(siteTree *models.SiteTree) (*models.SiteTree, error) {
	if err := r.db.Save(siteTree).Error; err != nil {
		return nil, err
	}

	return siteTree, nil
}

func (r *siteTreeRepository) DeleteSiteTree(siteTree *models.SiteTree) error {
	if err := r.db.Delete(siteTree).Error; err != nil {
		return err
	}

	return nil
}
