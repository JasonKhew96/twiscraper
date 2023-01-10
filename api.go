package twiscraper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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

	req.Header.Set("Authorization", "Bearer "+s.bearerToken)
	req.Header.Set("X-Guest-Token", s.guestToken)

	// use cookie
	if len(s.cookie) > 0 && len(s.xCsrfToken) > 0 {
		req.Header.Set("Cookie", s.cookie)
		req.Header.Set("x-csrf-token", s.xCsrfToken)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusForbidden {
		content, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("request api failed %s: %s", resp.Status, content)
	}

	if resp.Header.Get("X-Rate-Limit-Remaining") == "0" {
		s.guestToken = ""
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

func (s *Scraper) refreshGuestToken() error {
	req, err := http.NewRequest(http.MethodPost, "https://api.twitter.com/1.1/guest/activate.json", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+s.bearerToken)

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
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
