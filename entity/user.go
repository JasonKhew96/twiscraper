package entity

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"
)

// request

type UserByScreenNameVariables struct {
	ScreenName                 string `json:"screen_name"`
	WithSafetyModeUserFields   bool   `json:"withSafetyModeUserFields"`
	WithSuperFollowsUserFields bool   `json:"withSuperFollowsUserFields"`
}

func NewUserByScreenNameVariables(screenName string) UserByScreenNameVariables {
	return UserByScreenNameVariables{
		ScreenName:                 screenName,
		WithSafetyModeUserFields:   true,
		WithSuperFollowsUserFields: true,
	}
}

type UserByScreenNameFeatures struct {
	ResponsiveWebTwitterBlueVerifiedBadgeIsEnabled       bool `json:"responsive_web_twitter_blue_verified_badge_is_enabled"`
	ResponsiveWebGraphqlExcludeDirectiveEnabled          bool `json:"responsive_web_graphql_exclude_directive_enabled"`
	VerifiedPhoneLabelEnabled                            bool `json:"verified_phone_label_enabled"`
	ResponsiveWebGraphqlSkipUserProfileImageExtensionsEn bool `json:"responsive_web_graphql_skip_user_profile_image_extensions_enabled"`
	ResponsiveWebGraphqlTimelineNavigationEnabled        bool `json:"responsive_web_graphql_timeline_navigation_enabled"`
}

func NewUserByScreenNameFeatures() UserByScreenNameFeatures {
	return UserByScreenNameFeatures{
		ResponsiveWebTwitterBlueVerifiedBadgeIsEnabled:       true,
		ResponsiveWebGraphqlExcludeDirectiveEnabled:          false,
		VerifiedPhoneLabelEnabled:                            false,
		ResponsiveWebGraphqlSkipUserProfileImageExtensionsEn: false,
		ResponsiveWebGraphqlTimelineNavigationEnabled:        true,
	}
}

func NewUserByScreenNameParams(screenName string) (url.Values, error) {
	variables, err := json.Marshal(NewUserByScreenNameVariables(screenName))
	if err != nil {
		return nil, fmt.Errorf("failed to marshal variables: %v", err)
	}
	features, err := json.Marshal(NewUserByScreenNameFeatures())
	if err != nil {
		return nil, fmt.Errorf("failed to marshal features: %v", err)
	}
	return url.Values{
		"variables": {string(variables)},
		"features":  {string(features)},
	}, nil
}

// response

type Url struct {
	DisplayUrl  string `json:"display_url"`
	ExpandedUrl string `json:"expanded_url"`
	Url         string `json:"url"`
	Indices     []int  `json:"indices"`
}

type Urls struct {
	Urls []Url `json:"urls"`
}

type Entities struct {
	Description Urls `json:"description"`
	Url         Urls `json:"url"`
}

type BirthDate struct {
	Day            int    `json:"day"`
	Month          int    `json:"month"`
	Visibility     string `json:"visibility"`
	YearVisibility string `json:"year_visibility"`
}

type LegacyExtendedProfile struct {
	BirthDate BirthDate `json:"birthdate"`
}

type VerificationInfo struct {
	Reason struct {
		Description struct {
			Text     string     `json:"text"`
			Entities []Entities `json:"entities"`
		} `json:"description"`
	} `json:"reason"`
}

type UserLegacy struct {
	BlockedBy               bool          `json:"blocked_by"`
	Blocking                bool          `json:"blocking"`
	CanDm                   bool          `json:"can_dm"`
	CanMediaTag             bool          `json:"can_media_tag"`
	CreatedAt               string        `json:"created_at"`
	DefaultProfile          bool          `json:"default_profile"`
	DefaultProfileImage     bool          `json:"default_profile_image"`
	Description             string        `json:"description"`
	Entities                Entities      `json:"entities"`
	FastFollowersCount      int           `json:"fast_followers_count"`
	FavouritesCount         int           `json:"favourites_count"`
	FollowRequestSent       bool          `json:"follow_request_sent"`
	FollowedBy              bool          `json:"followed_by"`
	FollowersCount          int           `json:"followers_count"`
	Following               bool          `json:"following"`
	FriendsCount            int           `json:"friends_count"`
	HasCustomTimelines      bool          `json:"has_custom_timelines"`
	IsTranslator            bool          `json:"is_translator"`
	ListedCount             int           `json:"listed_count"`
	Location                string        `json:"location"`
	MediaCount              int           `json:"media_count"`
	Muting                  bool          `json:"muting"`
	Name                    string        `json:"name"`
	NormalFollowersCount    int           `json:"normal_followers_count"`
	Notifications           bool          `json:"notifications"`
	PinnedTweetIdsStr       []string      `json:"pinned_tweet_ids_str"`
	PossiblySensitive       bool          `json:"possibly_sensitive"`
	ProfileBannerExtensions interface{}   `json:"profile_banner_extensions"`
	ProfileBannerUrl        string        `json:"profile_banner_url"`
	ProfileImageExtensions  interface{}   `json:"profile_image_extensions"`
	ProfileImageUrlHttps    string        `json:"profile_image_url_https"`
	ProfileInterstitialType interface{}   `json:"profile_interstitial_type"`
	Protected               bool          `json:"protected"`
	ScreenName              string        `json:"screen_name"`
	StatusesCount           int           `json:"statuses_count"`
	TranslatorType          string        `json:"translator_type"`
	Url                     string        `json:"url"`
	Verified                bool          `json:"verified"`
	VerifiedType            string        `json:"verified_type"`
	WantRetweets            bool          `json:"want_retweets"`
	WithHeldInCountries     []interface{} `json:"withheld_in_countries"`
}

type UserResult struct {
	TypeName                   string                `json:"__typename"`
	Id                         string                `json:"id"`
	Reason                     string                `json:"reason"`
	RestId                     string                `json:"rest_id"`
	AffiliatesHighlightedLabel interface{}           `json:"affiliates_highlighted_label"`
	HasGraduatedAccess         bool                  `json:"has_graduated_access"`
	HasNftAvatar               bool                  `json:"has_nft_avatar"`
	IsBlueVerified             bool                  `json:"is_blue_verified"`
	Legacy                     UserLegacy            `json:"legacy"`
	SmartBlockedBy             bool                  `json:"smart_blocked_by"`
	SmartBlocking              bool                  `json:"smart_blocking"`
	SuperFollowEligible        bool                  `json:"super_follow_eligible"`
	SuperFollowedBy            bool                  `json:"super_followed_by"`
	SuperFollowing             bool                  `json:"super_following"`
	LegacyExtendedProfile      LegacyExtendedProfile `json:"legacy_extended_profile"`
	IsProfileTranslatable      bool                  `json:"is_profile_translatable"`
	VerificationInfo           VerificationInfo      `json:"verification_info"`
}

type UserByScreenName struct {
	Errors []TwitterError `json:"errors"`
	Data   struct {
		User struct {
			Result UserResult `json:"result"`
		} `json:"user"`
	} `json:"data"`
}

// request

type FollowersVariables struct {
	UserId                      string `json:"userId"`
	Count                       int    `json:"count"`
	Cursor                      string `json:"cursor,omitempty"`
	IncludePromotedContent      bool   `json:"includePromotedContent"`
	WithSuperFollowsUserFields  bool   `json:"withSuperFollowsUserFields"`
	WithDownvotePerspective     bool   `json:"withDownvotePerspective"`
	WithReactionsMetadata       bool   `json:"withReactionsMetadata"`
	WithReactionsPerspective    bool   `json:"withReactionsPerspective"`
	WithSuperFollowsTweetFields bool   `json:"withSuperFollowsTweetFields"`
}

func NewFollowersVariables(userId string, count int, cursor string) FollowersVariables {
	v := FollowersVariables{
		UserId:                      userId,
		Count:                       count, // 20
		IncludePromotedContent:      false,
		WithSuperFollowsUserFields:  true,
		WithDownvotePerspective:     false,
		WithReactionsMetadata:       false,
		WithReactionsPerspective:    false,
		WithSuperFollowsTweetFields: true,
	}
	if cursor != "" {
		v.Cursor = cursor
	}
	return v
}

type FollowersFeatures struct {
	ResponsiveWebTwitterBlueVerifiedBadgeIsEnabled                 bool `json:"responsive_web_twitter_blue_verified_badge_is_enabled"`
	ResponsiveWebGraphqlExcludeDirectiveEnabled                    bool `json:"responsive_web_graphql_exclude_directive_enabled"`
	VerifiedPhoneLabelEnabled                                      bool `json:"verified_phone_label_enabled"`
	ResponsiveWebGraphqlTimelineNavigationEnabled                  bool `json:"responsive_web_graphql_timeline_navigation_enabled"`
	ResponsiveWebGraphqlSkipUserProfileImageExtensionsEnabled      bool `json:"responsive_web_graphql_skip_user_profile_image_extensions_enabled"`
	TweetypieUnmentionOptimizationEnabled                          bool `json:"tweetypie_unmention_optimization_enabled"`
	VibeApiEnabled                                                 bool `json:"vibe_api_enabled"`
	ResponsiveWebEditTweetApiEnabled                               bool `json:"responsive_web_edit_tweet_api_enabled"`
	GraphqlIsTranslatableRwebTweetIsTranslatableEnabled            bool `json:"graphql_is_translatable_rweb_tweet_is_translatable_enabled"`
	ViewCountsEverywhereApiEnabled                                 bool `json:"view_counts_everywhere_api_enabled"`
	LongformNotetweetsConsumptionEnabled                           bool `json:"longform_notetweets_consumption_enabled"`
	FreedomOfSpeechNotReachAppealLabelEnabled                      bool `json:"freedom_of_speech_not_reach_appeal_label_enabled"`
	StandardizedNudgesMisinfo                                      bool `json:"standardized_nudges_misinfo"`
	TweetWithVisibilityResultsPreferGqlLimitedActionsPolicyEnabled bool `json:"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled"`
	InteractiveTextEnabled                                         bool `json:"interactive_text_enabled"`
	ResponsiveWebTextConversationsEnabled                          bool `json:"responsive_web_text_conversations_enabled"`
	ResponsiveWebEnhanceCardsEnabled                               bool `json:"responsive_web_enhance_cards_enabled"`
}

func NewFollowersFeatures() FollowersFeatures {
	return FollowersFeatures{
		ResponsiveWebTwitterBlueVerifiedBadgeIsEnabled:                 true,
		ResponsiveWebGraphqlExcludeDirectiveEnabled:                    false,
		VerifiedPhoneLabelEnabled:                                      false,
		ResponsiveWebGraphqlTimelineNavigationEnabled:                  true,
		ResponsiveWebGraphqlSkipUserProfileImageExtensionsEnabled:      false,
		TweetypieUnmentionOptimizationEnabled:                          true,
		VibeApiEnabled:                                                 true,
		ResponsiveWebEditTweetApiEnabled:                               true,
		GraphqlIsTranslatableRwebTweetIsTranslatableEnabled:            true,
		ViewCountsEverywhereApiEnabled:                                 true,
		LongformNotetweetsConsumptionEnabled:                           true,
		FreedomOfSpeechNotReachAppealLabelEnabled:                      false,
		StandardizedNudgesMisinfo:                                      true,
		TweetWithVisibilityResultsPreferGqlLimitedActionsPolicyEnabled: false,
		InteractiveTextEnabled:                                         true,
		ResponsiveWebTextConversationsEnabled:                          false,
		ResponsiveWebEnhanceCardsEnabled:                               false,
	}
}

func NewFollowersParams(userId string, count int, cursor string) (url.Values, error) {
	variables, err := json.Marshal(NewFollowersVariables(userId, count, cursor))
	if err != nil {
		return nil, fmt.Errorf("failed to marshal variables: %v", err)
	}
	features, err := json.Marshal(NewFollowersFeatures())
	if err != nil {
		return nil, fmt.Errorf("failed to marshal features: %v", err)
	}
	return url.Values{
		"variables": {string(variables)},
		"features":  {string(features)},
	}, nil
}

// response

type UserResults struct {
	Result UserResult `json:"result"`
}

type UserResultEntry struct {
	EntryId   string `json:"entryId"`
	SortIndex string `json:"sortIndex"`
	Content   struct {
		EntryType   string `json:"entryType"`
		TypeName    string `json:"__typename"`
		ItemContent struct {
			ItemType        string       `json:"itemType"`
			TypeName        string       `json:"__typename"`
			UserResults     *UserResults `json:"user_results"`
			UserDisplayType string       `json:"userDisplayType"`
		} `json:"itemContent"`
		ClientEventInfo struct {
			Component string `json:"component"`
			Element   string `json:"element"`
		} `json:"clientEventInfo"`
	} `json:"content"`
}

type TimelineCursorEntry struct {
	EntryId   string `json:"entryId"`
	SortIndex string `json:"sortIndex"`
	Content   struct {
		EntryType  string `json:"entryType"`
		TypeName   string `json:"__typename"`
		Value      string `json:"value"`
		CursorType string `json:"cursorType"`
	} `json:"content"`
}

type FollowersResponse struct {
	Errors []TwitterError `json:"errors"`
	Data   struct {
		User struct {
			Result struct {
				TypeName string `json:"__typename"`
				Timeline struct {
					Timeline struct {
						Instructions []TimelineInstruction `json:"instructions"`
					} `json:"timeline"`
				} `json:"timeline"`
			} `json:"result"`
		} `json:"user"`
	} `json:"data"`
}

// function

type ParsedUser struct {
	UserId         string
	IsBlueVerified bool
	BirthData      struct {
		Day   int
		Month int
	}
	IsBlockedBy     bool
	IsBlocked       bool
	CreatedAt       time.Time
	Description     string
	FavouritesCount int
	FollowedBy      bool
	FollowersCount  int
	IsFollowing     bool
	FollowingCount  int
	Location        string
	MediaCount      int
	PinnedTweetIds  []string
	IsProtected     bool
	ScreenName      string
	StatusesCount   int
	Url             string
	IsVerified      bool
	VerifiedType    string
}

func (u *UserResult) Parse() (*ParsedUser, error) {
	createdAt, err := time.Parse(time.RubyDate, u.Legacy.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse created_at: %v", err)
	}
	description := u.Legacy.Description
	for _, entity := range u.Legacy.Entities.Description.Urls {
		description = strings.ReplaceAll(description, entity.Url, entity.ExpandedUrl)
	}
	userUrl := ""
	if len(u.Legacy.Entities.Url.Urls) > 0 {
		userUrl = u.Legacy.Entities.Url.Urls[0].ExpandedUrl
	}
	parsedUser := ParsedUser{
		UserId:         u.RestId,
		IsBlueVerified: u.IsBlueVerified,
		BirthData: struct {
			Day   int
			Month int
		}{
			Day:   u.LegacyExtendedProfile.BirthDate.Day,
			Month: u.LegacyExtendedProfile.BirthDate.Month,
		},
		IsBlockedBy:     u.Legacy.BlockedBy,
		IsBlocked:       u.Legacy.Blocking,
		CreatedAt:       createdAt,
		Description:     description,
		FavouritesCount: u.Legacy.FavouritesCount,
		FollowedBy:      u.Legacy.FollowedBy,
		FollowersCount:  u.Legacy.FollowersCount,
		IsFollowing:     u.Legacy.Following,
		FollowingCount:  u.Legacy.FriendsCount,
		Location:        u.Legacy.Location,
		MediaCount:      u.Legacy.MediaCount,
		PinnedTweetIds:  u.Legacy.PinnedTweetIdsStr,
		IsProtected:     u.Legacy.Protected,
		ScreenName:      u.Legacy.ScreenName,
		StatusesCount:   u.Legacy.StatusesCount,
		Url:             userUrl,
		IsVerified:      u.Legacy.Verified,
		VerifiedType:    u.Legacy.VerifiedType,
	}

	return &parsedUser, nil
}
