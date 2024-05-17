package twiscraper

import (
	"time"
)

const (
	defaultDelay         = 1 * time.Second
	defaultClientTimeout = 10 * time.Second
	defaultBearerToken   = "AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"
	defaultUserAgent     = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36"

	scheme        = "https"
	endpoint      = "x.com"
	graphqlPath   = "/i/api/graphql"
	graphqlPrefix = scheme + "://" + endpoint + graphqlPath

	// TODO
	apiUserByRestId     = graphqlPrefix + "/gUIQEk2xDGzQTX8Ii0Yesw/UserByRestId"
	apiUserByScreenName = graphqlPrefix + "/rePnxwe9LZ51nQ7Sn_xN_A/UserByScreenName"
	apiFollowing        = graphqlPrefix + "/wh5eBj-w6PPSzTSbgrlHzw/Following"
	apiFollowers        = graphqlPrefix + "/EcJ6iHzKpwDjpC0Dm1Gkhw/Followers"

	apiUserTweets = graphqlPrefix + "/rCpYpqplOq3UJ2p6Oxy3tw/UserTweets"
	apiUserMedia  = graphqlPrefix + "/ghc-7mU9EvRC54PiccAsCA/UserMedia"

	apiHomeTimeline       = graphqlPrefix + "/6VUR2qFhg6jw55JEvJEmmA/HomeTimeline"
	apiHomeLatestTimeline = graphqlPrefix + "/vxwgV-TdXnjj9AscrP0mTA/HomeLatestTimeline"

	apiTweetDetail = graphqlPrefix + "/VaihYjIIeVg4gfvwMgQsUA/TweetDetail"

	apiFriendShipsCreate  = scheme + "://" + endpoint + "/i/api/1.1/friendships/create.json"
	apiFriendshipsDestroy = scheme + "://" + endpoint + "/i/api/1.1/friendships/destroy.json"
)
