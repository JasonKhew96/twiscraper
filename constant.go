package twiscraper

const (
	DEFAULT_BEARER_TOKEN = "AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"

	SCHEME         = "https"
	ENDPOINT       = "api.twitter.com"
	GRAPHQL_PATH   = "/graphql"
	GRAPHQL_PREFIX = SCHEME + "://" + ENDPOINT + GRAPHQL_PATH

	API_USER_BY_SCREEN_NAME = GRAPHQL_PREFIX + "/tgMiZwwhWR2sI0KsNsExrA/UserByScreenName"
	API_FOLLOWING           = GRAPHQL_PREFIX + "/cocC_CzoxzpwgXr3jhG7DA/Following"
	API_FOLLOWERS           = GRAPHQL_PREFIX + "/KwJEsSEIHz991Ansf4Y1tQ/Followers"

	API_USER_TWEETS = GRAPHQL_PREFIX + "/whN_WW_HT--6SW2bhDcx4Q/UserTweets"
	API_USER_MEDIA  = GRAPHQL_PREFIX + "/QqRNmKWm3uTs75PCYTGkFw/UserMedia"

	API_HOME_TIMELINE        = GRAPHQL_PREFIX + "/XMoTnsLCI_a4DyvHNLSoKQ/HomeTimeline"
	API_HOME_LATEST_TIMELINE = GRAPHQL_PREFIX + "/kM-kIqajFOTGQLsLXv8YxQ/HomeLatestTimeline"

	API_TWEET_DETAIL = GRAPHQL_PREFIX + "/d9VslTaZvKUSOh88ntOT_g/TweetDetail"

	API_FRIENDSHIPS_CREATE  = SCHEME + "://" + ENDPOINT + "/i/api/1.1/friendships/create.json"
	API_FRIENDSHIPS_DESTROY = SCHEME + "://" + ENDPOINT + "/i/api/1.1/friendships/destroy.json"
)
