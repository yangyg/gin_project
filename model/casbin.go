package model

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql" // sql driver
)

// CasbinModel struct
type CasbinModel struct {
	RoleName string `json:"role_name"`
	Path     string `json:"path"`
	Method   string `json:"method"`
}

// AddPolicy 增加策略
func (cm *CasbinModel) AddPolicy() (bool, error) {
	e := Casbin()
	return e.AddPolicy(cm.RoleName, cm.Path, cm.Method)
}

// RemovePolicy 删除策略
func (cm *CasbinModel) RemovePolicy() (bool, error) {
	e := Casbin()
	return e.RemovePolicy(cm.RoleName, cm.Path, cm.Method)
}

// AddRoleForUser  为用户添加角色
func AddRoleForUser(user string, role string) (bool, error) {
	e := Casbin()
	return e.AddRoleForUser(user, role)
}

// DeleteRoleForUser 删除用户的角色
func DeleteRoleForUser(user string, role string) (bool, error) {
	e := Casbin()
	return e.DeleteRoleForUser(user, role)
}

//Casbin 持久化到数据库
func Casbin() *casbin.Enforcer {
	a, _ := gormadapter.NewAdapterByDB(Db)
	e, _ := casbin.NewEnforcer("config/model.conf", a)
	e.LoadPolicy()
	return e
}
