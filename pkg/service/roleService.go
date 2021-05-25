package service

import (
	"apiserver/pkg/model"
	"apiserver/pkg/repository"
)

type RoleService struct {
	repoRoles repository.Role
}

func NewRoleService(repo repository.Role) *RoleService {
	return &RoleService{repoRoles: repo}
}

func (r *RoleService) CreateRole(role model.Role) (int64, error) {
	return r.repoRoles.CreateRole(role)
}
