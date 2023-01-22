package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

func printStruct(s interface{}) {
	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		fmt.Printf("%s %v \n", v.Type().Field(i).Name, f.Interface())
	}

}

func trimMessage(message []byte) []byte {
	return message[:6+int(binary.BigEndian.Uint32(message[2:6]))]
}

// SequenceSplitFunc accepts a sequence and returns a split function that will split
// the buffer into tokens of various byte length depending on the sequence.
// This allows to split the PDU item according to its type.
func SequenceSplitFunc(sequence []int) bufio.SplitFunc {
	cursor := 0 // keeps track of the index in the sequence

	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// base condition to stop the scan
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		// Stop the scan if the sequence is over, effectively returning the last
		// bytes of the buffer as a token.
		// This allows for variable length items to be dealt with later.
		if cursor >= len(sequence) {
			return 0, data, nil
		}

		advance = sequence[cursor]
		token = data[0:advance]
		cursor++
		return advance, token, nil
	}
}

func NewScannerFromScanner(scanner *bufio.Scanner) (*bufio.Scanner, error) {
	// exceptionnally we don't need a scanner.Scan() here because the scanner
	// is only one token long, so we can just read the bytes and create a new
	// scanner from it.
	buf := scanner.Bytes()
	scanner = bufio.NewScanner(bufio.NewReader(bytes.NewReader(buf[:])))
	return scanner, nil
}
