package main

import "net"

type wrappedFunc func(net.Conn)

func tcpWrapper(f wrappedFunc) {

	ln, err := net.Listen("tcp", ":11112")

	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			panic(err)
		}
		go f(conn)
	}
}
