package main

import (
	"Banosa_Project/webdir/server"
	"fmt"
	"log"
	"os"
)

var Logger *log.Logger

func main() {
	log.Println("Logging")
	log.SetFlags(0)
	fpLog, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fpLog.Close()

	Logger = log.New(fpLog, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	run()

	Logger.Println("End of Program")
	fmt.Println("START Banosa")
	server.Run()
}

func run() {
	Logger.Print("Test")
}
