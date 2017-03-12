package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
)

type item struct {
	Type     byte
	Name     string
	Selector string
	Host     string
	Port     int
}
type list []item

func (i item) String() string {
	switch i.Type {
	case 'i':
		return fmt.Sprintf("%c%s\t\tinfo.host\t1\r\n",
			i.Type, i.Name)
	default:
		return fmt.Sprintf("%c%s\t%s\t%s\t%d\r\n",
			i.Type, i.Name, i.Selector, i.Host, i.Port)
	}
}

func (l list) String() string {
	var b bytes.Buffer
	for _, i := range l {
		fmt.Fprint(&b, i)
	}
	fmt.Fprint(&b, ".\r\n")
	return b.String()
}

func handleConn(c net.Conn) {
	defer c.Close()
	buf := bufio.NewReader(c)

	log.Printf("Open: %v", c.RemoteAddr())

	req, _, err := buf.ReadLine()
	if err != nil {
		fmt.Fprint(c, responseError("Invalid request."))
	}
	log.Printf("%v: %s", c.RemoteAddr(), req)

	handleRequest(string(req), c)

	log.Printf("Close: %v", c.RemoteAddr())
}

func handleRequest(s string, c net.Conn) {
	if s == "" || s == "/" {
		s = "."
	}
	s = *root + filepath.Clean("/"+s)
	f, err := os.Open(s)
	defer f.Close()
	if err != nil {
		fmt.Fprint(c, responseError("Resource not found."))
		return
	}
	fi, _ := f.Stat()
	if fi.IsDir() {
		var l list
		filepath.Walk(s, (func(p string, info os.FileInfo, err error) error {
			log.Println(p)
			if p == "." {
				return nil
			}
			if info.IsDir() {
				if len(l) > 0 {
					l = append(l, item{'1', info.Name(), p[len(*root)-1:], *host, *port})
					return filepath.SkipDir
				}
				l = append(l, item{Type: 'i', Name: p[len(*root):]})
				return nil
			}
			l = append(l, item{'0', info.Name(), p[len(*root)-1:], *host, *port})
			return nil
		}))
		fmt.Fprint(c, l)
		return
	}
	io.Copy(c, f)
}

func responseError(s string) list {
	return list{item{Type: '3', Name: s}}
}

func ListenAndServe(addr string) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Serving %s at port %d", *root, *port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}
