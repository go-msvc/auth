package main

import (
	"github.com/go-msvc/config"
	"github.com/go-msvc/logger"
	_ "github.com/go-msvc/nats-utils"
	"github.com/go-msvc/utils/ms"
)

var log = logger.New().WithLevel(logger.LevelDebug)

func main() {
	svc := ms.New(
		//accounts
		ms.WithOper("add_account", operAddAccount),
		ms.WithOper("get_account", operGetAccount), //can return list of users/roles/permissions too...
		ms.WithOper("del_account", operDelAccount),

		//users
		// ms.WithOper("findUsers", operFindUsers),
		ms.WithOper("add_account_user", operAddAccountUser),
		ms.WithOper("add_user_role", operAddUserRole),
		//ms.WithOper("getUser", operGetUser),
		// ms.WithOper("updUser", operUpdUser),
		// ms.WithOper("delUser", operDelUser),
		// //roles
		// ms.WithOper("findUsers", operFindUsers),
		ms.WithOper("add_account_role", operAddAccountRole),
		ms.WithOper("add_role_permission", operAddRolePermission),
		// ms.WithOper("getUser", operGetUser),
		// ms.WithOper("updUser", operUpdUser),
		// ms.WithOper("delUser", operDelUser),
		// //permissions
		// ms.WithOper("findUsers", operFindUsers),
		ms.WithOper("add_account_permission", operAddAccountPermission),
		// ms.WithOper("getUser", operGetUser),
		// ms.WithOper("updUser", operUpdUser),
		// ms.WithOper("delUser", operDelUser),
	)
	if err := config.AddSource("file", config.File("./conf/config.json")); err != nil {
		panic(err)
	}
	if err := config.Load(); err != nil {
		panic(err)
	}
	svc.Serve()
} //main()
