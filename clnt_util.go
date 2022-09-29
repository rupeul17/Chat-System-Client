package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

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
