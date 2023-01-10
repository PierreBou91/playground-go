package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"reflect"
)

type myStruct struct {
	Field1 [205]byte
	Field2 [6]byte
}

type Associate struct {
	PduType         uint8
	Reserved        uint8
	Length          [4]byte
	ProtocolVersion [2]byte
	Reserved2       [2]byte
	CalledAETitle   [16]byte
	CallingAETitle  [16]byte
	Reserved3       [32]byte
	// variableItems   []byte
}

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
	// fmt.Printf("%s", myStruct.Field1)
	// fmt.Printf("%s", myStruct.Field2)
	// fmt.Printf("%+v\n", associate)

	printStruct(associate)

	// fmt.Printf("pduType: %v\n", associate.PduType)
	// fmt.Printf("reserved: %v\n", associate.Reserved)
	// fmt.Printf("length: %v\n", associate.Length)
	// fmt.Printf("protocol version: %v\n", associate.ProtocolVersion)
	// fmt.Printf("reserved2: %v\n", associate.Reserved2)
	// fmt.Printf("calledAET: %v\n", associate.CalledAETitle)
	// fmt.Printf("callingAET: %v\n", associate.CallingAETitle)
	// fmt.Printf("reserved3: %v\n", associate.Reserved3)

}

func printStruct(s interface{}) {
	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		fmt.Printf("%s %v \n", v.Type().Field(i).Name, f.Interface())
	}

}
