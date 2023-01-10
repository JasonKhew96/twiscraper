package twiscraper

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/JasonKhew96/twiscraper/entity"
)

func (s *Scraper) GetTweetDetail(focalTweetId string) (*entity.ParsedTweet, error) {
	vl, err := entity.NewTweetDetailParams(focalTweetId)
	if err != nil {
		return nil, err
	}
	apiUrl, err := url.Parse(apiTweetDetail)
	if err != nil {
		return nil, err
	}
	apiUrl.RawQuery = vl.Encode()
	req, err := http.NewRequest(http.MethodGet, apiUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	var tweetDetailResponse entity.TweetDetailResponse
	err = s.requestAPI(req, &tweetDetailResponse)
	if err != nil {
		return nil, err
	}
	if len(tweetDetailResponse.Errors) > 0 {
		return nil, errors.New(tweetDetailResponse.Errors[0].Message)
	}

	for _, instructionRaw := range tweetDetailResponse.Data.ThreadedConversationWithInjectionsV2.Instructions {
		var instruction entity.Instruction
		err := json.Unmarshal(instructionRaw, &instruction)
		if err != nil {
			fmt.Println(err)
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
				fmt.Println(err)
				continue
			}
			for _, entryRaw := range timelineAddEntries.Entries {
				var entry entity.Entry
				err := json.Unmarshal(entryRaw, &entry)
				if err != nil {
					fmt.Println(err)
					continue
				}
				switch entry.Content.EntryType {
				case "TimelineTimelineItem":
					var tweetEntry entity.TimelineTweetEntry
					err := json.Unmarshal(entryRaw, &tweetEntry)
					if err != nil {
						fmt.Println(err)
						continue
					}
					parsedTweet, err := tweetEntry.Content.ItemContent.TweetResults.Result.Parse()
					if err != nil {
						fmt.Println(err)
						continue
					}
					parsedTweet.IsRecommended = tweetEntry.Content.ItemContent.SocialContext != nil
					return parsedTweet, nil
				case "TimelineTimelineModule":
				default:
					fmt.Printf("unknown entry type: %s\n", entry.Content.EntryType)
				}
			}
		default:
			fmt.Printf("unknown instruction type: %s\n", instruction.Type)
		}
	}
	return nil, errors.New("no tweet found")
}
