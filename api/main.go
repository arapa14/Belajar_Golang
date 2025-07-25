package main 

import (
	"api/router"
)

func main() {
	r := router.SetupRoute()
	r.Run(":8000")
}