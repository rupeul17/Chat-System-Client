package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
)

func ReceiveLoginResponseFromServer(conn net.Conn) int {

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

			if serv_msg.Head.Res == OK {
				return 1
			} else {
				fmt.Printf("Login failed...\n")
				return -1
			}
		}
	}

	return -1
}

func TryLogin(conn net.Conn) int {

	fmt.Printf("Entered Your ID : ")
	LoginInfo.Id = input_string()

	fmt.Printf("Entered Your Passwd : ")
	LoginInfo.Pwd = input_string()

	jsondata, _ := json.Marshal(LoginInfo)

	msg := MyMsg{
		Head: Header{
			MsgType: TYPE_LOGIN,
			BodyLen: len(jsondata),
		},
		Body: jsondata,
	}

	bytedata := EncodeToBytes(msg)

	_, error := conn.Write(bytedata)
	if error != nil {
		log.Println(error.Error())
		return -1
	}

	if ReceiveLoginResponseFromServer(conn) < 0 {
		return -1
	} else {
		return 1
	}
}
