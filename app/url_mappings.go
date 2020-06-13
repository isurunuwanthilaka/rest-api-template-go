package app

import "github.com/isurunuwanthilaka/rest-api-template/controller"

func mapUrls() {
	router.GET("/ping", controller.Ping)
}
