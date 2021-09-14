package app

import (
	"rest-api-template-go/controller/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
}
