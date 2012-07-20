package main

import (
	"bufio"
	"io/ioutil"
	"errors"
	"net"
	"os/user"
	"flag"
	"strconv"
)

var port *int = flag.Int("port", 79, "Port to listen on")

func main() {
	flag.Parse()
	ln, err := net.Listen("tcp", ":"+ strconv.Itoa(*port))
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

	info, err := getUserInfo(string(usr))
	if err != nil {
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
