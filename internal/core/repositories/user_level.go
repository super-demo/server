package repositories

import "server/internal/core/models"

var (
	RootUserLevel = models.UserLevel{
		UserLevelId: 1,
		Slug:        "Root",
		Description: "Good for people who mange everything.",
	}
)

var (
	StaffUserLevel = models.UserLevel{
		UserLevelId: 2,
		Slug:        "Staff",
		Description: "Good for people who manage something.",
	}
)

var (
	OwnerUserLevel = models.UserLevel{
		UserLevelId: 3,
		Slug:        "Owner",
		Description: "Good for people who own everything.",
	}
)

var (
	SuperAdminUserLevel = models.UserLevel{
		UserLevelId: 4,
		Slug:        "Super Admin",
		Description: "Good for people who can manage everything.",
	}
)

var (
	AdminUserLevel = models.UserLevel{
		UserLevelId: 5,
		Slug:        "Admin",
		Description: "Good for people who just need to manage something.",
	}
)

var (
	MemberUserLevel = models.UserLevel{
		UserLevelId: 6,
		Slug:        "Member",
		Description: "Good for poeple who just need to view something.",
	}
)
