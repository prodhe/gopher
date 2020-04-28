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
	"path/filepath"
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
	flag.StringVar(&root, "d", "/var/gopher/", "Root `directory` to serve")
}

func main() {
	flag.Parse()

	if help || root == "" {
		fmt.Println("usage: gopher [options]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// check and correct root directory
	if !filepath.IsAbs(root) {
		path, err := filepath.Abs(root)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
		}
		root = path
	}
	if root[len(root)-1:] != "/" {
		root = root + "/"
	}

	addr := net.JoinHostPort("0.0.0.0", strconv.Itoa(port))
	ListenAndServe(addr)
}
