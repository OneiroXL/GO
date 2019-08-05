package user

import (
	"net"
)

type UserInfo struct {
	ID int
	Name string
	Conn net.Conn
}

