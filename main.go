package main

import (
	"example/web-service-gin/di"

	_ "github.com/lib/pq"
)

func main() {
	app, err := di.InitializeEvent()
	if err != nil {
		panic(err)
	}

	app.Run(":8080")
}
