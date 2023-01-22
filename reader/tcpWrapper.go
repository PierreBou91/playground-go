package main

import "net"

type wrappedFunc func(net.Conn)

// Base TCP connection listener that accepts a connection handler function.
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
