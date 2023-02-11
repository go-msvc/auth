package main

import (
	"sort"
	"strings"
	"sync"

	"github.com/google/uuid"
)

type account struct {
	sync.Mutex
	id               string
	userByID         map[string]*user
	userByName       map[string]*user
	roleByID         map[string]*role
	roleByName       map[string]*role
	permissionByID   map[string]*permission
	permissionByName map[string]*permission
}

var (
	accountsMutex sync.Mutex
	accounts      = map[string]*account{}
)

func AddAccount() (*account, error) {
	acc := &account{
		id:               uuid.New().String(),
		userByID:         map[string]*user{},
		userByName:       map[string]*user{},
		roleByID:         map[string]*role{},
		roleByName:       map[string]*role{},
		permissionByID:   map[string]*permission{},
		permissionByName: map[string]*permission{},
	}

	accountsMutex.Lock()
	defer accountsMutex.Unlock()
	accounts[acc.id] = acc
	return acc, nil
} //AddAccount()

func GetAccount(id string) (*account, bool) {
	if acc, ok := accounts[id]; ok {
		return acc, true
	}
	return nil, false
}

func DelAccount(id string) error {
	accountsMutex.Lock()
	defer accountsMutex.Unlock()
	delete(accounts, id)
	return nil
}

func (acc *account) AddPermission(name string) (*permission, error) {
	p := &permission{
		account: acc,
		name:    name,
		id:      uuid.New().String(),
	}
	acc.Lock()
	defer acc.Unlock()
	acc.permissionByID[p.id] = p
	acc.permissionByName[p.name] = p
	return p, nil
} //Account.AddPermission()

func (acc *account) findPermissions(filter string, limit int) []*permission {
	names := []string{}
	for n := range acc.permissionByName {
		if filter == "" || strings.Index(n, filter) >= 0 {
			names = append(names, n)
		}
	}
	sort.Slice(names, func(i, j int) bool { return names[i] > names[j] })
	if len(names) > limit {
		names = names[0:limit]
	}
	list := []*permission{}
	for _, n := range names {
		p := acc.permissionByName[n]
		list = append(list, p)
	}
	return list
} //account.findPermissions()

func (acc *account) AddRole(name string) (*role, error) {
	r := &role{
		account: acc,
		name:    name,
		id:      uuid.New().String(),
	}
	acc.Lock()
	defer acc.Unlock()
	acc.roleByID[r.id] = r
	acc.roleByName[r.name] = r
	return r, nil
} //Account.AddRole()

func (acc *account) findRoles(filter string, limit int) []*role {
	names := []string{}
	for n := range acc.roleByName {
		if filter == "" || strings.Index(n, filter) >= 0 {
			names = append(names, n)
		}
	}
	sort.Slice(names, func(i, j int) bool { return names[i] > names[j] })
	if len(names) > limit {
		names = names[0:limit]
	}
	list := []*role{}
	for _, n := range names {
		r := acc.roleByName[n]
		list = append(list, r)
	}
	return list
} //account.findRoles()

func (acc *account) AddUser(name string) (*user, error) {
	u := &user{
		account: acc,
		name:    name,
		id:      uuid.New().String(),
	}
	acc.Lock()
	defer acc.Unlock()
	acc.userByID[u.id] = u
	acc.userByName[u.name] = u
	return u, nil
} //Account.AddUser()

func (acc *account) findUsers(filter string, limit int) []*user {
	names := []string{}
	for n := range acc.userByName {
		if filter == "" || strings.Index(n, filter) >= 0 {
			names = append(names, n)
		}
	}
	sort.Slice(names, func(i, j int) bool { return names[i] > names[j] })
	if len(names) > limit {
		names = names[0:limit]
	}
	list := []*user{}
	for _, n := range names {
		u := acc.userByName[n]
		list = append(list, u)
	}
	return list
} //account.findUsers()
