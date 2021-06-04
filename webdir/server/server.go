package server

import (
	"Banosa_Project/webdir/route"
	"log"
)

// Run : Server Start
// 21.01.11 port 1323 to 80
func Run() {
	e := route.Route()
	e.Logger.Fatal(e.Start(":8848"))
	log.Println("TI")
}
