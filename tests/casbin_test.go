package tests

import (
	`testing`
	
	`github.com/casbin/casbin/v2`
	adapter "github.com/casbin/gorm-adapter/v3"
	`github.com/chaodoing/boot/database`
	`github.com/chaodoing/boot/logger`
)

func TestCasbin(t *testing.T) {
	env := database.Config{
		Type:     "mysql",
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "123.com",
		Database: "test",
		Charset:  "utf8mb4",
		Logger: logger.Logger{
			Stdout: true,
			Level:  4,
			File:   "./logs/mysql-%F.log",
		},
	}
	
	db, err := env.Connection()
	if err != nil {
		t.Error(err)
	}
	
	orm, err := adapter.NewAdapterByDBUseTableName(db, "develop", "rule")
	if err != nil {
		t.Error(err)
		return
	}
	
	e, err := casbin.NewEnforcer("./rbac_model.conf", orm)
	if err != nil {
		t.Error(err)
		return
	}
	e.EnableLog(true)
	err = e.LoadPolicy()
	if err != nil {
		t.Error(err)
		return
	}
	_, _ = e.AddPolicy("alice", "data1", "read")
	_, _ = e.AddPolicy("bob", "data2", "write")
	_, _ = e.AddGroupingPolicy("admin", "alice")
	_, _ = e.AddGroupingPolicy("admin", "bob")
	var ok bool
	ok, err = e.Enforce("bob", "data2", "write")
	if ok {
		t.Log("Success")
	} else {
		t.Error(err)
	}
	
	_, _ = e.AddRolesForUser("何烨霖", []string{"bob"})
	t.Log(e.GetRolesForUser("何烨霖"))
	_, _ = e.AddRoleForUserInDomain("何烨霖", "admin", "chao")
	t.Log(e.Enforce("何烨霖", "data2", "write"))
}
