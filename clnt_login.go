package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func proc_recv_login_resopnse(conn net.Conn) int {

	recv := make([]byte, 4096)

	for {
		n, error := conn.Read(recv)
		if error != nil {
			if error == io.EOF {
				fmt.Println("Connection is Close by Client\n", conn.RemoteAddr().String())
			}

			log.Println(error.Error())
			break
		}
		if n > 0 {

			serv_msg := DecodeToMyMsg(recv)

			if serv_msg.Head.res == 200 {
				return 1
			} else {
				fmt.Printf("Login failed..., msgtype(%d), ip(%s), res(%d) ", serv_msg.Head.MsgType, serv_msg.Head.Ip, serv_msg.Head.res)
				return -1
			}
		}
	}

	return -1
}

func proc_login(conn net.Conn) int {

	var LoginInfo Login

	fmt.Printf("Entered Your ID : ")
	LoginInfo.Id = input_string()

	fmt.Printf("Entered Your Passwd : ")
	LoginInfo.Pwd = input_string()

	send_msg := EncodeToBytes(LoginInfo)

	msg := MyMsg{
		Head: Header{
			MsgType: 2,
			Ip:      conn.LocalAddr().String(),
			BodyLen: len(send_msg),
		},
		Body: []byte(send_msg),
	}

	bytedata := EncodeToBytes(msg)

	_, error := conn.Write(bytedata)
	if error != nil {
		log.Println(error.Error())
		return -1
	}

	if proc_recv_login_resopnse(conn) < 0 {
		return -1
	} else {
		return 1
	}
}
