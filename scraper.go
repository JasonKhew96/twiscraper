package twiscraper

import (
	"net/http"
	"sync"
	"time"
)

type Scraper struct {
	bearerToken    string
	client         *http.Client
	delay          time.Duration
	guestToken     string
	guestCreatedAt time.Time
	wg             sync.WaitGroup

	cookie     string
	xCsrfToken string
}

const DefaultClientTimeout = 10 * time.Second

var defaultScraper *Scraper

func New() *Scraper {
	return &Scraper{
		bearerToken: DEFAULT_BEARER_TOKEN,
		client:      &http.Client{Timeout: DefaultClientTimeout},
	}
}

func (s *Scraper) SetBearerToken(bearerToken string) {
	s.bearerToken = bearerToken
	s.guestToken = ""
}

func (s *Scraper) HasGuestToken() bool {
	return s.guestToken != ""
}

func (s *Scraper) WithDelay(delay time.Duration) *Scraper {
	s.delay = delay
	return s
}

func (s *Scraper) WithTimeout(timeout time.Duration) *Scraper {
	s.client.Timeout = timeout
	return s
}

func (s *Scraper) WithCookie(cookie string) *Scraper {
	s.cookie = cookie
	return s
}

func (s *Scraper) WithXCsrfToken(token string) *Scraper {
	s.xCsrfToken = token
	return s
}

func init() {
	defaultScraper = New()
}
