package models

// This is the root Json object.  It does not contain data that we care about though.
type RedditJsonContent struct {
	Kind string                `json:"kind"`
	Data RedditJsonContentData `json:"data"`
}

type RedditJsonContentData struct {
	After    string                      `json:"after"`
	Dist     int                         `json:"dist"`
	Modhash  string                      `json:"modhash"`
	Children []RedditJsonContentChildren `json:"children"`
}

type RedditJsonContentChildren struct {
	Kind string     `json:"kind"`
	Data RedditPost `json:"data"`
}

// RedditPost contains the information that was posted by a user.
type RedditPost struct {
	Subreddit           string          `json:"subreddit"`
	Title               string          `json:"title"`
	Content             string          `json:"selftext"`
	ContentHtml         string          `json:"selftext_html"`
	Author              string          `json:"author"`
	Permalink           string          `json:"permalink"`
	IsVideo             bool            `json:"is_video"`
	Media               RedditPostMedia `json:"media"`
	Url                 string          `json:"url"`
	UrlOverriddenByDest string          `json:"url_overridden_by_dest"`

	Thumbnail string `json:"thumbnail"`
}

// RedditPostMedia defines if the post contains a video that is hosted on Reddit.
type RedditPostMedia struct {
	RedditVideo RedditPostMediaRedditVideo `json:"reddit_video"`
}

// RedditVideo contains information about the video in the post.
type RedditPostMediaRedditVideo struct {
	BitrateKbps       int    `json:"bitrate_kpbs"`
	FallBackUrl       string `json:"fallback_url"`
	Height            int    `json:"height"`
	Width             int    `json:"width"`
	ScrubberMediaUrl  string `json:"scrubber_media_url"`
	DashUrl           string `json:"dash_url"`
	Duration          int    `json:"duration"`
	HlsUrl            string `json:"hls_url"`
	IsGif             bool   `json:"is_gif"`
	TranscodingStatus string `json:"transcoding_status"`
}
