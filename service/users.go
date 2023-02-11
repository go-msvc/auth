package main

import "sync"

type user struct {
	sync.Mutex
	account    *account
	id         string
	name       string
	roleByID   map[string]*role
	roleByName map[string]*role
}
