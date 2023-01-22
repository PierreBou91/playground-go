package main

import (
	"log"
	"net/rpc"
	"rpc_structs"
)

func main() {
	// Connect to the RPC server.
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing error:", err)
	}

	// Send an RPC request and receive a response.
	args := &rpc_structs.Args{A: 2, B: 3}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}

	// Print the response.
	log.Printf("Result: %d", reply)
}
