package repositories

import (
	"server/internal/core/models"

	"gorm.io/gorm"
)

type SiteTreeRepository interface {
	BeginLog() (SiteTreeRepository, error)
	Commit() error
	Rollback() error
	CreateSiteTree(siteTree *models.SiteTree) (*models.SiteTree, error)
	GetListSiteTreeBySiteId(siteId int) ([]models.SiteTree, error)
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

func (r *siteTreeRepository) GetListSiteTreeBySiteId(siteId int) ([]models.SiteTree, error) {
	var siteTrees []models.SiteTree

	query := `
	WITH RECURSIVE site_hierarchy AS (
		SELECT
			s.site_id,
			s.name,
			s.site_type_id,
			s.description,
			st.site_parent_id,
			1 AS depth,
			ARRAY[s.site_id] AS path
		FROM sites s
		LEFT JOIN site_trees st ON s.site_id = st.site_child_id
		WHERE s.site_id = ?
		
		UNION ALL
		
		SELECT
			s.site_id,
			s.name,
			s.site_type_id,
			s.description,
			st.site_parent_id,
			sh.depth + 1,
			sh.path || s.site_id
		FROM sites s
		INNER JOIN site_trees st ON s.site_id = st.site_child_id
		INNER JOIN site_hierarchy sh ON st.site_parent_id = sh.site_id
	)
	SELECT * FROM site_hierarchy ORDER BY path;
	`

	err := r.db.Raw(query, siteId).Scan(&siteTrees).Error
	if err != nil {
		return nil, err
	}

	return siteTrees, nil
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
