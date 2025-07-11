package domain

type PermissionSet map[string]bool

type RoleResourcePermission map[string]PermissionSet

type AccessPolicyManager struct {
	Roles       map[string]bool
	Resources   map[string]bool
	Permissions map[string]RoleResourcePermission // role -> resource -> permissions
}

func NewAccessPolicyManager() *AccessPolicyManager {
	return &AccessPolicyManager{
		Roles:       make(map[string]bool),
		Resources:   make(map[string]bool),
		Permissions: make(map[string]RoleResourcePermission),
	}
}
