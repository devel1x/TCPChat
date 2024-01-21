package handler

import (
	"fmt"
	"os"
	"time"
)

func (u *user) welcome() {
	var welcommsg []byte
	if u.ver == "right" {
		welcommsg, _ = os.ReadFile("static/helloRight.txt")
	} else {
		welcommsg, _ = os.ReadFile("static/hello.txt")
	}

	fmt.Fprint(u.con, u.col, string(welcommsg))
}

func (u *user) isValidName(name string) bool {
	if name == "" {
		if u.ver == "right" {
			fmt.Fprint(u.con, u.col, "Trow me some letters, Buddy\nPlease try again...\n[ENTER YOUR NAME]:")
		} else {
			fmt.Fprint(u.con, u.col, "Username must be from 2 to 15 characters and contain only Ascii code characters\n[ENTER YOUR NAME]:")
		}
		return false
	}

	var count int
	for range name {
		count++
	}
	if count > 15 {
		if u.ver == "right" {
			fmt.Fprint(u.con, u.col, "AAAAgh, its too big, Buddy\nPlease try again...\n[ENTER YOUR NAME]:")
		} else {
			fmt.Fprint(u.con, u.col, "Username must be from 2 to 15 characters and contain only Ascii code characters\n[ENTER YOUR NAME]:")
		}
		return false
	}

	if count < 2 {
		if u.ver == "right" {
			fmt.Fprint(u.con, u.col, "Too short for the club, Buddy\nPlease try again...\n[ENTER YOUR NAME]:")
		} else {
			fmt.Fprint(u.con, u.col, "Username must be from 2 to 15 characters and contain only Ascii code characters\n[ENTER YOUR NAME]:")
		}
		return false
	}

	for _, ch := range name {
		if ch >= 0 && ch <= 31 {
			if u.ver == "right" {
				fmt.Fprint(u.con, u.col, "some LETTERS, Buddy\nPlease try again...\n[ENTER YOUR NAME]:")
			} else {
				fmt.Fprint(u.con, u.col, "Username must be from 2 to 15 characters and contain only Ascii code characters\n[ENTER YOUR NAME]:")
			}
			
			return false
		}
	}
	u.mu.Lock()
	if _, ok := connMap[name]; ok {
		if u.ver == "right" {
			fmt.Fprint(u.con, u.col, "He is already in the club Buddy, throw me some other letters\nPlease try again...\n[ENTER YOUR NAME]:")
		} else {
			fmt.Fprint(u.con, u.col, "Username is already taken\nPlease try again...\n[ENTER YOUR NAME]:")
		}
		u.mu.Unlock()
		return false
	}
	u.mu.Unlock()
	return true
}

func isValidMsg(msg string) bool {
	if msg == "" {
		return false
	}
	for _, ch := range msg {
		if ch >= 0 && ch <= 31 {
			return false
		}
	}
	return true
}

func (u *user) formatMsg(msg message) string {
	return fmt.Sprintf("[%s][%s]:%s", string(time.Now().Format("2006-01-02 15:04:05")), msg.from, string(msg.payload))
}
