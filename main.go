package main

import (
	"testgo/src/app"
)

const configPath = "config/config.yaml"

func main() {
	app.Run(configPath)
}
