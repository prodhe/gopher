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
	help = flag.Bool("h", false, "Show usage")
	host = flag.String("a", "localhost", "Public host `address`")
	port = flag.Int("p", 70, "Listening `port`")
	root = flag.String("d", "/var/gopher", "Root `directory` to serve")
)

func main() {
	flag.Parse()

	if *help {
		fmt.Println("usage: gopher [options]")
		flag.PrintDefaults()
		os.Exit(1)
	}

	addr := net.JoinHostPort("0.0.0.0", strconv.Itoa(*port))
	ListenAndServe(addr)
}
