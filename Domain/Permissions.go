package domain

type Permissions struct {
	RoleID string
	UserID string
}

type PermissionsRepository interface {
	AddPermission(RoleID string, UserID string) error
	CheckPermission(userID string, roleID string) (bool, error)
	Permissions(userID string) ([]string, error)
	DeletePermission(RoleID string, UserID string) error
	Purge(UserID string) error
}
