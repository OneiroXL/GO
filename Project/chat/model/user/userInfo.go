package user

import (
	"net"
)

type UserInfo struct {
	UserCode string
	Name string
	Conn net.Conn
}

