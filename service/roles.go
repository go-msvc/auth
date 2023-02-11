package main

import (
	"sync"

	"github.com/go-msvc/errors"
)

type role struct {
	sync.Mutex
	account          *account
	id               string
	name             string
	permissionByID   map[string]*permission
	permissionByName map[string]*permission
}

func (r *role) AddPermission(p *permission) error {
	if r == nil || p == nil {
		return errors.Errorf("r=%p,p=%p", r, p)
	}
	r.Lock()
	defer r.Unlock()
	if _, ok := r.permissionByName[p.name]; ok {
		return nil //already got this permission
	}
	r.permissionByID[p.id] = p
	r.permissionByName[p.name] = p
	return nil
} //role.AddPermission()
