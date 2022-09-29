package main

import (
	"log"
	"net"
	"time"
)

func proc_send_msg(conn net.Conn) {

	for {
		send := input_string()
		_, error := conn.Write([]byte(send))
		if error != nil {
			log.Println(error.Error())
		}

		time.Sleep(1 * time.Second)
	}
}
