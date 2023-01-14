package twiscraper

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/JasonKhew96/twiscraper/entity"
)

type TimelineResult struct {
	ParsedTweet entity.ParsedTweet
	Error       error
}

func (s *Scraper) GetTimelineTweets(ctx context.Context, screenName string, count int) <-chan *TimelineResult {
	return s.getTimelineStream(ctx, FetchTweets, screenName, count, s.fetchTimeline)
}

func (s *Scraper) GetTimelineMedia(ctx context.Context, screenName string, count int) <-chan *TimelineResult {
	return s.getTimelineStream(ctx, FetchMedia, screenName, count, s.fetchTimeline)
}

func (s *Scraper) getTimelineStream(ctx context.Context, opt fetchOptions, screenName string, count int, fetchFunc fetchTimelineFunc) <-chan *TimelineResult {
	ch := make(chan *TimelineResult)
	go func(screenName string) {
		defer close(ch)

		id, err := s.GetUserIdByScreenName(screenName)
		if err != nil {
			ch <- &TimelineResult{Error: err}
			return
		}

		var nextCursor string
		tweetsCount := 0
		for tweetsCount < count {
			select {
			case <-ctx.Done():
				ch <- &TimelineResult{Error: ctx.Err()}
				return
			default:
			}

			tweets, cursor, err := fetchFunc(opt, id, count, nextCursor)
			if err != nil {
				ch <- &TimelineResult{Error: err}
				return
			}

			if len(tweets) == 0 {
				break
			}

			for _, tweet := range tweets {
				select {
				case <-ctx.Done():
					ch <- &TimelineResult{Error: ctx.Err()}
					return
				default:
				}

				if tweetsCount < count {
					nextCursor = cursor
					ch <- &TimelineResult{ParsedTweet: tweet}
				} else {
					break
				}
				tweetsCount++
			}
		}
	}(screenName)
	return ch
}

func (s *Scraper) fetchTimeline(opt fetchOptions, id string, count int, cursor string) ([]entity.ParsedTweet, string, error) {
	var err error
	var u string
	var vl url.Values
	switch opt {
	case FetchTweets:
		if count > 40 {
			count = 40
		}
		vl, err = entity.NewUserTweetsParams(id, count, cursor)
		u = apiUserTweets
	case FetchMedia:
		if !s.IsLogined() {
			return nil, "", ErrorNotLogined
		}
		if count > 20 {
			count = 20
		}
		vl, err = entity.NewUserMediaParams(id, count, cursor)
		u = apiUserMedia
	default:
		return nil, "", fmt.Errorf("invalid fetch option")
	}
	if err != nil {
		return nil, "", fmt.Errorf("failed to create params: %v", err)
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
	var timelineTweets entity.TimelineTweetsResponse
	err = s.requestAPI(req, &timelineTweets)
	if err != nil {
		return nil, "", fmt.Errorf("failed to request api: %v", err)
	}
	// if len(timelineTweets.Errors) > 0 {
	// 	return nil, "", errors.New(timelineTweets.Errors[0].Message)
	// }

	var tweetResults []entity.ParsedTweet
	var nextCursor string

	for _, instructionRaw := range timelineTweets.Data.User.Result.TimelineV2.Timeline.Instructions {
		var instruction entity.Instruction
		err := json.Unmarshal(instructionRaw, &instruction)
		if err != nil {
			s.sugar.Errorln(err)
			continue
		}
		switch instruction.Type {
		case "TimelinePinEntry":
			// TODO
			continue
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
					var tweetEntry entity.TimelineTweetEntry
					err := json.Unmarshal(entryRaw, &tweetEntry)
					if err != nil {
						s.sugar.Errorln(err)
						continue
					}
					parsedTweet, err := tweetEntry.Content.ItemContent.TweetResults.Result.Parse()
					if err != nil {
						s.sugar.Errorln(err)
						continue
					}
					parsedTweet.IsRecommended = tweetEntry.Content.ItemContent.SocialContext != nil
					tweetResults = append(tweetResults, *parsedTweet)
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
				case "TimelineTimelineModule":
					// TODO
					continue
				default:
					fmt.Printf("unknown entry type: %s\n", entry.Content.EntryType)
				}
			}
		default:
			fmt.Printf("unknown instruction type: %s\n", instruction.Type)
		}
	}

	return tweetResults, nextCursor, nil
}
