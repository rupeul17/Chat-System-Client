package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func proc_recv_msg(Conn net.Conn) {

	/*serv_msg := make(map[string]MyMsg)*/
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

			serv_msg := DecodeToMyMsg(recv)

			fmt.Printf("(%s) >> %s\n", serv_msg.Head.Ip, serv_msg.Body)
			//fmt.Println(serv_msg["Ip"])
		}
	}
}
