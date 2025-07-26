package main 

import (
	"try1/router"
)

func main() {
	r := router.SetupRoute()
	r.Run(":8000")
}