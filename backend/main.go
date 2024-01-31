package main

import (
	"Syojincoder2/router"
)

func main(){
	r := router.SetupRouter()
	r.Run(":8080")
}