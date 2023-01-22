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

	// splitFunc := computeSplitFunc(buf)
	// a, _ := NewAssocFromScanner(scanner)
	// printStruct(*a)

	a2, _ := NewAssocFromScanner2(scanner)
	printStruct(*a2)

	// a3, _ := NewAssocFromScanner3(scanner)
	// printStruct(*a3)

}

// ComputeSplitFunc accepts a message buffer and returns a split function that will split
// the scanner into tokens of various byte length depending on the message.
func computeSplitFunc(buf []byte) bufio.SplitFunc {
	// determine the type of the message
	switch buf[0] {
	case 0x01:
		// association request
		return computeAssocRQSplitFunc(buf)
	case 0x02:
		// TODO: Other cases
	}
	return nil
}

func computeAssocRQSplitFunc(buf []byte) bufio.SplitFunc {
	// initial byte size sequence of association request
	// initialSequence := []int{1, 1, 4, 2, 2, 16, 16, 32}

	// Get the length of the application context item
	// applicationContextLength := int(binary.BigEndian.Uint16(buf[76:78]))
	// applicationContextSequence := []int{1, 1, 2, applicationContextLength}
	return nil
}
