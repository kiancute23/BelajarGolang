package main

import "golang/router"

func main() {
	var PORT = ":8080"

	router.StarServer().Run(PORT)
}
