package twiscraper

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/JasonKhew96/twiscraper/entity"
)

func (s *Scraper) GetHomeTimeline(ctx context.Context, count int) <-chan *TimelineResult {
	return s.getHomeTimelineStream(ctx, FetchHomeTimeline, count, s.fetchHomeTimeline)
}

func (s *Scraper) GetHomeLatestTimeline(ctx context.Context, count int) <-chan *TimelineResult {
	return s.getHomeTimelineStream(ctx, FetchHomeLatestTimeline, count, s.fetchHomeTimeline)
}

func (s *Scraper) getHomeTimelineStream(ctx context.Context, opt fetchOptions, count int, fetchFunc fetchHomeTimelineFunc) <-chan *TimelineResult {
	ch := make(chan *TimelineResult)
	go func() {
		defer close(ch)
		var nextCursor string
		tweetsCount := 0
		for tweetsCount < count {
			select {
			case <-ctx.Done():
				ch <- &TimelineResult{Error: ctx.Err()}
				return
			default:
			}

			tweets, cursor, err := fetchFunc(opt, count, nextCursor)
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
	}()
	return ch
}

func (s *Scraper) fetchHomeTimeline(opt fetchOptions, count int, cursor string) ([]entity.ParsedTweet, string, error) {
	if !s.IsLogined() {
		return nil, "", ErrorNotLogined
	}

	if count > 20 {
		count = 20
	}

	var err error
	var u string
	var vl url.Values
	switch opt {
	case FetchHomeTimeline:
		vl, err = entity.NewHomeTimelineParams(count, cursor)
		u = apiHomeTimeline
	case FetchHomeLatestTimeline:
		vl, err = entity.NewHomeLatestTimelineParams(count, cursor)
		u = apiHomeLatestTimeline
	default:
		return nil, "", fmt.Errorf("invalid fetch option")
	}
	if err != nil {
		return nil, "", fmt.Errorf("failed to create home timeline params: %v", err)
	}

	apiUrl, err := url.Parse(u)
	if err != nil {
		return nil, "", fmt.Errorf("failed to parse api url: %v", err)
	}
	apiUrl.RawQuery = vl.Encode()
	req, err := http.NewRequest(http.MethodGet, apiUrl.String(), nil)
	if err != nil {
		return nil, "", fmt.Errorf("failed to create request: %v", err)
	}
	var homeTimelineTweets entity.HomeTimelineResponse
	err = s.requestAPI(req, &homeTimelineTweets)
	if err != nil {
		return nil, "", fmt.Errorf("failed to request api: %v", err)
	}
	// if len(homeTimelineTweets.Errors) > 0 {
	// 	return nil, "", errors.New(homeTimelineTweets.Errors[0].Message)
	// }

	var tweetResults []entity.ParsedTweet
	var nextCursor string

	for _, instructionRaw := range homeTimelineTweets.Data.Home.HomeTimelineUrt.Instructions {
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
		case "TimelineShowCover":
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
					if tweetEntry.Content.ItemContent.TweetResults == nil {
						s.sugar.Errorln("tweet results is nil")
						continue
					}
					if tweetEntry.Content.ItemContent.PromotedMetadata != nil {
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
					if !strings.HasPrefix(entry.EntryId, "home-conversation-") {
						continue
					}
					var tweetEntry entity.TimelineTweetEntry
					err := json.Unmarshal(entryRaw, &tweetEntry)
					if err != nil {
						s.sugar.Errorln(err)
						continue
					}
					for _, item := range tweetEntry.Content.Items {
						if item.Item.ItemContent.TweetResults == nil {
							s.sugar.Errorln("tweet results is nil")
							continue
						}
						if item.Item.ItemContent.PromotedMetadata != nil {
							continue
						}
						parsedTweet, err := item.Item.ItemContent.TweetResults.Result.Parse()
						if err != nil {
							s.sugar.Errorln(err)
							continue
						}
						tweetResults = append(tweetResults, *parsedTweet)
					}
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
