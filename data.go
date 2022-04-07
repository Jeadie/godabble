package godabble

type ArticleMetaTags struct {
	ArticleModifiedTime  int64    `json:"article_modified_time"`
	ArticlePublishedTime int64    `json:"article_published_time"`
	ArticleSection       string   `json:"article_section"`
	ArticleTags          []string `json:"article_tags"`
	Canonical            string   `json:"canonical"`
	Description          string   `json:"description"`
	Image                string   `json:"image"`
	Oembed               string   `json:"oembed"`
	Title                string   `json:"title"`
	Type                 string   `json:"type"`
}

type Category struct {
	Emoji   string `json:"emoji"`
	ID      string `json:"id"`
	Slug    string `json:"slug"`
	TagText string `json:"tag_text"`
	Title   string `json:"title"`
}

type ChartOptions struct {
	Fidelity   string `json:"fidelity"`
	IsSelected bool   `json:"is_selected"`
	Title      string `json:"title"`
}

type ChartPoint struct {
	T string  `json:"t"`
	X int64   `json:"x"`
	Y float64 `json:"y"`
}

type Comment struct {
	Comment        string `json:"comment"`
	CreatedAt      string `json:"created_at"`
	DabbleOfficial int64  `json:"dabble_official"`
	DeleteEnabled  bool   `json:"delete_enabled"`
	ID             string `json:"id"`
	IsAuthor       int64  `json:"is_author"`
	IsBearish      int64  `json:"is_bearish"`
	IsBullish      int64  `json:"is_bullish"`
	IsLiked        bool   `json:"is_liked"`
	LikeCount      int64  `json:"like_count"`
	ReplyCount     int64  `json:"reply_count"`
	ReportEnabled  bool   `json:"report_enabled"`
	UserID         string `json:"user_id"`
	UserPicture    string `json:"user_picture"`
	UserSlug       string `json:"user_slug"`
	UserUsername   string `json:"user_username"`
}

type Country struct {
	Color string  `json:"color"`
	Label string  `json:"label"`
	Value float64 `json:"value"`
}

type Holding struct {
	AllocationPercentage float64 `json:"allocation_percentage"`
	Color                string  `json:"color"`
	ID                   string  `json:"id"`
	Movement1y           float64 `json:"movement_1y"`
	Movement24h          float64 `json:"movement_24h"`
	Movement7d           float64 `json:"movement_7d"`
	Price                float64 `json:"price"`
	Price1y              float64 `json:"price_1y"`
	Price24h             float64 `json:"price_24h"`
	Price7d              float64 `json:"price_7d"`
	Slug                 string  `json:"slug"`
	Ticker               string  `json:"ticker"`
	TickerID             string  `json:"ticker_id"`
	Title                string  `json:"title"`
}

type KeyStats struct {
	Chart            []ChartPoint `json:"chart"`
	LinkedPortfolios interface{}  `json:"linked_portfolios"`
	MarketCap        int64        `json:"market_cap"`
	PeRatio          float64      `json:"pe_ratio"`
	PriceLastClose   float64      `json:"price_last_close"`
	PriceOpen        float64      `json:"price_open"`
	Range52WeekHigh  float64      `json:"range_52_week_high"`
	Range52WeekLow   float64      `json:"range_52_week_low"`
	RangeDayHigh     float64      `json:"range_day_high"`
	RangeDayLow      float64      `json:"range_day_low"`
}

type MonthlyReturn struct {
	Apr float64 `json:"Apr"`
	Aug float64 `json:"Aug"`
	Dec float64 `json:"Dec"`
	Feb float64 `json:"Feb"`
	Jan float64 `json:"Jan"`
	Jul float64 `json:"Jul"`
	Jun float64 `json:"Jun"`
	Mar float64 `json:"Mar"`
	May float64 `json:"May"`
	Nov float64 `json:"Nov"`
	Oct float64 `json:"Oct"`
	Sep float64 `json:"Sep"`
}

type News struct {
	Description   string      `json:"description"`
	ID            string      `json:"id"`
	Picture       string      `json:"picture"`
	PublishedOn   string      `json:"published_on"`
	PublisherName string      `json:"publisher_name"`
	PublisherURL  string      `json:"publisher_url"`
	Slug          string      `json:"slug"`
	Tickers       interface{} `json:"tickers"`
	Title         string      `json:"title"`
	URL           string      `json:"url"`
}

type PageMetaTags struct {
	ArticleModifiedTime  int64       `json:"article_modified_time"`
	ArticlePublishedTime int64       `json:"article_published_time"`
	ArticleSection       string      `json:"article_section"`
	ArticleTags          interface{} `json:"article_tags"`
	Canonical            string      `json:"canonical"`
	Description          string      `json:"description"`
	Image                string      `json:"image"`
	Oembed               string      `json:"oembed"`
	Title                string      `json:"title"`
	Type                 string      `json:"type"`
}

type Portfolio struct {
	Chart              interface{} `json:"chart"`
	CurrentVersionID   string      `json:"current_version_id"`
	Description        string      `json:"description"`
	Emoji              string      `json:"emoji"`
	HoldingCount       int64       `json:"holding_count"`
	ID                 string      `json:"id"`
	IsMine             bool        `json:"is_mine"`
	IsWatched          bool        `json:"is_watched"`
	Movement1y         float64     `json:"movement_1y"`
	Movement24h        float64     `json:"movement_24h"`
	Movement7d         float64     `json:"movement_7d"`
	MovementLastClose  float64     `json:"movement_last_close"`
	Picture            string      `json:"picture"`
	Price              float64     `json:"price"`
	Price1y            float64     `json:"price_1y"`
	Price24h           float64     `json:"price_24h"`
	Price7d            float64     `json:"price_7d"`
	PriceLastClose     float64     `json:"price_last_close"`
	PriceOpen          float64     `json:"price_open"`
	Slug               string      `json:"slug"`
	Title              string      `json:"title"`
	UserDabbleOfficial int64       `json:"user_dabble_official"`
	UserID             string      `json:"user_id"`
	UserPicture        string      `json:"user_picture"`
	UserSlug           string      `json:"user_slug"`
	UserUsername       string      `json:"user_username"`
	WatchCount         int64       `json:"watch_count"`
}

type PortfolioList struct {
	Portfolios []Portfolio `json:"portfolios"`
	Slug       string      `json:"slug"`
}

type RecentReturns struct {
	Averages MonthlyReturn `json:"averages"`
	Returns  []YearReturn  `json:"returns"`
}

type RelatedTags struct {
	Emoji string `json:"emoji"`
	ID    string `json:"id"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

type RelatedTickers struct {
	AllocationPercentage float64 `json:"allocation_percentage"`
	Color                string  `json:"color"`
	ID                   string  `json:"id"`
	Movement1y           float64 `json:"movement_1y"`
	Movement24h          float64 `json:"movement_24h"`
	Movement7d           float64 `json:"movement_7d"`
	Price                float64 `json:"price"`
	Price1y              float64 `json:"price_1y"`
	Price24h             float64 `json:"price_24h"`
	Price7d              float64 `json:"price_7d"`
	Slug                 string  `json:"slug"`
	Ticker               string  `json:"ticker"`
	TickerID             string  `json:"ticker_id"`
	Title                string  `json:"title"`
}

type Section struct {
	Gainers     []Portfolio `json:"gainers"`
	Losers      []Portfolio `json:"losers"`
	NewsLeft    []News      `json:"news_left"`
	NewsRight   []News      `json:"news_right"`
	Portfolios  []Portfolio `json:"portfolios"`
	SectionType string      `json:"section_type"`
	Slug        string      `json:"slug"`
	Title       string      `json:"title"`
}

type Sector struct {
	Color string  `json:"color"`
	Label string  `json:"label"`
	Value float64 `json:"value"`
}

type SmallTag struct {
	Emoji string `json:"emoji"`
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

type StockKeyStats struct {
	Chart            []ChartPoint `json:"chart"`
	DivYield         float64      `json:"div_yield"`
	LinkedPortfolios interface{}  `json:"linked_portfolios"`
	MarketCap        int64        `json:"market_cap"`
	MarketCapRank    int64        `json:"market_cap_rank"`
	PeRatio          float64      `json:"pe_ratio"`
	Price24h         float64      `json:"price_24h"`
	PriceLastClose   float64      `json:"price_last_close"`
	PriceOpen        float64      `json:"price_open"`
	Range52WeekHigh  float64      `json:"range_52_week_high"`
	Range52WeekLow   float64      `json:"range_52_week_low"`
	RangeDayHigh     float64      `json:"range_day_high"`
	RangeDayLow      float64      `json:"range_day_low"`
}

type Tag struct {
	CategoryID    string `json:"category_id"`
	CategoryTitle string `json:"category_title"`
	Emoji         string `json:"emoji"`
	ID            string `json:"id"`
	Picture       string `json:"picture"`
	Slug          string `json:"slug"`
	Title         string `json:"title"`
}

type User struct {
	Bio            string `json:"bio"`
	DabbleOfficial int64  `json:"dabble_official"`
	ID             string `json:"id"`
	Picture        string `json:"picture"`
	Slug           string `json:"slug"`
	Username       string `json:"username"`
}

type YearReturn struct {
	Months MonthlyReturn `json:"months"`
	Sum    float64       `json:"sum"`
	Year   string        `json:"year"`
}
