package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func proc_recv_msg(Conn net.Conn) {

	recv := make([]byte, 4096)

	for {
		n, error := Conn.Read(recv)
		if error != nil {
			if error == io.EOF {
				fmt.Println("Connection is Close by Client\n", Conn.RemoteAddr().String())
			}

			log.Println(error.Error())
			break
		}
		if n > 0 {
			fmt.Printf("%s\n", string(recv[:n]))
		}
	}
}
