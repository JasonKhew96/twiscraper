package entity

// response

type Friendships struct {
	ID          int64  `json:"id"`
	IDStr       string `json:"id_str"`
	Name        string `json:"name"`
	ScreenName  string `json:"screen_name"`
	Location    string `json:"location"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Entities    struct {
		URL struct {
			Urls []struct {
				URL         string `json:"url"`
				ExpandedURL string `json:"expanded_url"`
				DisplayURL  string `json:"display_url"`
				Indices     []int  `json:"indices"`
			} `json:"urls"`
		} `json:"url"`
		Description struct {
			Urls []interface{} `json:"urls"`
		} `json:"description"`
	} `json:"entities"`
	Protected                      bool          `json:"protected"`
	FollowersCount                 int           `json:"followers_count"`
	FastFollowersCount             int           `json:"fast_followers_count"`
	NormalFollowersCount           int           `json:"normal_followers_count"`
	FriendsCount                   int           `json:"friends_count"`
	ListedCount                    int           `json:"listed_count"`
	CreatedAt                      string        `json:"created_at"`
	FavouritesCount                int           `json:"favourites_count"`
	UtcOffset                      interface{}   `json:"utc_offset"`
	TimeZone                       interface{}   `json:"time_zone"`
	GeoEnabled                     bool          `json:"geo_enabled"`
	Verified                       bool          `json:"verified"`
	StatusesCount                  int           `json:"statuses_count"`
	MediaCount                     int           `json:"media_count"`
	Lang                           interface{}   `json:"lang"`
	ContributorsEnabled            bool          `json:"contributors_enabled"`
	IsTranslator                   bool          `json:"is_translator"`
	IsTranslationEnabled           bool          `json:"is_translation_enabled"`
	ProfileBackgroundColor         string        `json:"profile_background_color"`
	ProfileBackgroundImageURL      interface{}   `json:"profile_background_image_url"`
	ProfileBackgroundImageURLHTTPS interface{}   `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool          `json:"profile_background_tile"`
	ProfileImageURL                string        `json:"profile_image_url"`
	ProfileImageURLHTTPS           string        `json:"profile_image_url_https"`
	ProfileBannerURL               string        `json:"profile_banner_url"`
	ProfileLinkColor               string        `json:"profile_link_color"`
	ProfileSidebarBorderColor      string        `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string        `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string        `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool          `json:"profile_use_background_image"`
	HasExtendedProfile             bool          `json:"has_extended_profile"`
	DefaultProfile                 bool          `json:"default_profile"`
	DefaultProfileImage            bool          `json:"default_profile_image"`
	PinnedTweetIds                 []int64       `json:"pinned_tweet_ids"`
	PinnedTweetIdsStr              []string      `json:"pinned_tweet_ids_str"`
	HasCustomTimelines             bool          `json:"has_custom_timelines"`
	CanDm                          interface{}   `json:"can_dm"`
	CanMediaTag                    bool          `json:"can_media_tag"`
	Following                      bool          `json:"following"`
	FollowRequestSent              bool          `json:"follow_request_sent"`
	Notifications                  bool          `json:"notifications"`
	Muting                         bool          `json:"muting"`
	Blocking                       bool          `json:"blocking"`
	BlockedBy                      bool          `json:"blocked_by"`
	WantRetweets                   bool          `json:"want_retweets"`
	AdvertiserAccountType          string        `json:"advertiser_account_type"`
	AdvertiserAccountServiceLevels []interface{} `json:"advertiser_account_service_levels"`
	ProfileInterstitialType        string        `json:"profile_interstitial_type"`
	BusinessProfileState           string        `json:"business_profile_state"`
	TranslatorType                 string        `json:"translator_type"`
	WithheldInCountries            []interface{} `json:"withheld_in_countries"`
	FollowedBy                     bool          `json:"followed_by"`
	ExtHasNftAvatar                bool          `json:"ext_has_nft_avatar"`
	ExtIsBlueVerified              bool          `json:"ext_is_blue_verified"`
	RequireSomeConsent             bool          `json:"require_some_consent"`
}
