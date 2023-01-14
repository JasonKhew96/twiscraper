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
	endpoint      = "api.twitter.com"
	graphqlPath   = "/graphql"
	graphqlPrefix = scheme + "://" + endpoint + graphqlPath

	// TODO
	apiUserByRestId     = graphqlPrefix + "/mi_IjXgFyr41N9zkszPz9w/UserByRestId"
	apiUserByScreenName = graphqlPrefix + "/tgMiZwwhWR2sI0KsNsExrA/UserByScreenName"
	apiFollowing        = graphqlPrefix + "/cocC_CzoxzpwgXr3jhG7DA/Following"
	apiFollowers        = graphqlPrefix + "/KwJEsSEIHz991Ansf4Y1tQ/Followers"

	apiUserTweets = graphqlPrefix + "/whN_WW_HT--6SW2bhDcx4Q/UserTweets"
	apiUserMedia  = graphqlPrefix + "/QqRNmKWm3uTs75PCYTGkFw/UserMedia"

	apiHomeTimeline       = graphqlPrefix + "/XMoTnsLCI_a4DyvHNLSoKQ/HomeTimeline"
	apiHomeLatestTimeline = graphqlPrefix + "/kM-kIqajFOTGQLsLXv8YxQ/HomeLatestTimeline"

	apiTweetDetail = graphqlPrefix + "/d9VslTaZvKUSOh88ntOT_g/TweetDetail"

	apiFriendShipsCreate  = scheme + "://" + endpoint + "/i/api/1.1/friendships/create.json"
	apiFriendshipsDestroy = scheme + "://" + endpoint + "/i/api/1.1/friendships/destroy.json"
)
