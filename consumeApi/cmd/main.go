package main

import "consume-api/internal/routes"

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}
