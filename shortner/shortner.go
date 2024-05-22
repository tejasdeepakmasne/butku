// logic for shortening the url
package shortner

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/go-redis/redis/v8"
	"github.com/lithammer/shortuuid/v4"
	redisfunc "github.com/tejasdeepakmasne/butku/redisfunc"
)

type URLData struct {
	URL string `json:"url"`
}

type Response struct {
	ShortURL    string        `json:"shortURL"`
	RedirectURL string        `json:"redirectURL"`
	ExpiryTime  time.Duration `json:"expiryTime"`
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var data URLData
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !govalidator.IsURL(data.URL) {
		http.Error(w, "Not a URL", http.StatusBadRequest)
		return
	}

	rdb := redisfunc.InitializeRDB(0)
	defer rdb.Close()

	var newResponse Response
	newResponse.RedirectURL = data.URL
	newResponse.ShortURL = shortuuid.NewWithNamespace(data.URL)
	newResponse.ExpiryTime = 24 * time.Hour

	if _, err := rdb.Get(redisfunc.Ctx, newResponse.ShortURL).Result(); err != redis.Nil {
		http.Error(w, "URL already shortened", http.StatusForbidden)
		return
	}

	if _, err := rdb.Set(redisfunc.Ctx, newResponse.ShortURL, newResponse.RedirectURL, newResponse.ExpiryTime).Result(); err != nil {
		http.Error(w, "Redis Server connection failed", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
