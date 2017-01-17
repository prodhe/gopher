/*
gopher - Set up a server to communicate over gopher

https://tools.ietf.org/html/rfc1436
*/
package main

import "net"

const (
	host string = "localhost"
	port string = "7070"
	root string = "."
)

func main() {
	addr := net.JoinHostPort(host, port)
	ListenAndServe(addr)
}
