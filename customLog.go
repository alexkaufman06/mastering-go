package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "mGo.log")
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	fmt.Println(f.Name())

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close() // defer tells Go to execute statement just before function returns
	iLog := log.New(f, "iLog ", log.LstdFlags)
	iLog.Println("Hello there!")
	iLog.Println("Masting Go 3rd edition!")
}