package service

import "go-machine-boilerplate/internal/accessmanager/domain"

type AccessManagerService struct {
	domain *domain.AccessPolicyManager
}

func NewAccessManagerService() *AccessManagerService {
	return &AccessManagerService{
		domain: domain.NewAccessPolicyManager(),
	}
}

func (s *AccessManagerService) AddRole(role string) bool {
	if s.domain.Roles[role] {
		return false
	}

	s.domain.Roles[role] = true
	s.domain.Permissions[role] = make(domain.RoleResourcePermission)
	return true
}

func (s *AccessManagerService) AddResource(resource string) bool {
	if s.domain.Resources[resource] {
		return false
	}

	s.domain.Resources[resource] = true
	return true
}

func (s *AccessManagerService) AddPermissions(role, resource string, permissions []string) bool {
	if !s.domain.Roles[role] || !s.domain.Resources[resource] {
		return false
	}

	if s.domain.Permissions[role][resource] == nil {
		s.domain.Permissions[role][resource] = make(domain.PermissionSet)
	}

	for _, permission := range permissions {
		s.domain.Permissions[role][resource][permission] = true
	}

	return true
}

func (s *AccessManagerService) CheckAccess(role, resource, permission string) bool {
	return s.domain.Permissions[role][resource][permission]
}
