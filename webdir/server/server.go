package server

import (
	"Banosa_Project/webdir/route"
)

// Run : Server Start
func Run() {
	e := route.Route()
	e.Logger.Fatal(e.Start(":1323"))
}
