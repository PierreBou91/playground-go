package main

import (
	"bufio"
	"bytes"
	"net"
)

func MyScanner(conn net.Conn) {
	defer conn.Close()

	// read from connection
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)

	if err != nil {
		panic(err)
	}

	buf = trimMessage(buf) // TODO: Evaluate the need for this

	reader := bufio.NewReader(bytes.NewReader(buf[:n]))
	scanner := bufio.NewScanner(reader)
	// dicomSplitFunc := SequenceSplitFunc([]int{1, 1, 4, 2, 2, 16, 16, 32, int(binary.BigEndian.Uint32(buf[2:6])) - 68})

	// a, _ := NewAssocFromScanner(scanner)
	// printStruct(*a)

	a2, _ := NewAssocFromScanner2(scanner)
	printStruct(*a2)

	// a3, _ := NewAssocFromScanner3(scanner)
	// printStruct(*a3)

}
