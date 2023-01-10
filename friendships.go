package twiscraper

import (
	"fmt"
	"net/http"
)

func (s *Scraper) Follow(screenName string) error {
	parsedUser, err := s.GetUserByScreenName(screenName)
	if err != nil {
		return err
	}
	if parsedUser.IsFollowing {
		return fmt.Errorf("already following %s", screenName)
	}

	req, err := http.NewRequest(http.MethodPost, apiFriendShipsCreate, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("user_id", parsedUser.UserId)
	req.URL.RawQuery = q.Encode()

	err = s.requestAPI(req, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *Scraper) Unfollow(screenName string) error {
	parsedUser, err := s.GetUserByScreenName(screenName)
	if err != nil {
		return err
	}
	if !parsedUser.IsFollowing {
		return fmt.Errorf("not following %s", screenName)
	}

	req, err := http.NewRequest(http.MethodPost, apiFriendshipsDestroy, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("user_id", parsedUser.UserId)
	req.URL.RawQuery = q.Encode()

	err = s.requestAPI(req, nil)
	if err != nil {
		return err
	}
	return nil
}
