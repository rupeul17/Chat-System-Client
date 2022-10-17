package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func SendMessageToServer(conn net.Conn, wg *sync.WaitGroup) {

	for ServiceFlag == 1 {
		fmt.Printf(">> ")
		send_msg := input_string()

		msginfo := MsgInfo{
			Id:      LoginInfo.Id,
			Message: []byte(send_msg),
		}

		jsondata, _ := json.Marshal(msginfo)

		msg := MyMsg{
			Head: Header{
				MsgType: TYPE_MESSAGE,
				BodyLen: len(jsondata),
			},
			Body: jsondata,
		}

		bytedata := EncodeToBytes(msg)

		_, error := conn.Write(bytedata)
		if error != nil {
			log.Println(error.Error())
		}

		time.Sleep(1 * time.Second)
	}

	fmt.Println("send_msg goroutine exiting...")
	wg.Done()
}
