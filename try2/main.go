package main

import (
	"try2/router"
)

func main() {
	r := router.SetupRoute()
	r.Run(":8000")
}