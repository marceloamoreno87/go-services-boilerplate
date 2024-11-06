package main

import (
	"sendzap-checkout/common/core"
	route "sendzap-checkout/services/checkout"
)

func main() {
	app := core.Application{}

	err := app.Run(&route.Router{})
	if err != nil {
		panic(err)
	}
}
