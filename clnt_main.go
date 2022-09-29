package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	conn, error := net.Dial("tcp", "127.0.0.1:10000")
	if error != nil {
		log.Println(error.Error())
		os.Exit(0)
	}

	fmt.Println("OK, Connecting to Server Success...")
	fmt.Println("Entered Your Msg to below.")

	go proc_send_msg(conn)
	go proc_recv_msg(conn)
	for {

	}
}
