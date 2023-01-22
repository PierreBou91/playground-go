package main

import (
	"log"
	"net"
	"net/rpc"

	"rpc_structs"
)

// The RPC service we are providing.
type Arith int

// The RPC method we are exposing.
func (t *Arith) Multiply(args *rpc_structs.Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func main() {
	// Create a new RPC server.
	arith := new(Arith)
	server := rpc.NewServer()
	server.Register(arith)

	// Start listening on a port.
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	// Accept incoming connections and handle them.
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		go server.ServeConn(conn)
	}
}
