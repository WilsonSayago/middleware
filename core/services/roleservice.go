package services

import (
	"github.com/WilsonSayago/authBase/core"
	"github.com/WilsonSayago/authBase/core/domain"
	"github.com/WilsonSayago/authBase/core/port"
	"github.com/WilsonSayago/initModules"
)

type RoleService struct {
	port port.RolePort
}

func GetRoleServiceInstance(port port.RolePort) core.RoleUseCase {
	instance := initModules.GetInstance("RoleService", func() interface{} {
		return &RoleService{
			port: port,
		}
	})
	return instance.(*RoleService)
}

func (r *RoleService) GetRoleById(id string) (domain.Role, error) {
	return r.port.FindById(id)
}

func (r *RoleService) GetRoles() ([]domain.Role, error) {
	return r.port.FindAll()
}

func (r *RoleService) CreateRole(role domain.Role) (domain.Role, error) {
	return r.port.Save(role)
}
