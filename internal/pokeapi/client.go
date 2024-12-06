package pokeapi

import (
	"net/http"
	"time"

	"github.com/bootdotdev/go-api-gate/courses/projects/bootdev_pokedex/pokecache"
)

// Client
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// New Client
func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
