package models

import (
	"cooky-go/models"
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
)

var Enforcer *casbin.Enforcer

func Init() *casbin.Enforcer {

	a := gormadapter.NewAdapterByDB(models.DB)
	Enforcer = casbin.NewEnforcer("conf/authz_model.conf", a)

	return Enforcer
}

func AddPolicy(role Role, menus []Menu) {
	DeleteRolePolycy(role)
	for i := 0; i < len(menus); i++ {
		Enforcer.AddPolicy(role.RoleName, menus[i].Url, menus[i].Method)
	}
}

func DeleteRolePolycy(role Role) {
	filteredPolicy := Enforcer.GetFilteredPolicy(0, role.RoleName)
	for _, policy := range filteredPolicy {
		Enforcer.RemovePolicy(policy)
	}
	//Enforcer.DeleteRole(role.RoleName)
}

func RemoveRole(role Role) {
	Enforcer.DeleteRole(role.RoleName)
}

func AddRoleToUser(user User, roles []Role) {
	DeleteUserRole(user)
	for _, role := range roles {
		Enforcer.AddRoleForUser(user.Username, role.RoleName)
	}
}

func DeleteUserRole(user User) {
	Enforcer.DeleteRolesForUser(user.Username)
}
