package repositories

import "server/internal/core/models"

var (
	OwnerUserLevel = models.UserLevel{
		UserLevelId: 1,
		Slug:        "Owner",
		Description: "Good for people who own everything.",
	}
)

var (
	SuperAdminUserLevel = models.UserLevel{
		UserLevelId: 2,
		Slug:        "Super Admin",
		Description: "Good for people who can manage everything.",
	}
)

var (
	AdminUserLevel = models.UserLevel{
		UserLevelId: 3,
		Slug:        "Admin",
		Description: "Good for people who just need to manage something.",
	}
)

var (
	MemberUserLevel = models.UserLevel{
		UserLevelId: 4,
		Slug:        "Member",
		Description: "Good for poeple who just need to view something.",
	}
)
