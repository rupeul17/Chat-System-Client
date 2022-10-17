package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func DecodeToMyMsg(s []byte) MyMsg {

	myMsg := MyMsg{}
	dec := gob.NewDecoder(bytes.NewReader(s))

	err := dec.Decode(&myMsg)
	if err != nil {
		log.Fatal(err)
	}

	return myMsg
}

func EncodeToBytes(p interface{}) []byte {

	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("uncompressed size (bytes): ", len(buf.Bytes()))
	return buf.Bytes()
}

func input_string() string {

	rd := bufio.NewReader(os.Stdin)

	TmpStr, err := rd.ReadString('\n')
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	TmpStr = strings.TrimSpace(TmpStr)

	return TmpStr
}

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ClearTerminal() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}
