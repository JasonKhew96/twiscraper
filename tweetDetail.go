package twiscraper

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/JasonKhew96/twiscraper/entity"
)

func (s *Scraper) GetTweetDetail(focalTweetId string) (*entity.ParsedTweet, error) {
	vl, err := entity.NewTweetDetailParams(focalTweetId)
	if err != nil {
		return nil, fmt.Errorf("failed to create tweet detail params: %v", err)
	}
	apiUrl, err := url.Parse(apiTweetDetail)
	if err != nil {
		return nil, fmt.Errorf("failed to parse tweet detail url: %v", err)
	}
	apiUrl.RawQuery = vl.Encode()
	req, err := http.NewRequest(http.MethodGet, apiUrl.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	var tweetDetailResponse entity.TweetDetailResponse
	err = s.requestAPI(req, &tweetDetailResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to request api: %v", err)
	}
	// if len(tweetDetailResponse.Errors) > 0 {
	// 	return nil, errors.New(tweetDetailResponse.Errors[0].Message)
	// }

	var parsedTweet *entity.ParsedTweet
	for _, instructionRaw := range tweetDetailResponse.Data.ThreadedConversationWithInjectionsV2.Instructions {
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
			var replies [][]*entity.ParsedTweet
			for _, entryRaw := range timelineAddEntries.Entries {
				var entry entity.Entry
				err := json.Unmarshal(entryRaw, &entry)
				if err != nil {
					s.sugar.Errorln(err)
					continue
				}
				if strings.HasPrefix(entry.EntryId, "cursor-") {
					continue
				}
				switch entry.Content.EntryType {
				case "TimelineTimelineItem":
					var tweetEntry entity.TimelineTweetEntry
					err := json.Unmarshal(entryRaw, &tweetEntry)
					if err != nil {
						return nil, fmt.Errorf("failed to unmarshal tweet entry: %v", err)
					}
					if tweetEntry.Content.ItemContent.TweetResults == nil {
						return nil, fmt.Errorf("tweet entry has no tweet results")
					}
					if tweetEntry.Content.ItemContent.PromotedMetadata != nil {
						return nil, fmt.Errorf("tweet is promoted tweet")
					}
					parsedTweet, err = tweetEntry.Content.ItemContent.TweetResults.Result.Parse()
					if err != nil {
						return nil, fmt.Errorf("failed to parse tweet: %v", err)
					}
					if parsedTweet.TweetId != focalTweetId {
						continue
					}
					parsedTweet.IsRecommended = tweetEntry.Content.ItemContent.SocialContext != nil
					// return parsedTweet, nil
				case "TimelineTimelineModule":
					var tweetEntry entity.TimelineTweetEntry
					err := json.Unmarshal(entryRaw, &tweetEntry)
					if err != nil {
						return nil, fmt.Errorf("failed to unmarshal tweet entry: %v", err)
					}
					if len(tweetEntry.Content.Items) <= 0 {
						continue
					}
					var innerReplies []*entity.ParsedTweet
					for _, conversations := range tweetEntry.Content.Items {
						if conversations.Item.ItemContent.ItemType == "TimelineTimelineCursor" {
							// show more replies
							continue
						}
						if conversations.Item.ItemContent.TweetResults == nil {
							return nil, fmt.Errorf("conversationthread has no tweet results")
						}
						parsedReplyTweet, err := conversations.Item.ItemContent.TweetResults.Result.Parse()
						if err != nil {
							return nil, fmt.Errorf("failed to parse conversationthread: %v", err)
						}
						if conversations.Item.ItemContent.PromotedMetadata != nil {
							// is promoted content
							continue
						}
						innerReplies = append(innerReplies, parsedReplyTweet)
					}
					replies = append(replies, innerReplies)
				default:
					fmt.Printf("unknown entry type: %s\n", entry.Content.EntryType)
				}
			}
			parsedTweet.Replies = replies
		default:
			fmt.Printf("unknown instruction type: %s\n", instruction.Type)
		}
	}

	if parsedTweet == nil {
		return nil, errors.New("no tweet found")
	}

	return parsedTweet, nil
}
