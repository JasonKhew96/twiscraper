package twiscraper

import (
	"errors"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var ErrorNotLogined = errors.New("cookie and x-csrf-token are required")

type Scraper struct {
	bearerToken    string
	client         *http.Client
	guestToken     string
	guestCreatedAt time.Time
	wg             sync.WaitGroup

	delay      time.Duration
	cookie     string
	xCsrfToken string
	userAgent  string
}

type ScraperOptions struct {
	Delay time.Duration

	Cookie     string
	XCsrfToken string

	Timeout   time.Duration
	Proxy     string
	UserAgent string
}

func (s *Scraper) hasGuestToken() bool {
	return s.guestToken != ""
}

func (s *Scraper) IsLogined() bool {
	return s.cookie != "" && s.xCsrfToken != ""
}

func New(opts ScraperOptions) (*Scraper, error) {
	scraper := Scraper{
		bearerToken: defaultBearerToken,
		client:      &http.Client{Timeout: defaultClientTimeout},
		userAgent:   defaultUserAgent,
	}
	if opts.Delay != 0 {
		scraper.delay = opts.Delay
	}
	if opts.Cookie != "" && opts.XCsrfToken != "" {
		scraper.cookie = opts.Cookie
		scraper.xCsrfToken = opts.XCsrfToken
	}
	if opts.Proxy != "" {
		u, err := url.Parse(opts.Proxy)
		if err != nil {
			return nil, err
		}
		scraper.client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(u),
			},
		}
	}
	if opts.Timeout != 0 {
		scraper.client.Timeout = opts.Timeout
	}
	if opts.UserAgent != "" {
		scraper.userAgent = opts.UserAgent
	}
	return &scraper, nil
}
