package handler

import (
	"net"
	"sync"
)

type UserI interface {
	AcceptLoop()
	CloseChan()
}

type user struct {
	name    string
	msgch   chan message
	con     net.Conn
	mu      *sync.Mutex
	ver     string
	col     string
	colname string
	res     string
}

func NewUser(mu *sync.Mutex, con net.Conn, ver string) (UserI, bool) {
	mu.Lock()
	defer mu.Unlock()
	if len(used) > 2 {
		return nil, false
	}
	var col, colname string
	for key, value := range colors {
		if _, ok := used[key]; !ok {
			col = value
			colname = key
			used[key] = value
			break
		}
	}

	return &user{
		msgch:   make(chan message, 1),
		mu:      mu,
		con:     con,
		ver:     ver,
		col:     col,
		colname: colname,
		res:     "\u001b[0m",
	}, true
}

var colors = map[string]string{
	"red":      "\u001b[31m",
	"green":    "\u001b[32m",
	"yellow":   "\u001b[33m",
	"blue":     "\u001b[34m",
	"magenta":  "\u001b[35m",
	"cyan":     "\u001b[36m",
	"white":    "\u001b[37m",
	"pencil":   "\u001b[38;2;253;182;0m",
	"lavender": "\u001b[38;5;147m",
	"pink":     "\u001b[38;5;201m",
}

var used = make(map[string]string)
