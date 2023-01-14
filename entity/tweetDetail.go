package entity

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// request

type TweetDetailVariables struct {
	FocalTweetID                           string `json:"focalTweetId"`
	WithRuxInjections                      bool   `json:"with_rux_injections"`
	IncludePromotedContent                 bool   `json:"includePromotedContent"`
	WithCommunity                          bool   `json:"withCommunity"`
	WithQuickPromoteEligibilityTweetFields bool   `json:"withQuickPromoteEligibilityTweetFields"`
	WithBirdwatchNotes                     bool   `json:"withBirdwatchNotes"`
	WithSuperFollowsUserFields             bool   `json:"withSuperFollowsUserFields"`
	WithDownvotePerspective                bool   `json:"withDownvotePerspective"`
	WithReactionsMetadata                  bool   `json:"withReactionsMetadata"`
	WithReactionsPerspective               bool   `json:"withReactionsPerspective"`
	WithSuperFollowsTweetFields            bool   `json:"withSuperFollowsTweetFields"`
	WithVoice                              bool   `json:"withVoice"`
	WithV2Timeline                         bool   `json:"withV2Timeline"`
}

type TweetDetailFeatures struct {
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

// response

type TweetDetailResponse struct {
	Errors []TwitterError `json:"errors"`
	Data struct {
		ThreadedConversationWithInjectionsV2 struct {
			Instructions []TimelineInstruction `json:"instructions"`
		} `json:"threaded_conversation_with_injections_v2"`
	} `json:"data"`
}

// function

func NewTweetDetailVariables(focalTweetID string) TweetDetailVariables {
	return TweetDetailVariables{
		FocalTweetID:                           focalTweetID,
		WithRuxInjections:                      false,
		IncludePromotedContent:                 true,
		WithCommunity:                          true,
		WithQuickPromoteEligibilityTweetFields: true,
		WithBirdwatchNotes:                     true,
		WithSuperFollowsUserFields:             true,
		WithDownvotePerspective:                false,
		WithReactionsMetadata:                  false,
		WithReactionsPerspective:               false,
		WithSuperFollowsTweetFields:            true,
		WithVoice:                              true,
		WithV2Timeline:                         true,
	}
}

func NewTweetDetailFeatures() TweetDetailFeatures {
	return TweetDetailFeatures{
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
		ResponsiveWebEnhanceCardsEnabled:                               false,
	}
}

func NewTweetDetailParams(focalTweetId string) (url.Values, error) {
	variables, err := json.Marshal(NewTweetDetailVariables(focalTweetId))
	if err != nil {
		return nil, fmt.Errorf("failed to marshal variables: %v", err)
	}
	features, err := json.Marshal(NewTweetDetailFeatures())
	if err != nil {
		return nil, fmt.Errorf("failed to marshal features: %v", err)
	}
	return url.Values{
		"variables": {string(variables)},
		"features":  {string(features)},
	}, nil
}
