package admin

import (
	"log"
	"net/http"

	"github.com/husobee/vestigo"
)

func Admin() {
	router := vestigo.NewRouter()
	vestigo.AllowTrace = true

	router.Get("/*", RequestHandler)
	router.Post("/*", RequestHandler)

	log.Print("Web Admin Start listen port 1235")
	log.Fatal(http.ListenAndServe(":1235", router))

}


func RequestHandler(w http.ResponseWriter, r *http.Request) {
	uri := r.RequestURI
	log.Print(uri)
	w.WriteHeader(200)
	w.Write([]byte("WEb Admin Gotta catch em all!"))
}
