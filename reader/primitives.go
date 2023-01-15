package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
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
	VariableItems   []byte
}

type PDUSequence []int

var (
	AssociateRQ PDUSequence = []int{1, 1, 4, 2, 2, 16, 16, 32}
	AppContext              = []int{1, 1, 2}
	AssociateAC             = []int{1, 1, 4, 2, 2, 16, 16, 32}
	AssociateRJ             = []int{1, 1, 4, 2, 2, 16, 16, 32}
	DataTF                  = []int{1, 1, 4}
)

type Associate2 struct {
	PduType         []byte
	Reserved        []byte
	Length          []byte
	ProtocolVersion []byte
	Reserved2       []byte
	CalledAETitle   []byte
	CallingAETitle  []byte
	Reserved3       []byte
	VariableItems   VariableItems
}

type VariableItems struct {
	ApplicationContext ApplicationContext
	// PresentationContextList []PresentationContext
	// UserInfo                UserInfo
}

// type VariableItems struct {
// 	ApplicationContext []byte
// }

type ApplicationContext struct {
	ItemType               []byte
	Reserved               []byte
	Length                 []byte
	ApplicationContextName []byte
}

type PresentationContext struct {
	ItemType              []byte
	Reserved              []byte
	Length                []byte
	PresentationContextID []byte
	Reserved2             []byte
	ResultReason          []byte
	Reserved3             []byte
	AbstractSyntax        AbstractSyntax
	TransferSyntaxList    []TransferSyntax
}

type AbstractSyntax struct {
	ItemType           []byte
	Reserved           []byte
	Length             []byte
	AbstractSyntaxName []byte
}

type TransferSyntax struct {
	ItemType           []byte
	Reserved           []byte
	Length             []byte
	TransferSyntaxName []byte
}

type UserInfo struct {
	ItemType            []byte
	Reserved            []byte
	Length              []byte
	UserInfoSubItemList []UserInfoSubItem
}

type UserInfoSubItem struct {
	ItemType  []byte
	Reserved  []byte
	Length    []byte
	MaxLength []byte
}

// NewAssocFromScanner2 creates a new Associate2 struct from a bufio.Scanner
// this function implements its specific SequenceSplitFunc
func NewAssocFromScanner2(scanner *bufio.Scanner) (*Associate2, error) {
	assocSplitFunc := SequenceSplitFunc(AssociateRQ) // accomodate a switch for RQ, AC, RJ
	scanner.Split(assocSplitFunc)
	a := &Associate2{}
	val := reflect.ValueOf(a).Elem()
	for i := 0; i < val.NumField(); i++ {
		// check if field is VariableItems
		if val.Type().Field(i).Name == "VariableItems" {
			varItems, _ := NewVariableItemsFromScanner(scanner)
			val.Field(i).Set(reflect.ValueOf(*varItems))
			continue
		}
		scanner.Scan()
		val.Field(i).SetBytes(scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return a, nil
}

//	func NewVariableItemsFromScanner(scanner *bufio.Scanner) (*VariableItems, error) {
//		v := &VariableItems{}
//		scanner.Scan()
//		v.ApplicationContext = scanner.Bytes()
//		return v, nil
//	}
func NewVariableItemsFromScanner(scanner *bufio.Scanner) (*VariableItems, error) {
	v := &VariableItems{}
	appContext, _ := NewApplicationContextFromScanner(scanner)
	v.ApplicationContext = *appContext
	return v, nil
}

func NewApplicationContextFromScanner(scanner *bufio.Scanner) (*ApplicationContext, error) {
	scanner.Scan()                                                           // necessary to get the length of the app context
	appContextLength := int(binary.BigEndian.Uint16((scanner.Bytes()[2:4]))) // Lenght is necessary to give the correct sequence to the SequenceSplitFunc
	scanner, _ = NewScannerFromScanner(scanner)
	sequence := append(AppContext, appContextLength)
	fmt.Println(sequence)
	appContextSplitFunc := SequenceSplitFunc(sequence)
	scanner.Split(appContextSplitFunc)
	a := &ApplicationContext{}
	val := reflect.ValueOf(a).Elem()
	for i := 0; i < val.NumField(); i++ {
		scanner.Scan()
		val.Field(i).SetBytes(scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return a, nil
}

type Associate3 struct {
	ItemKeys []string
	Items    map[string][]byte
}

func NewAssocFromScanner3(scanner *bufio.Scanner) (*Associate3, error) {
	a := &Associate3{}
	a.ItemKeys = []string{"PduType", "Reserved", "Length", "ProtocolVersion", "Reserved2", "CalledAETitle", "CallingAETitle", "Reserved3", "VariableItems"}
	a.Items = make(map[string][]byte)
	for _, v := range a.ItemKeys {
		scanner.Scan()
		a.Items[v] = scanner.Bytes()
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return a, nil
}

// func NewAssocFromScanner(scanner *bufio.Scanner) (*Associate2, error) {
// 	var a Associate2
// 	scanner.Scan()
// 	a.PduType = scanner.Bytes()
// 	scanner.Scan()
// 	a.Reserved = scanner.Bytes()
// 	scanner.Scan()
// 	a.Length = scanner.Bytes()
// 	scanner.Scan()
// 	a.ProtocolVersion = scanner.Bytes()
// 	scanner.Scan()
// 	a.Reserved2 = scanner.Bytes()
// 	scanner.Scan()
// 	a.CalledAETitle = scanner.Bytes()
// 	scanner.Scan()
// 	a.CallingAETitle = scanner.Bytes()
// 	scanner.Scan()
// 	a.Reserved3 = scanner.Bytes()
// 	scanner.Scan()
// 	a.VariableItems = scanner.Bytes()
// 	return &a, nil
// }
