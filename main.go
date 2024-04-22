package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const INDEX = `<!DOCTYPE html>
<html>
  <head>
    <title>Powered By Paketo Buildpacks</title>
  </head>
  <body>
    <img style="display: block; margin-left: auto; margin-right: auto; width: 50%;" src="https://paketo.io/images/paketo-logo-full-color.png"></img>
  </body>
</html>`

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		for _, e := range os.Environ() {
			log.Println(e)
		}

		fmt.Fprintln(w, INDEX)
	})
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), loggedRouter))
}
