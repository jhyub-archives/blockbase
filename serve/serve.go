package serve

import (
	"fmt"
	"net"
	"strings"
	"os"
	"github.com/blbase/base"
)

func Serve(adr string) {
	base.Create("1")
	fmt.Printf("Check Database Folder...\n")
	if _, err := os.Stat("database"); os.IsNotExist(err) {
		fmt.Printf("Database Folder doesn't exist.\n")
		os.MkdirAll("database", 0777)
		fmt.Printf("Created Database Folder.\n")
	} else {
		fmt.Printf("Database Folder exists.\n")
	}
	listener, err := net.Listen("tcp", adr)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("The BlockBase Server has started to %s", adr)
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
			msg := ""
			conn.Read(buffer)
			for _, data := range buffer {
				if data != 0 {
					msg = msg + string(data)
				}
			}
			if err == nil {
				// to be proceed
				a := strings.Split(msg, " ")
				print(a)
			}
		}()
	}

}
