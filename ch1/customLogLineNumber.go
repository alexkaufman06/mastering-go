package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "mGo.log")
	fmt.Println("Path: ", LOGFILE)
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	LstdFlags := log.Ldate | log.Lshortfile
	iLog := log.New(f, "LNum ", LstdFlags)
	iLog.Println("Masting Go, 3rd edition!")
	iLog.SetFlags(LstdFlags)
	iLog.Println("Another log entry!")
}
