package main

import (
	"errors"
	"fmt"
	"log"
	"net"
)

func server_with_panic() {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal()
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		go ClientHandler(conn)
	}
}

func ClientHandler(c net.Conn) {
	// для корректной обработки паники
	defer func() {
		if v := recover(); v != nil {
			fmt.Println("captured panic", v)
		}
		c.Close()
	}()
	panic(errors.New("err internal err"))
}
