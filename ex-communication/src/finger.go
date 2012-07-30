package main

import (
	"bufio"
	"errors"
	"flag"
	"io/ioutil"
	"net"
	"os/user"
	"strconv"
)

func main() {
	flag.Parse()
	ln, err := net.Listen("tcp", ":79")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	usr, _, _ := reader.ReadLine()

	if info, err := getUserInfo(string(usr)); err != nil {
		conn.Write([]byte(err.Error()))
	} else {
		conn.Write(info)
	}
}

func getUserInfo(usr string) ([]byte, error) {
	u, e := user.Lookup(usr)
	if e != nil {
		return nil, e
	}
	data, err := ioutil.ReadFile(u.HomeDir + ".plan")
	if err != nil {
		return data, errors.New("User doesn't have a .plan file!\n")
	}
	return data, nil
}
