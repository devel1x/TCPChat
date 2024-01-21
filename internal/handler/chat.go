package handler

import (
	"fmt"
	"io"
	"strings"
)

var chatHistory []string

type message struct {
	from    string
	payload []byte
}

func (u *user) printMessage() {
	u.mu.Lock()

	msg := <-u.msgch
	if isValidMsg(strings.TrimSpace(string(msg.payload))) {
		for key, value := range connMap {
			if key != u.name {

				value.con.Write([]byte("\n"))
				fmt.Fprint(value.con, u.col, u.formatMsg(msg), u.res)
			}

			fmt.Fprint(value.con, value.color, u.formatMsg(message{key, nil}))
		}
		chatHistory = append(chatHistory, u.col, u.formatMsg(msg), u.res)
	} else {
		u.con.Write([]byte(u.formatMsg(message{u.name, nil})))
	}

	u.mu.Unlock()
}

func (u *user) printHistory() {
	for _, value := range chatHistory {
		u.con.Write([]byte(value))
	}
}

func (u *user) readLoop() {
	defer u.CloseChan()
	buf := make([]byte, 2048)
	for {
		n, err := u.con.Read(buf)
		if err == io.EOF {
			return
		}
		u.msgch <- message{
			from:    u.name,
			payload: buf[:n],
		}
		u.printMessage()
	}
}
