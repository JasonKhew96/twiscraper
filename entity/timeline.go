package entity

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type TimelineInstruction = json.RawMessage

type TimelineEntry = json.RawMessage

type TimelineInstructionAddEntries struct {
	Type    string            `json:"type"`
	Entries []TimelineEntry `json:"entries"`
}

// request

type HomeTimelineVariables struct {
	Cursor                      string `json:"cursor,omitempty"`
	Count                       int    `json:"count"`
	IncludePromotedContent      bool   `json:"includePromotedContent"`
	LatestControlAvailable      bool   `json:"latestControlAvailable"`
	WithCommunity               bool   `json:"withCommunity"`
	WithSuperFollowsUserFields  bool   `json:"withSuperFollowsUserFields"`
	WithDownvotePerspective     bool   `json:"withDownvotePerspective"`
	WithReactionsMetadata       bool   `json:"withReactionsMetadata"`
	WithReactionsPerspective    bool   `json:"withReactionsPerspective"`
	WithSuperFollowsTweetFields bool   `json:"withSuperFollowsTweetFields"`
}

type HomeLatestTimelineVariables struct {
	Cursor                      string `json:"cursor,omitempty"`
	Count                       int    `json:"count"`
	IncludePromotedContent      bool   `json:"includePromotedContent"`
	LatestControlAvailable      bool   `json:"latestControlAvailable"`
	WithSuperFollowsUserFields  bool   `json:"withSuperFollowsUserFields"`
	WithDownvotePerspective     bool   `json:"withDownvotePerspective"`
	WithReactionsMetadata       bool   `json:"withReactionsMetadata"`
	WithReactionsPerspective    bool   `json:"withReactionsPerspective"`
	WithSuperFollowsTweetFields bool   `json:"withSuperFollowsTweetFields"`
}

// response

type HomeTimelineResponse struct {
	Errors []TwitterError `json:"errors"`
	Data struct {
		Home struct {
			HomeTimelineUrt struct {
				Instructions    []TimelineInstruction `json:"instructions"`
				ResponseObjects interface{}       `json:"responseObjects"`
			} `json:"home_timeline_urt"`
		} `json:"home"`
	} `json:"data"`
}

// function

func NewHomeTimelineVariables(count int, cursor string) HomeTimelineVariables {
	v := HomeTimelineVariables{
		Count:                       count,
		IncludePromotedContent:      true,
		LatestControlAvailable:      true,
		WithCommunity:               true,
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

func NewHomeLatestTimelineVariables(count int, cursor string) HomeLatestTimelineVariables {
	v := HomeLatestTimelineVariables{
		Count:                       count,
		IncludePromotedContent:      true,
		LatestControlAvailable:      true,
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

func NewHomeTimelineParams(count int, cursor string) (url.Values, error) {
	variables, err := json.Marshal(NewHomeTimelineVariables(count, cursor))
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

func NewHomeLatestTimelineParams(count int, cursor string) (url.Values, error) {
	variables, err := json.Marshal(NewHomeLatestTimelineVariables(count, cursor))
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
