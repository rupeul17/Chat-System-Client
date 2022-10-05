package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/*
	global variable
*/
var ServiceFlag int

func main() {

	ServiceFlag = 1

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	wg := sync.WaitGroup{}
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	conn, error := net.Dial("tcp", "127.0.0.1:10000")
	if error != nil {
		log.Println(error.Error())
		os.Exit(0)
	}

	fmt.Println("OK, Connecting to Server Success...")

	if proc_login(conn) < 0 {
		log.Println("login failed...")
		os.Exit(0)
	}

	fmt.Println("Entered Your Msg to below.")

	wg.Add(4)
	go proc_send_msg(conn, &wg)
	go proc_recv_msg(conn, &wg)
	go func() {
		/*
			signal check
		*/
		sig := <-sigs
		fmt.Println(sig)
		ServiceFlag = 0
		done <- true
		wg.Done()
	}()

	<-done
	wg.Wait()
	fmt.Println("Main goroutine exiting...")
}
