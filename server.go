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

const (
	G_ERROR = '3'
	G_INFO  = 'i'
	G_MENU  = '1'
	G_TEXT  = '0'
)

type Item struct {
	Type     byte
	Name     string
	Selector string
	Host     string
	Port     int
}
type List []Item

func (i Item) String() string {
	switch i.Type {
	case 'i':
		return fmt.Sprintf("%c%s\t\tinfo.host\t1\r\n",
			i.Type, i.Name)
	default:
		return fmt.Sprintf("%c%s\t%s\t%s\t%d\r\n",
			i.Type, i.Name, i.Selector, i.Host, i.Port)
	}
}

func (l List) String() string {
	var b bytes.Buffer
	for _, i := range l {
		fmt.Fprint(&b, i)
	}
	fmt.Fprint(&b, ".\r\n")
	return b.String()
}

// Row returns a gopher item ready to be served
func Row(t byte, n, s, h string, p int) Item {
	switch t {
	case G_ERROR:
		s = ""
		h = "error.host"
		p = 1
	case G_INFO:
		s = ""
		h = "info.host"
		p = 1
	case G_MENU:
		if h == "" {
			h = host
			p = port
		}
	case G_TEXT:
		if h == "" {
			h = host
			p = port
		}
	default:
		return Row(G_ERROR, "Internal server error", "", "", 0)
	}
	return Item{t, n, s, h, p}
}

// Exists returns whether the given file or directory exists or not
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// ListDir scans the given 'path' and returns a gopher list of entries
func ListDir(path string) List {
	var l List
	filepath.Walk(path, (func(p string, info os.FileInfo, err error) error {
		if p == "." {
			return nil
		}
		if info.IsDir() {
			if len(l) > 0 {
				l = append(l, Row(G_MENU, info.Name(), p[len(root)-1:], host, port))
				return filepath.SkipDir
			}
			l = append(l, Row(G_INFO, p[len(root):], "", "", 0))
			return nil
		}
		l = append(l, Row(G_TEXT, info.Name(), p[len(root)-1:], host, port))
		return nil
	}))
	return l
}

// ListenAndServe starts a gopher server at 'addr'
func ListenAndServe(addr string) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Serving %s at %s:%d", root, host, port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

// handleConn manages open and close of network conn
func handleConn(c net.Conn) {
	defer c.Close()

	buf := bufio.NewReader(c)
	req, _, err := buf.ReadLine()
	if err != nil {
		fmt.Fprint(c, responseError("Invalid request."))
	}

	log.Printf("%v: %s", c.RemoteAddr(), req)

	handleRequest(string(req), c)
}

// handleRequest parses the request and sends an answer
func handleRequest(req string, c net.Conn) {
	req = root + filepath.Clean("/"+req)

	f, err := os.Open(req)
	defer f.Close()
	if err != nil {
		fmt.Fprint(c, responseError("Resource not found."))
		return
	}

	fi, _ := f.Stat()
	if fi.IsDir() {
		var l List
		if ok, err := Exists(req + "/gophermap"); ok == true && err == nil {
			l = Gophermap(req + "/gophermap")
		} else {
			l = ListDir(req)
		}
		fmt.Fprint(c, l)
		return
	}

	io.Copy(c, f)
	return
}

// responseError returns a full response with a gopher-formatted error 's'
func responseError(s string) List {
	return List{Row(G_ERROR, s, "", "", 0)}
}
