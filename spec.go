package godabble

type Categories struct {
	Categories   []Category `json:"categories"`
	ErrorMessage string     `json:"error_message"`
	Status       string     `json:"status"`
}
type CategoryPage struct {
	Description     string        `json:"description"`
	Emoji           string        `json:"emoji"`
	ErrorMessage    string        `json:"error_message"`
	ID              string        `json:"id"`
	MetaTags        PageMetaTags  `json:"meta_tags"`
	OtherCategories []Category    `json:"other_categories"`
	Portfolios      []Portfolio   `json:"portfolios"`
	Slug            string        `json:"slug"`
	Status          string        `json:"status"`
	Tags            []RelatedTags `json:"tags"`
	Title           string        `json:"title"`
}

type Comments struct {
	Comments         []Comment `json:"comments"`
	ErrorMessage     string    `json:"error_message"`
	HasMore          bool      `json:"has_more"`
	NextCursor       int64     `json:"next_cursor"`
	ReplyEnabled     bool      `json:"reply_enabled"`
	SentimentEnabled bool      `json:"sentiment_enabled"`
	Status           string    `json:"status"`
}

type Home struct {
	MetaTags PageMetaTags `json:"meta_tags"`
	Sections []Section    `json:"sections"`
}

type Me struct {
	EmailNotificationsEnabled int64       `json:"email_notifications_enabled"`
	ErrorMessage              string      `json:"error_message"`
	ID                        string      `json:"id"`
	IsAuthenticated           bool        `json:"is_authenticated"`
	Name                      string      `json:"name"`
	Picture                   string      `json:"picture"`
	PortfoliosPrivate         interface{} `json:"portfolios_private"`
	PortfoliosPublished       interface{} `json:"portfolios_published"`
	PortfoliosWatched         interface{} `json:"portfolios_watched"`
	PushNotificationsEnabled  int64       `json:"push_notifications_enabled"`
	Slug                      string      `json:"slug"`
	Status                    string      `json:"status"`
	Timezone                  string      `json:"timezone"`
	Username                  string      `json:"username"`
	WeeklyEmailEnabled        int64       `json:"weekly_email_enabled"`
}

type NewsPage struct {
	Description       string           `json:"description"`
	ErrorMessage      string           `json:"error_message"`
	ID                string           `json:"id"`
	IsAuthenticated   bool             `json:"is_authenticated"`
	LatestNews        []News           `json:"latest_news"`
	MetaTags          ArticleMetaTags  `json:"meta_tags"`
	Picture           string           `json:"picture"`
	PublishedOn       string           `json:"published_on"`
	PublisherName     string           `json:"publisher_name"`
	PublisherURL      string           `json:"publisher_url"`
	RelatedPortfolios interface{}      `json:"related_portfolios"`
	RelatedTags       []RelatedTags    `json:"related_tags"`
	RelatedTickers    []RelatedTickers `json:"related_tickers"`
	Slug              string           `json:"slug"`
	Status            string           `json:"status"`
	Tickers           interface{}      `json:"tickers"`
	Title             string           `json:"title"`
	URL               string           `json:"url"`
}

type PortfolioPage struct {
	AllocationMethod string         `json:"allocation_method"`
	Chart            interface{}    `json:"chart"`
	ChartOptions     []ChartOptions `json:"chart_options"`
	Countries        []Country      `json:"countries"`
	Description      string         `json:"description"`
	Emoji            string         `json:"emoji"`
	ErrorMessage     string         `json:"error_message"`
	HoldingCount     int64          `json:"holding_count"`
	Holdings         []Holding      `json:"holdings"`
	ID               string         `json:"id"`
	IsAuthenticated  bool           `json:"is_authenticated"`
	IsMine           bool           `json:"is_mine"`
	IsPublished      int64          `json:"is_published"`
	IsQueuedForArt   int64          `json:"is_queued_for_art"`
	IsWatched        bool           `json:"is_watched"`
	KeyStats         KeyStats       `json:"key_stats"`
	MetaTags         PageMetaTags   `json:"meta_tags"`
	Movement1y       float64        `json:"movement_1y"`
	Movement24h      float64        `json:"movement_24h"`
	Movement7d       float64        `json:"movement_7d"`
	News             []News         `json:"news"`
	Picture          string         `json:"picture"`
	Price            float64        `json:"price"`
	Price1y          float64        `json:"price_1y"`
	Price24h         float64        `json:"price_24h"`
	Price7d          float64        `json:"price_7d"`
	PublishedAt      string         `json:"published_at"`
	RecentReturns    RecentReturns  `json:"recent_returns"`
	Sectors          []Sector       `json:"sectors"`
	Slug             string         `json:"slug"`
	Status           string         `json:"status"`
	Tags             []Tag          `json:"tags"`
	Title            string         `json:"title"`
	User             User           `json:"user"`
	WatchCount       int64          `json:"watch_count"`
}

type Rankings struct {
	Gainers  PortfolioList `json:"gainers"`
	Losers   PortfolioList `json:"losers"`
	MetaTags PageMetaTags  `json:"meta_tags"`
}

type Stock struct {
	Ceo                 string         `json:"ceo"`
	CeoPicture          string         `json:"ceo_picture"`
	Chart               interface{}    `json:"chart"`
	ChartOptions        []ChartOptions `json:"chart_options"`
	Color               string         `json:"color"`
	Description         string         `json:"description"`
	ErrorMessage        string         `json:"error_message"`
	ExchangeDescription string         `json:"exchange_description"`
	ID                  string         `json:"id"`
	KeyStats            KeyStats       `json:"key_stats"`
	LocaleDescription   string         `json:"locale_description"`
	MetaTags            PageMetaTags   `json:"meta_tags"`
	Movement1y          float64        `json:"movement_1y"`
	Movement24h         float64        `json:"movement_24h"`
	Movement7d          float64        `json:"movement_7d"`
	News                []News         `json:"news"`
	Portfolios          []Portfolio    `json:"portfolios"`
	Price               float64        `json:"price"`
	Price1y             float64        `json:"price_1y"`
	Price24h            float64        `json:"price_24h"`
	Price7d             float64        `json:"price_7d"`
	RecentReturns       RecentReturns  `json:"recent_returns"`
	Slug                string         `json:"slug"`
	Status              string         `json:"status"`
	Tags                []SmallTag     `json:"tags"`
	Ticker              string         `json:"ticker"`
	Title               string         `json:"title"`
	Type                string         `json:"type"`
}

type Tags struct {
	Tags []Tag `json:"tags"`
}

type UserPage struct {
	Bio                 string       `json:"bio"`
	ErrorMessage        string       `json:"error_message"`
	FollowerCount       int64        `json:"follower_count"`
	Followers           []Follower   `json:"followers"`
	Following           []Follower   `json:"following"`
	FollowingCount      int64        `json:"following_count"`
	ID                  string       `json:"id"`
	IsAuthenticated     bool         `json:"is_authenticated"`
	IsFollowing         bool         `json:"is_following"`
	IsMine              bool         `json:"is_mine"`
	MetaTags            PageMetaTags `json:"meta_tags"`
	Picture             string       `json:"picture"`
	PortfolioCount      int64        `json:"portfolio_count"`
	PortfoliosPrivate   interface{}  `json:"portfolios_private"`
	PortfoliosPublished []Portfolio  `json:"portfolios_published"`
	PortfoliosWatched   []Portfolio  `json:"portfolios_watched"`
	Slug                string       `json:"slug"`
	Status              string       `json:"status"`
	Username            string       `json:"username"`
}
