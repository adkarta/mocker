package main

import (
	"log"
	"net/http"

	"mocker/api"
	"github.com/husobee/vestigo"

	"github.com/boltdb/bolt"
	"time"
)

func main() {


	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

	api := api.NewApi(db)

	router := vestigo.NewRouter()
	vestigo.AllowTrace = true

	router.Get("/*", api.RequestHandler)
	router.Post("/*", api.RequestHandler)

	log.Print("API Start listen port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))

}
