package twiscraper

import (
	"compress/flate"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/andybalholm/brotli"
)

func (s *Scraper) requestAPI(req *http.Request, target interface{}) error {
	s.wg.Wait()
	if s.delay > 0 {
		defer func() {
			s.wg.Add(1)
			go func() {
				time.Sleep(s.delay)
				s.wg.Done()
			}()
		}()
	}

	if !s.hasGuestToken() || s.guestCreatedAt.Before(time.Now().Add(-time.Hour*3)) {
		err := s.refreshGuestToken()
		if err != nil {
			return err
		}
	}

	req.Header.Set("User-Agent", s.userAgent)
	req.Header.Set("Authorization", "Bearer "+s.bearerToken)
	req.Header.Set("X-Guest-Token", s.guestToken)
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")

	// use cookie
	if s.IsLogined() {
		req.Header.Set("Cookie", s.cookie)
		req.Header.Set("x-csrf-token", s.xCsrfToken)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.Header.Get("x-rate-limit-remaining") == "0" || resp.StatusCode == http.StatusTooManyRequests {
		s.guestToken = ""
		reset, err := strconv.ParseInt(resp.Header.Get("x-rate-limit-reset"), 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse x-rate-limit-reset: %v", err)
		}
		resetTime := time.Unix(reset, 0)
		until := time.Until(resetTime)
		s.sugar.Infof("rate limit exceeded, reset at %s, sleep for %.2f", resetTime, until.Seconds())
		time.Sleep(until)
		return s.requestAPI(req, target)
	}

	var content []byte
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		gr, err := gzip.NewReader(resp.Body)
		if err != nil {
			return err
		}
		defer gr.Close()
		content, err = io.ReadAll(gr)
		if err != nil {
			return err
		}
	case "deflate":
		fr := flate.NewReader(resp.Body)
		defer fr.Close()
		content, err = io.ReadAll(fr)
		if err != nil {
			return err
		}
	case "br":
		br := brotli.NewReader(resp.Body)
		content, err = io.ReadAll(br)
		if err != nil {
			return err
		}
	default:
		content, err = io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusForbidden && resp.StatusCode != http.StatusTooManyRequests {
		return fmt.Errorf("request api failed %s: %s", resp.Status, content)
	}

	s.sugar.Debugf("request api %s\n%s", req.URL, content)
	return json.Unmarshal(content, target)
}

func (s *Scraper) refreshGuestToken() error {
	req, err := http.NewRequest(http.MethodPost, "https://api.x.com/1.1/guest/activate.json", nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", s.userAgent)
	req.Header.Set("Authorization", "Bearer "+s.bearerToken)

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("refresh guest token failed %s: %s", resp.Status, body)
	}

	var jsn map[string]interface{}
	if err := json.Unmarshal(body, &jsn); err != nil {
		return err
	}
	var ok bool
	if s.guestToken, ok = jsn["guest_token"].(string); !ok {
		return fmt.Errorf("guest_token not found")
	}
	s.guestCreatedAt = time.Now()

	return nil
}
