package handler

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var err error

type umap struct {
	con   net.Conn
	color string
}

var connMap = make(map[string]umap)

func (u *user) AcceptLoop() {
	isJoin := false
	defer func() {
		if !isJoin {
			u.CloseChan()
		}
	}()
	u.welcome()
	u.getName()
	if u.name == "" {
		return
	}
	u.mu.Lock()
	connMap[u.name] = umap{u.con, u.col}

	for key, value := range connMap {
		if key != u.name {
			if u.ver == "right" {
				fmt.Fprint(value.con, u.col, "\nWelcome to the club ", u.name, "!\n")
			} else {
				fmt.Fprint(value.con, u.col, "\n", u.name, " has joined our chat...\n")
			}
		}
		if key == u.name {
			u.printHistory()
		}
		fmt.Fprint(value.con, value.color, u.formatMsg(message{key, nil}))
	}

	u.mu.Unlock()
	isJoin = true
	u.readLoop()
}

func (u *user) CloseChan() {
	u.mu.Lock()
	for key, value := range connMap {
		if u.name == "" {
		} else if u.name != key {
			if u.ver == "right" {
				fmt.Fprint(value.con, "\n", u.col, u.name, " has left the gym!\n", u.res, value.color, u.formatMsg(message{key, nil}))
			} else {
				fmt.Fprint(value.con, "\n", u.col, u.name, " has left our chat...\n", u.res, value.color, u.formatMsg(message{key, nil}))
			}
		}
	}
	delete(connMap, u.name)
	delete(used, u.colname)
	close(u.msgch)
	u.mu.Unlock()
}

func (u *user) getName() error {
	name, err := bufio.NewReader(u.con).ReadString('\n')
	if err != nil {
		return err
	}

	name = strings.TrimSpace(name)
	if !u.isValidName(name) {
		return u.getName()
	}

	u.name = name
	return nil
}
