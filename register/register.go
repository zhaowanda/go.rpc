package register

import "net"

type IRegister interface {
	Register(string, interface{}, ...string) error
	Start() error
	HandleConnAccept(conn net.Conn) (net.Conn, bool)
	Unregister(name string) error
	Name() string
}
