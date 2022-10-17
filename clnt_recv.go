package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

func ResponseHandler(Conn net.Conn, wg *sync.WaitGroup) {

	/*serv_msg := make(map[string]MyMsg)*/
	recv := make([]byte, 4096)

	for ServiceFlag == 1 {
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

			inputdata := MsgInfo{}
			json.Unmarshal([]byte(serv_msg.Body), &inputdata)

			fmt.Printf("(%s) %s\n", inputdata.Id, inputdata.Message)
		}
	}

	fmt.Println("recv_msg goroutine exiting...")
	wg.Done()
}
