package server

import (
	"fmt"
	"net"
)

func Server(adr string) {
	listener, err := net.Listen("tcp", adr)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("The BlockBase Server has started to %v", adr)
	}

	for {
		conn, err := listener.Accept()
		if nil != err {
			fmt.Printf("\nConnection Error: %v", err)
			continue
		} else {
			fmt.Printf("\nConnected from: %v\n", conn.RemoteAddr())
		}
		go func() {
			buffer := make([]byte, 3000)
			stringA := ""
			conn.Read(buffer)
			for _, data := range buffer {
				if data != 0 {
					stringA = stringA + string(data)
				}
			}
			if err == nil {
				// to be proceed
				fmt.Print(stringA)
			}
		}()
	}

}
