// logic for shortening the url
package shortner

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"github.com/lithammer/shortuuid/v4"
	redisfunc "github.com/tejasdeepakmasne/butku/redisfunc"
)

type URLData struct {
	URL string `json:"url"`
}

type Response struct {
	ShortURL         string    `json:"shortURL"`
	CompleteShortURL string    `json:"fullShortUrl"`
	RedirectURL      string    `json:"redirectURL"`
	ExpiryTime       time.Time `json:"expiryTime"`
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error Loading env file %v", err)
	}
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
	newResponse.ExpiryTime = time.Now().Add(24 * time.Hour)
	newResponse.CompleteShortURL = os.Getenv("DOMAIN") + os.Getenv("SERVE_PORT") + "/" + newResponse.ShortURL

	if _, err := rdb.Get(redisfunc.Ctx, newResponse.ShortURL).Result(); err != redis.Nil {
		http.Error(w, "URL already shortened", http.StatusForbidden)
		return
	}

	if _, err := rdb.Set(redisfunc.Ctx, newResponse.ShortURL, newResponse.RedirectURL, 24*time.Hour).Result(); err != nil {
		http.Error(w, "Redis Server connection failed", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newResponse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
