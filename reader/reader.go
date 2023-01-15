package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

func Reader(conn net.Conn) {
	defer conn.Close()
	// read from connection
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)

	if err != nil {
		panic(err)
	}

	var myStruct myStruct
	var associate Associate

	err = binary.Read(bytes.NewReader(buf[:n]), binary.BigEndian, &myStruct)

	if err != nil {
		panic(err)
	}

	err = binary.Read(bytes.NewReader(buf[:n]), binary.BigEndian, &associate)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d\n", buf[:n])

	printStruct(associate)

}
