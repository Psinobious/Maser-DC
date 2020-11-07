package domain

type Role struct {
	RoleID      string `json:"RoleID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
	RoleType    string `json:"RoleType"`
}

type RoleRepository interface {
	FindById(roleID string) (Role, error)
	FindByRoleTypes(RoleType string) ([]Role, error)
	RemoveRole(roleID string) error
	UpdateRole(role Role) error
	Store(role Role) error
}
