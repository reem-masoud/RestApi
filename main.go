package main

import (
	"fmt"

	"github.com/restapi/routes"
)

//"strconv"

func main() {
	routes.init()

	fmt.Println("Running...")
	routes.Router()

	e.Logger.Fatal(e.Start(":1221"))

}
