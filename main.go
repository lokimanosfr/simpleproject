package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/takama/router"
)

// func home(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Processing %s ....", r.URL.Path)
// }
var log = logrus.New()

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("Port is not set")
	}
	r := router.New()
	r.Logger = logger

	r.GET("/", home)
	r.Listen(":" + port)

}
