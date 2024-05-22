package resolver

import (
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	redisfunc "github.com/tejasdeepakmasne/butku/redisfunc"
)

func ResolveURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortUrl := vars["shortURL"]

	rdb := redisfunc.InitializeRDB(0)
	defer rdb.Close()

	val, err := rdb.Get(redisfunc.Ctx, shortUrl).Result()
	if err == redis.Nil {
		http.Error(w, "shortURL not found in redis DB", http.StatusNotFound)
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, val, http.StatusMovedPermanently)

}
