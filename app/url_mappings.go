package app

import (
	"github.com/isurunuwanthilaka/rest-api-template/controller/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
}
