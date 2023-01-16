package twiscraper

import (
	"fmt"
	"net/http"

	"github.com/JasonKhew96/twiscraper/entity"
)

func (s *Scraper) Follow(screenName string) error {
	if !s.IsLogined() {
		return ErrorNotLogined
	}
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

	var friendships entity.Friendships
	err = s.requestAPI(req, &friendships)
	if err != nil {
		return err
	}
	return nil
}

func (s *Scraper) UnFollow(screenName string) error {
	if !s.IsLogined() {
		return ErrorNotLogined
	}
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

	var friendships entity.Friendships
	err = s.requestAPI(req, &friendships)
	if err != nil {
		return err
	}
	return nil
}
