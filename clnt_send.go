package main

import (
	"log"
	"net"
	"time"
)

func proc_send_msg(conn net.Conn) {

	for {
		send_msg := input_string()

		msg := MyMsg{
			Head: Header{
				MsgType: 1,
				Ip:      conn.LocalAddr().String(),
				BodyLen: len(send_msg),
			},
			Body: []byte(send_msg),
		}

		bytedata := EncodeToBytes(msg)

		_, error := conn.Write(bytedata)
		if error != nil {
			log.Println(error.Error())
		}

		time.Sleep(1 * time.Second)
	}
}
