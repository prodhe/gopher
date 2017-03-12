/*
gopher - Set up a server to communicate over gopher

https://tools.ietf.org/html/rfc1436
*/
package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
)

var (
	help bool
	host string
	port int
	root string
)

func init() {
	flag.BoolVar(&help, "h", false, "Show usage")
	flag.StringVar(&host, "a", "localhost", "Public host `address`")
	flag.IntVar(&port, "p", 70, "Listening `port`")
	flag.StringVar(&root, "d", "/var/gopher", "Root `directory` to serve")
}

func main() {
	flag.Parse()

	if help {
		fmt.Println("usage: gopher [options]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// check trailing slash for root directory and append if needed
	if root != "" {
		if root[len(root)-1:] != "/" {
			root = root + "/"
		}
	}

	addr := net.JoinHostPort("0.0.0.0", strconv.Itoa(port))
	ListenAndServe(addr)
}
