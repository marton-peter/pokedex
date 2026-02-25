package pokeapi

import (
	"net/http"
	"time"

	"github.com/marton-peter/pokedex/internal/pokecache"
)

type Client struct {
	client http.Client
	cache  *pokecache.Cache
}

func NewClient(timeout time.Duration, interval time.Duration) *Client {
	c := Client{client: http.Client{Timeout: timeout}, cache: pokecache.NewCache(interval)}
	return &c
}
