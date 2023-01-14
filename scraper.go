package twiscraper

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"go.uber.org/zap"
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

	sugar *zap.SugaredLogger
	debug bool
}

type ScraperOptions struct {
	Delay time.Duration

	Cookie     string
	XCsrfToken string

	Timeout   time.Duration
	Proxy     string
	UserAgent string

	Debug bool
}

func (s *Scraper) hasGuestToken() bool {
	return s.guestToken != ""
}

func (s *Scraper) IsLogined() bool {
	return s.cookie != "" && s.xCsrfToken != ""
}

func New(opts *ScraperOptions) (*Scraper, error) {
	scraper := Scraper{
		bearerToken: defaultBearerToken,
		delay:       defaultDelay,
		client:      &http.Client{Timeout: defaultClientTimeout},
		userAgent:   defaultUserAgent,
	}
	if opts != nil {
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
				return nil, fmt.Errorf("invalid proxy url: %v", err)
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
		if opts.Debug {
			scraper.debug = opts.Debug
		}
	}

	var logger *zap.Logger
	var err error
	if opts != nil && opts.Debug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %v", err)
	}
	scraper.sugar = logger.Sugar()
	scraper.sugar.Infoln("Scraper initialized")

	return &scraper, nil
}
