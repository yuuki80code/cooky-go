package models

import (
	"cooky-go/models"
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
)

var Enforcer *casbin.Enforcer

func Init() {

	a := gormadapter.NewAdapterByDB(models.DB)
	Enforcer = casbin.NewEnforcer("conf/authz_model.conf", a)

	return
}

func AddPolicy(role Role, menus []Menu) {
	DeleteRolePolycy(role)
	for i := 0; i < len(menus); i++ {
		Enforcer.AddPolicy(role.RoleName, menus[i].Url, menus[i].Method)
	}
}

func DeleteRolePolycy(role Role) {
	Enforcer.DeleteRole(role.RoleName)
}