package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	filename string = "gophermap"
)

// Gophermap parses the given file 'fn' and returns a proper gopher list of items
func Gophermap(fn string) List {
	var l List

	f, err := os.Open(fn)
	defer f.Close()
	if err != nil {
		return Error("gophermap error")
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		if strings.Contains(row, "\t") {
			itemtype, cols := parse(row)
			p, _ := strconv.Atoi(cols[3]) // port
			switch itemtype {
			case G_MENU:
				fallthrough
			case G_TEXT:
				i := Row(itemtype, cols[0], cols[1], cols[2], p)
				l = append(l, i)
			case '!':
				if cols[0] == "!" && cols[1] == "list" {
					l = append(l, ListDir(filepath.Dir(fn))...)
				}
			default:
				i := Row(G_ERROR, strings.Replace(row, "\t", "\\t", -1), "", "", 0)
				l = append(l, i)
			}
		} else {
			l = append(l, Row(G_INFO, row, "", "", 0))
		}
	}

	return l
}

func parse(s string) (byte, []string) {
	cols := strings.Split(s, "\t")
	itemtype := []byte(cols[0][:1])[0]
	f_name := cols[0][1:len(cols[0])]
	f_selector := ""
	f_host := ""
	f_port := ""

	if len(cols) >= 4 {
		f_selector = cols[1]
		f_host = cols[2]
		f_port = cols[3]
	} else if len(cols) >= 3 {
		f_selector = cols[1]
		f_host = cols[2]
	} else if len(cols) >= 2 {
		f_selector = cols[1]
	}

	fields := []string{f_name, f_selector, f_host, f_port}
	return itemtype, fields
}
