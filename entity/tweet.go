package entity

import (
	"encoding/json"
	"fmt"
	"html"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// request

type UserTweetsVariables struct {
	UserId                                 string `json:"userId"`
	Cursor                                 string `json:"cursor,omitempty"`
	Count                                  int    `json:"count"`
	IncludePromotedContent                 bool   `json:"includePromotedContent"`
	WithQuickPromoteEligibilityTweetFields bool   `json:"withQuickPromoteEligibilityTweetFields"`
	WithSuperFollowsUserFields             bool   `json:"withSuperFollowsUserFields"`
	WithDownvotePerspective                bool   `json:"withDownvotePerspective"`
	WithReactionsMetadata                  bool   `json:"withReactionsMetadata"`
	WithReactionsPerspective               bool   `json:"withReactionsPerspective"`
	WithSuperFollowsTweetFields            bool   `json:"withSuperFollowsTweetFields"`
	WithVoice                              bool   `json:"withVoice"`
	WithV2Timeline                         bool   `json:"withV2Timeline"`
}

type UserTweetsFeatures struct {
	ResponsiveWebTwitterBlueVerifiedBadgeIsEnabled                 bool `json:"responsive_web_twitter_blue_verified_badge_is_enabled"`
	VerifiedPhoneLabelEnabled                                      bool `json:"verified_phone_label_enabled"`
	ResponsiveWebGraphqlTimelineNavigationEnabled                  bool `json:"responsive_web_graphql_timeline_navigation_enabled"`
	ViewCountsPublicVisibilityEnabled                              bool `json:"view_counts_public_visibility_enabled"`
	ViewCountsEverywhereApiEnabled                                 bool `json:"view_counts_everywhere_api_enabled"`
	TweetypieUnmentionOptimizationEnabled                          bool `json:"tweetypie_unmention_optimization_enabled"`
	ResponsiveWebUcGqlEnabled                                      bool `json:"responsive_web_uc_gql_enabled"`
	VibeApiEnabled                                                 bool `json:"vibe_api_enabled"`
	ResponsiveWebEditTweetApiEnabled                               bool `json:"responsive_web_edit_tweet_api_enabled"`
	GraphqlIsTranslatableRwebTweetIsTranslatableEnabled            bool `json:"graphql_is_translatable_rweb_tweet_is_translatable_enabled"`
	StandardizedNudgesMisinfo                                      bool `json:"standardized_nudges_misinfo"`
	TweetWithVisibilityResultsPreferGqlLimitedActionsPolicyEnabled bool `json:"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled"`
	InteractiveTextEnabled                                         bool `json:"interactive_text_enabled"`
	ResponsiveWebTextConversationsEnabled                          bool `json:"responsive_web_text_conversations_enabled"`
	ResponsiveWebEnhanceCardsEnabled                               bool `json:"responsive_web_enhance_cards_enabled"`
}

type UserMediaVariables struct {
	UserId                      string `json:"userId"`
	Cursor                      string `json:"cursor,omitempty"`
	Count                       int    `json:"count"`
	IncludePromotedContent      bool   `json:"includePromotedContent"`
	WithSuperFollowsUserFields  bool   `json:"withSuperFollowsUserFields"`
	WithDownvotePerspective     bool   `json:"withDownvotePerspective"`
	WithReactionsMetadata       bool   `json:"withReactionsMetadata"`
	WithReactionsPerspective    bool   `json:"withReactionsPerspective"`
	WithSuperFollowsTweetFields bool   `json:"withSuperFollowsTweetFields"`
	WithClientEventToken        bool   `json:"withClientEventToken"`
	WithBirdwatchNotes          bool   `json:"withBirdwatchNotes"`
	WithVoice                   bool   `json:"withVoice"`
	WithV2Timeline              bool   `json:"withV2Timeline"`
}

// response
type MediaPhoto struct {
	DisplayUrl    string      `json:"display_url"`
	ExpandedUrl   string      `json:"expanded_url"`
	ExtAltText    string      `json:"ext_alt_text"`
	IdStr         string      `json:"id_str"`
	Indices       []int       `json:"indices"`
	MediaUrlHttps string      `json:"media_url_https"`
	Type          string      `json:"type"`
	Url           string      `json:"url"`
	Features      interface{} `json:"features"`
	Sizes         interface{} `json:"sizes"`
	OriginalInfo  struct {
		Height int `json:"height"`
		Width  int `json:"width"`
	} `json:"original_info"`
}

type MediaVideo struct {
	DisplayUrl          string      `json:"display_url"`
	ExpandedUrl         string      `json:"expanded_url"`
	ExtAltText          string      `json:"ext_alt_text"`
	IdStr               string      `json:"id_str"`
	Indices             []int       `json:"indices"`
	MediaKey            string      `json:"media_key"`
	MediaUrlHttps       string      `json:"media_url_https"`
	Type                string      `json:"type"`
	Url                 string      `json:"url"`
	AdditionalMediaInfo interface{} `json:"additional_media_info"`
	ExtMediaColor       interface{} `json:"ext_media_color"`
	MediaStats          struct {
		ViewCount int `json:"viewCount"`
	} `json:"mediaStats"`
	ExtMediaAvailability struct {
		Status string `json:"status"`
	} `json:"ext_media_availability"`
	Features     interface{} `json:"features"`
	Sizes        interface{} `json:"sizes"`
	OriginalInfo struct {
		Height int `json:"height"`
		Width  int `json:"width"`
	} `json:"original_info"`
	VideoInfo struct {
		AspectRation   []int `json:"aspect_ratio"`
		DurationMillis int   `json:"duration_millis"`
		Variants       []struct {
			Bitrate     int    `json:"bitrate"`
			ContentType string `json:"content_type"`
			Url         string `json:"url"`
		} `json:"variants"`
	} `json:"video_info"`
}

type MediaType struct {
	Type string `json:"type"`
}

type Media = json.RawMessage

type MediaEntities struct {
	Media        []Media `json:"media"`
	UserMentions []struct {
		IdStr      string `json:"id_str"`
		Name       string `json:"name"`
		ScreenName string `json:"screen_name"`
		Indices    []int  `json:"indices"`
	} `json:"user_mentions"`
	Urls []struct {
		DisplayUrl  string `json:"display_url"`
		ExpandedUrl string `json:"expanded_url"`
		Url         string `json:"url"`
		Indices     []int  `json:"indices"`
	} `json:"urls"`
	Hashtags []struct {
		Indices []int  `json:"indices"`
		Text    string `json:"text"`
	} `json:"hashtags"`
	Symbols []interface{} `json:"symbols"`
}

type TweetLegacy struct {
	CreatedAt                 string        `json:"created_at"`
	ConversationIdStr         string        `json:"conversation_id_str"`
	DisplayTextRange          []int         `json:"display_text_range"`
	Entities                  MediaEntities `json:"entities"`
	ExtendedEntities          MediaEntities `json:"extended_entities"`
	FavoriteCount             int           `json:"favorite_count"`
	Favorited                 bool          `json:"favorited"`
	FullText                  string        `json:"full_text"`
	IsQuoteStatus             bool          `json:"is_quote_status"`
	Lang                      string        `json:"lang"`
	PossiblySensitive         bool          `json:"possibly_sensitive"`
	PossiblySensitiveEditable bool          `json:"possibly_sensitive_editable"`
	QuoteCount                int           `json:"quote_count"`
	ReplyCount                int           `json:"reply_count"`
	RetweetCount              int           `json:"retweet_count"`
	Retweeted                 bool          `json:"retweeted"`
	Source                    string        `json:"source"`
	UserIdStr                 string        `json:"user_id_str"`
	IdStr                     string        `json:"id_str"`
	RetweetedStatusResult     *TweetResults `json:"retweeted_status_result"`
}

type TweetResult struct {
	TypeName string       `json:"__typename"`
	Tweet    *TweetResult `json:"tweet"`
	RestId   string       `json:"rest_id"`
	Core     struct {
		UserResults struct {
			Result UserResult `json:"result"`
		} `json:"user_results"`
	} `json:"core"`
	UnmentionData           interface{}   `json:"unmention_data"`
	EditControl             interface{}   `json:"edit_control"`
	EditPerspective         interface{}   `json:"edit_perspective"`
	IsTranslateable         bool          `json:"is_translateable"`
	QuotedStatusResult      *TweetResults `json:"quoted_status_result"`
	Legacy                  TweetLegacy   `json:"legacy"`
	QuickPromoteEligibility interface{}   `json:"quick_promote_eligibility"`
	Views                   struct {
		Count string `json:"count"`
		State string `json:"state"`
	} `json:"views"`
}

type SocialContext struct {
	Type        string `json:"type"`
	ContextType string `json:"context_type"`
	Text        string `json:"text"`
	LandingUrl  struct {
		Url                string `json:"url"`
		UrlType            string `json:"url_type"`
		UrtEndPointOptions struct {
			CacheId       string `json:"cache_id"`
			Title         string `json:"title"`
			RequestParams []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"requestParams"`
		} `json:"urtEndpointOptions"`
	} `json:"landing_url"`
}

type TweetResults struct {
	Result TweetResult `json:"result"`
}

type PromotedMetadata struct {
}

type TimelineTweetEntry struct {
	EntryId   string `json:"entryId"`
	SortIndex string `json:"sortIndex"`
	Content   struct {
		EntryType   string `json:"entryType"`
		TypeName    string `json:"__typename"`
		ItemContent struct {
			ItemType         string            `json:"itemType"`
			TypeName         string            `json:"__typename"`
			TweetResults     *TweetResults     `json:"tweet_results"`
			TweetDisplayType string            `json:"tweetDisplayType"`
			PromotedMetadata *PromotedMetadata `json:"promotedMetadata"`
			SocialContext    *SocialContext    `json:"socialContext"`
			ReactiveTriggers interface{}       `json:"reactive_triggers"`
		} `json:"itemContent"`
		FeedbackInfo    interface{} `json:"feedbackInfo"`
		ClientEventInfo interface{} `json:"clientEventInfo"`
	} `json:"content"`
}

type TimelineTweetsResponse struct {
	Errors []TwitterError `json:"errors"`
	Data   struct {
		User struct {
			Result struct {
				TypeName   string `json:"__typename"`
				TimelineV2 struct {
					Timeline struct {
						Instructions    []TimelineInstruction `json:"instructions"`
						ResponseObjects interface{}           `json:"responseObjects"`
						Metadata        interface{}           `json:"metadata"`
					} `json:"timeline"`
				} `json:"timeline_v2"`
			} `json:"result"`
		} `json:"user"`
	} `json:"data"`
}

// function

func NewUserTweetsVariables(userId string, count int, cursor string) UserTweetsVariables {
	v := UserTweetsVariables{
		UserId:                                 userId,
		Count:                                  count,
		IncludePromotedContent:                 true,
		WithQuickPromoteEligibilityTweetFields: true,
		WithSuperFollowsUserFields:             true,
		WithDownvotePerspective:                false,
		WithReactionsMetadata:                  false,
		WithReactionsPerspective:               false,
		WithSuperFollowsTweetFields:            true,
		WithVoice:                              true,
		WithV2Timeline:                         true,
	}
	if cursor != "" {
		v.Cursor = cursor
	}
	return v
}

func NewUserTweetsFeatures() UserTweetsFeatures {
	return UserTweetsFeatures{
		ResponsiveWebTwitterBlueVerifiedBadgeIsEnabled:                 true,
		VerifiedPhoneLabelEnabled:                                      false,
		ResponsiveWebGraphqlTimelineNavigationEnabled:                  true,
		ViewCountsPublicVisibilityEnabled:                              true,
		ViewCountsEverywhereApiEnabled:                                 true,
		TweetypieUnmentionOptimizationEnabled:                          true,
		ResponsiveWebUcGqlEnabled:                                      true,
		VibeApiEnabled:                                                 true,
		ResponsiveWebEditTweetApiEnabled:                               true,
		GraphqlIsTranslatableRwebTweetIsTranslatableEnabled:            true,
		StandardizedNudgesMisinfo:                                      true,
		TweetWithVisibilityResultsPreferGqlLimitedActionsPolicyEnabled: false,
		InteractiveTextEnabled:                                         true,
		ResponsiveWebTextConversationsEnabled:                          false,
		ResponsiveWebEnhanceCardsEnabled:                               true,
	}
}

func NewUserMediaVariables(userId string, count int, cursor string) UserMediaVariables {
	v := UserMediaVariables{
		UserId:                      userId,
		Count:                       count,
		IncludePromotedContent:      true,
		WithSuperFollowsUserFields:  true,
		WithDownvotePerspective:     false,
		WithReactionsMetadata:       false,
		WithReactionsPerspective:    false,
		WithSuperFollowsTweetFields: true,
		WithClientEventToken:        true,
		WithBirdwatchNotes:          true,
		WithVoice:                   true,
		WithV2Timeline:              true,
	}
	if cursor != "" {
		v.Cursor = cursor
	}
	return v
}

func NewUserTweetsParams(userId string, count int, cursor string) (url.Values, error) {
	variables, err := json.Marshal(NewUserTweetsVariables(userId, count, cursor))
	if err != nil {
		return nil, fmt.Errorf("failed to marshal variables: %v", err)
	}
	features, err := json.Marshal(NewUserTweetsFeatures())
	if err != nil {
		return nil, fmt.Errorf("failed to marshal features: %v", err)
	}
	return url.Values{
		"variables": {string(variables)},
		"features":  {string(features)},
	}, nil
}

func NewUserMediaParams(userId string, count int, cursor string) (url.Values, error) {
	variables, err := json.Marshal(NewUserMediaVariables(userId, count, cursor))
	if err != nil {
		return nil, fmt.Errorf("failed to marshal variables: %v", err)
	}
	features, err := json.Marshal(NewUserTweetsFeatures())
	if err != nil {
		return nil, fmt.Errorf("failed to marshal features: %v", err)
	}
	return url.Values{
		"variables": {string(variables)},
		"features":  {string(features)},
	}, nil
}

type ParsedMedia interface{}

type ParsedMediaPhoto struct {
	Height  int
	Width   int
	AltText string
	Url     string
}

type ParsedMediaVideo struct {
	Height        int
	Width         int
	AltText       string
	IsAnimatedGif bool
	ThumbUrl      string
	Bitrate       int
	ContentType   string
	DurationMs    int
	Url           string
}

type UserMentions struct {
	UserId     string
	Name       string
	ScreenName string
}

type ParsedTweet struct {
	CreatedAt time.Time
	TweetId   string
	Entities  struct {
		Media        []ParsedMedia
		UserMentions []UserMentions
		Urls         []string
		Hashtags     []string
	}
	FavouriteCount    int
	Favourited        bool
	FullText          string
	Lang              string
	PossiblySensitive bool
	QuoteCount        int
	ReplyCount        int
	RetweetedCount    int
	Retweeted         bool
	IsRetweet         bool
	RetweetedTweet    *ParsedTweet
	IsReply           bool
	RepliedTweet      *ParsedTweet
	IsRecommended     bool
	Url               string
	Views             int
	ParsedUser        ParsedUser
}

func (t *TweetResult) Parse() (*ParsedTweet, error) {
	if t.TypeName == "TweetWithVisibilityResults" { // wtf
		return t.Tweet.Parse()
	}

	userResult, err := t.Core.UserResults.Result.Parse()
	if err != nil {
		return nil, fmt.Errorf("failed to parse user: %v", err)
	}

	var repliedTweet *ParsedTweet
	isReply := t.QuotedStatusResult != nil
	if isReply {
		repliedTweet, err = t.QuotedStatusResult.Result.Parse()
		if err != nil {
			return nil, fmt.Errorf("failed to parse replied tweet: %v", err)
		}
	}

	createdAt, err := time.Parse(time.RubyDate, t.Legacy.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse created at: %v", err)
	}

	fullText := html.UnescapeString(t.Legacy.FullText)

	var medias []ParsedMedia
	extendedEntities := t.Legacy.ExtendedEntities
	for i, mediaRaw := range extendedEntities.Media {
		var media MediaType
		err = json.Unmarshal(mediaRaw, &media)
		if err != nil {
			fmt.Printf("error parsing media entity %d: %v", i, err)
			continue
		}
		switch media.Type {
		case "photo":
			mediaPhoto := MediaPhoto{}
			err = json.Unmarshal(mediaRaw, &mediaPhoto)
			if err != nil {
				fmt.Printf("error parsing photo entity %d: %v", i, err)
				continue
			}
			fullText = strings.Replace(fullText, mediaPhoto.Url, "", 1)
			medias = append(medias, ParsedMediaPhoto{
				Height:  mediaPhoto.OriginalInfo.Height,
				Width:   mediaPhoto.OriginalInfo.Width,
				AltText: mediaPhoto.ExtAltText,
				Url:     mediaPhoto.MediaUrlHttps,
			})
		case "video", "animated_gif":
			mediaVideo := MediaVideo{}
			err = json.Unmarshal(mediaRaw, &mediaVideo)
			if err != nil {
				fmt.Printf("error parsing video entity %d: %v", i, err)
				continue
			}
			fullText = strings.Replace(fullText, mediaVideo.Url, "", 1)
			parsedMediaVideo := ParsedMediaVideo{
				Height:        mediaVideo.OriginalInfo.Height,
				Width:         mediaVideo.OriginalInfo.Width,
				AltText:       mediaVideo.ExtAltText,
				IsAnimatedGif: media.Type == "animated_gif",
				ThumbUrl:      mediaVideo.MediaUrlHttps,
				DurationMs:    mediaVideo.VideoInfo.DurationMillis,
			}
			if len(mediaVideo.VideoInfo.Variants) > 0 {
				maxBitrate := 0
				for _, variant := range mediaVideo.VideoInfo.Variants {
					if variant.Bitrate > maxBitrate {
						maxBitrate = variant.Bitrate
						parsedMediaVideo.Bitrate = variant.Bitrate
						parsedMediaVideo.ContentType = variant.ContentType
						parsedMediaVideo.Url = variant.Url
					}
				}
			}
			medias = append(medias, parsedMediaVideo)
		default:
			fmt.Printf("unknown media type %s", media.Type)
		}
	}

	entities := t.Legacy.Entities
	var userMentions []UserMentions
	for _, entity := range entities.UserMentions {
		userMentions = append(userMentions, UserMentions{
			UserId:     entity.IdStr,
			Name:       entity.Name,
			ScreenName: entity.ScreenName,
		})
	}
	var urls []string
	for _, entity := range entities.Urls {
		fullText = strings.ReplaceAll(fullText, entity.Url, entity.ExpandedUrl)
		urls = append(urls, entity.ExpandedUrl)
	}
	var hashtags []string
	for _, entity := range entities.Hashtags {
		hashtags = append(hashtags, entity.Text)
	}

	isRetweet := t.Legacy.RetweetedStatusResult != nil
	var retweetedTweet *ParsedTweet
	if isRetweet {
		retweetedTweet, err = t.Legacy.RetweetedStatusResult.Result.Parse()
		if err != nil {
			return nil, fmt.Errorf("failed to parse retweet: %v", err)
		}
	}

	views := 0
	viewsRaw := t.Views.Count
	if viewsRaw != "" {
		views, err = strconv.Atoi(viewsRaw)
		if err != nil {
			return nil, fmt.Errorf("failed to parse views: %v", err)
		}
	}

	parsedTweet := ParsedTweet{
		CreatedAt: createdAt,
		TweetId:   t.Legacy.IdStr,
		Entities: struct {
			Media        []ParsedMedia
			UserMentions []UserMentions
			Urls         []string
			Hashtags     []string
		}{
			Media:        medias,
			UserMentions: userMentions,
			Urls:         urls,
			Hashtags:     hashtags,
		},
		FavouriteCount:    t.Legacy.FavoriteCount,
		Favourited:        t.Legacy.Favorited,
		FullText:          strings.TrimSpace(fullText),
		Lang:              t.Legacy.Lang,
		PossiblySensitive: t.Legacy.PossiblySensitive,
		QuoteCount:        t.Legacy.QuoteCount,
		ReplyCount:        t.Legacy.ReplyCount,
		RetweetedCount:    t.Legacy.RetweetCount,
		Retweeted:         t.Legacy.Retweeted,
		IsRetweet:         isRetweet,
		RetweetedTweet:    retweetedTweet,
		IsReply:           isReply,
		RepliedTweet:      repliedTweet,
		Url:               fmt.Sprintf("https://twitter.com/%s/status/%s", userResult.ScreenName, t.Legacy.IdStr),
		Views:             views,
		ParsedUser:        *userResult,
	}
	return &parsedTweet, nil
}
