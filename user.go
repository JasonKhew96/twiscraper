package twiscraper

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"

	"github.com/JasonKhew96/twiscraper/entity"
)

var cachedUserIds sync.Map

// FollowersResult is a struct that contains user information and error
type FollowersResult struct {
	ParsedUser entity.ParsedUser
	Error      error
}

// GetUserByScreenName returns user information by screen name
func (s *Scraper) GetUserByScreenName(screenName string) (*entity.ParsedUser, error) {
	vl, err := entity.NewUserByScreenNameParams(screenName)
	if err != nil {
		return nil, fmt.Errorf("failed to create params: %v", err)
	}
	apiUrl, err := url.Parse(apiUserByScreenName)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %v", err)
	}
	apiUrl.RawQuery = vl.Encode()
	req, err := http.NewRequest(http.MethodGet, apiUrl.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	var user entity.UserByScreenName
	err = s.requestAPI(req, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to request api: %v", err)
	}
	// if len(user.Errors) > 0 {
	// 	return nil, errors.New(user.Errors[0].Message)
	// }

	if user.Data.User.Result.TypeName == "" {
		return nil, fmt.Errorf("user %s not found", screenName)
	}

	if user.Data.User.Result.TypeName == "UserUnavailable" {
		return nil, fmt.Errorf("user %s %s", screenName, user.Data.User.Result.Reason)
	}

	return user.Data.User.Result.Parse()
}

func (s *Scraper) GetUserIdByScreenName(screenName string) (string, error) {
	id, ok := cachedUserIds.Load(screenName)
	if ok {
		return id.(string), nil
	}

	user, err := s.GetUserByScreenName(screenName)
	if err != nil {
		return "", fmt.Errorf("failed to get user by screen name: %v", err)
	}

	cachedUserIds.Store(screenName, user.UserId)

	return user.UserId, nil
}

// GetFollowing returns a channel of following
func (s *Scraper) GetFollowing(ctx context.Context, screenName string, count int) <-chan *FollowersResult {
	return s.getFollowersStream(ctx, FetchFollowing, screenName, count, s.fetchFollowers)
}

// GetFollowers returns a channel of followers
func (s *Scraper) GetFollowers(ctx context.Context, screenName string, count int) <-chan *FollowersResult {
	return s.getFollowersStream(ctx, FetchFollowers, screenName, count, s.fetchFollowers)
}

func (s *Scraper) getFollowersStream(ctx context.Context, opt fetchOptions, screenName string, count int, fetchFunc fetchFollowersFunc) <-chan *FollowersResult {
	ch := make(chan *FollowersResult)
	go func(screenName string) {
		defer close(ch)

		id, err := s.GetUserIdByScreenName(screenName)
		if err != nil {
			ch <- &FollowersResult{Error: err}
			return
		}

		var nextCursor string
		followersCount := 0
		for followersCount < count {
			select {
			case <-ctx.Done():
				ch <- &FollowersResult{Error: ctx.Err()}
				return
			default:
			}

			userResults, next, err := fetchFunc(opt, id, count, nextCursor)
			if err != nil {
				ch <- &FollowersResult{Error: err}
				return
			}

			if len(userResults) == 0 {
				break
			}

			for _, user := range userResults {
				select {
				case <-ctx.Done():
					ch <- &FollowersResult{Error: ctx.Err()}
					return
				default:
				}

				if followersCount < count {
					nextCursor = next
					ch <- &FollowersResult{ParsedUser: user}
				} else {
					break
				}
				followersCount++
			}
		}
	}(screenName)
	return ch
}

func (s *Scraper) fetchFollowers(opt fetchOptions, userId string, count int, cursor string) ([]entity.ParsedUser, string, error) {
	if !s.IsLogined() {
		return nil, "", ErrorNotLogined
	}

	if count > 20 {
		count = 20
	}

	vl, err := entity.NewFollowersParams(userId, count, cursor)
	if err != nil {
		return nil, "", fmt.Errorf("failed to create params: %v", err)
	}
	var u string
	switch opt {
	case FetchFollowing:
		u = apiFollowing
	case FetchFollowers:
		u = apiFollowers
	}
	apiUrl, err := url.Parse(u)
	if err != nil {
		return nil, "", fmt.Errorf("failed to parse url: %v", err)
	}
	apiUrl.RawQuery = vl.Encode()
	req, err := http.NewRequest(http.MethodGet, apiUrl.String(), nil)
	if err != nil {
		return nil, "", fmt.Errorf("failed to create request: %v", err)
	}
	var followers entity.FollowersResponse
	err = s.requestAPI(req, &followers)
	if err != nil {
		return nil, "", fmt.Errorf("failed to request api: %v", err)
	}
	// if len(followers.Errors) > 0 {
	// 	return nil, "", errors.New(followers.Errors[0].Message)
	// }

	var userResults []entity.ParsedUser
	var nextCursor string

	for _, instructionRaw := range followers.Data.User.Result.Timeline.Timeline.Instructions {
		var instruction entity.Instruction
		err := json.Unmarshal(instructionRaw, &instruction)
		if err != nil {
			s.sugar.Errorln(err)
			continue
		}
		switch instruction.Type {
		case "TimelineClearCache":
			continue
		case "TimelineTerminateTimeline":
			continue
		case "TimelineAddEntries":
			var timelineAddEntries entity.TimelineInstructionAddEntries
			err := json.Unmarshal(instructionRaw, &timelineAddEntries)
			if err != nil {
				s.sugar.Errorln(err)
				continue
			}
			for _, entryRaw := range timelineAddEntries.Entries {
				var entry entity.Entry
				err := json.Unmarshal(entryRaw, &entry)
				if err != nil {
					s.sugar.Errorln(err)
					continue
				}
				switch entry.Content.EntryType {
				case "TimelineTimelineItem":
					var userResultEntry entity.UserResultEntry
					err := json.Unmarshal(entryRaw, &userResultEntry)
					if err != nil {
						s.sugar.Errorln(err)
						continue
					}
					if userResultEntry.Content.ItemContent.UserResults == nil {
						s.sugar.Errorln("userResults is nil")
						continue
					}
					parsedUser, err := userResultEntry.Content.ItemContent.UserResults.Result.Parse()
					if err != nil {
						s.sugar.Errorln(err)
						continue
					}

					// cache ids
					_, ok := cachedUserIds.Load(parsedUser.ScreenName)
					if !ok {
						s.sugar.Debugln("caching user id", parsedUser.ScreenName)
						cachedUserIds.Store(parsedUser.ScreenName, parsedUser.UserId)
					}

					userResults = append(userResults, *parsedUser)
				case "TimelineTimelineCursor":
					var cursorEntry entity.TimelineCursorEntry
					err := json.Unmarshal(entryRaw, &cursorEntry)
					if err != nil {
						s.sugar.Errorln(err)
						continue
					}
					if cursorEntry.Content.CursorType == "Bottom" {
						nextCursor = cursorEntry.Content.Value
					}
				default:
					fmt.Printf("unknown entry type: %s\n", entry.Content.EntryType)
				}
			}

		default:
			fmt.Printf("unknown instruction type: %s\n", instruction.Type)
		}
	}

	return userResults, nextCursor, nil
}
