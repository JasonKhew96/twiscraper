package twiscraper

import "github.com/JasonKhew96/twiscraper/entity"

const (
	FetchFollowers fetchOptions = iota
	FetchFollowing
	FetchTweets
	FetchMedia
	FetchHomeTimeline
	FetchHomeLatestTimeline
)

type (
	fetchOptions int

	fetchFollowersFunc    func(opt fetchOptions, id string, count int, cursor string) ([]entity.ParsedUser, string, error)
	fetchTimelineFunc     func(opt fetchOptions, id string, count int, cursor string) ([]entity.ParsedTweet, string, error)
	fetchHomeTimelineFunc func(opt fetchOptions, count int, cursor string) ([]entity.ParsedTweet, string, error)
)
