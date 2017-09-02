package api

import (
	"net/http"
	"log"

	"github.com/boltdb/bolt"
	"fmt"
)

type Api struct {
	DB *bolt.DB
}

type Content struct {
	http_status int
	body string
}

func NewApi(db *bolt.DB) *Api {
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	return &Api{
		DB: db,
	}
}

func (api *Api) RequestHandler(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	uri := r.RequestURI

	key := method + "|" + uri

	result := ""
	e := api.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte(key))
		if v == nil {
			return fmt.Errorf("Key [%s] not found", key)
		}
		result = string(v)
		return nil
	})
	if e != nil {
		api.DB.Update(func(tx *bolt.Tx) error {
			b, e := tx.CreateBucketIfNotExists([]byte("MyBucket"))
			if e != nil {
				fmt.Printf("%s|", e)
				return e
			}
			return b.Put([]byte(key), []byte(""))
		})
	}

	log.Print(uri)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
	w.Write([]byte(`{"nama":"arman"}`))
}