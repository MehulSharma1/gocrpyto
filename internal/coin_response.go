package internal

import "time"

type Currency struct {
	Aed  float64 `json:"aed"`
	Ars  float64 `json:"ars"`
	Aud  float64 `json:"aud"`
	Bch  float64 `json:"bch"`
	Bdt  float64 `json:"bdt"`
	Bhd  float64 `json:"bhd"`
	Bmd  float64 `json:"bmd"`
	Bnb  float64 `json:"bnb"`
	Brl  float64 `json:"brl"`
	Btc  float64 `json:"btc"`
	Cad  float64 `json:"cad"`
	Chf  float64 `json:"chf"`
	Clp  float64 `json:"clp"`
	Cny  float64 `json:"cny"`
	Czk  float64 `json:"czk"`
	Dkk  float64 `json:"dkk"`
	Dot  float64 `json:"dot"`
	Eos  float64 `json:"eos"`
	Eth  float64 `json:"eth"`
	Eur  float64 `json:"eur"`
	Gbp  float64 `json:"gbp"`
	Hkd  float64 `json:"hkd"`
	Huf  float64 `json:"huf"`
	Idr  float64 `json:"idr"`
	Ils  float64 `json:"ils"`
	Inr  float64 `json:"inr"`
	Jpy  float64 `json:"jpy"`
	Krw  float64 `json:"krw"`
	Kwd  float64 `json:"kwd"`
	Lkr  float64 `json:"lkr"`
	Ltc  float64 `json:"ltc"`
	Mmk  float64 `json:"mmk"`
	Mxn  float64 `json:"mxn"`
	Myr  float64 `json:"myr"`
	Ngn  float64 `json:"ngn"`
	Nok  float64 `json:"nok"`
	Nzd  float64 `json:"nzd"`
	Php  float64 `json:"php"`
	Pkr  float64 `json:"pkr"`
	Pln  float64 `json:"pln"`
	Rub  float64 `json:"rub"`
	Sar  float64 `json:"sar"`
	Sek  float64 `json:"sek"`
	Sgd  float64 `json:"sgd"`
	Thb  float64 `json:"thb"`
	Try  float64 `json:"try"`
	Twd  float64 `json:"twd"`
	Uah  float64 `json:"uah"`
	Usd  float64 `json:"usd"`
	Vef  float64 `json:"vef"`
	Vnd  float64 `json:"vnd"`
	Xag  float64 `json:"xag"`
	Xau  float64 `json:"xau"`
	Xdr  float64 `json:"xdr"`
	Xlm  float64 `json:"xlm"`
	Xrp  float64 `json:"xrp"`
	Yfi  float64 `json:"yfi"`
	Zar  float64 `json:"zar"`
	Bits float64 `json:"bits"`
	Link float64 `json:"link"`
	Sats float64 `json:"sats"`
}
type CoinResponse struct {
	ID              string      `json:"id"`
	Symbol          string      `json:"symbol"`
	Name            string      `json:"name"`
	AssetPlatformID interface{} `json:"asset_platform_id"`
	Platforms       struct {
		string            `json:""`
		BinanceSmartChain string `json:"binance-smart-chain"`
	} `json:"platforms"`
	BlockTimeInMinutes int           `json:"block_time_in_minutes"`
	HashingAlgorithm   string        `json:"hashing_algorithm"`
	Categories         []string      `json:"categories"`
	PublicNotice       interface{}   `json:"public_notice"`
	AdditionalNotices  []interface{} `json:"additional_notices"`
	Description        struct {
		En string `json:"en"`
	} `json:"description"`
	Links struct {
		Homepage                    []string `json:"homepage"`
		BlockchainSite              []string `json:"blockchain_site"`
		OfficialForumURL            []string `json:"official_forum_url"`
		ChatURL                     []string `json:"chat_url"`
		AnnouncementURL             []string `json:"announcement_url"`
		TwitterScreenName           string   `json:"twitter_screen_name"`
		FacebookUsername            string   `json:"facebook_username"`
		BitcointalkThreadIdentifier int      `json:"bitcointalk_thread_identifier"`
		TelegramChannelIdentifier   string   `json:"telegram_channel_identifier"`
		SubredditURL                string   `json:"subreddit_url"`
		ReposURL                    struct {
			Github    []string      `json:"github"`
			Bitbucket []interface{} `json:"bitbucket"`
		} `json:"repos_url"`
	} `json:"links"`
	Image struct {
		Thumb string `json:"thumb"`
		Small string `json:"small"`
		Large string `json:"large"`
	} `json:"image"`
	CountryOrigin                string  `json:"country_origin"`
	GenesisDate                  string  `json:"genesis_date"`
	SentimentVotesUpPercentage   float64 `json:"sentiment_votes_up_percentage"`
	SentimentVotesDownPercentage float64 `json:"sentiment_votes_down_percentage"`
	MarketCapRank                int     `json:"market_cap_rank"`
	CoingeckoRank                int     `json:"coingecko_rank"`
	CoingeckoScore               float64 `json:"coingecko_score"`
	DeveloperScore               float64 `json:"developer_score"`
	CommunityScore               float64 `json:"community_score"`
	LiquidityScore               float64 `json:"liquidity_score"`
	PublicInterestScore          float64 `json:"public_interest_score"`
	MarketData                   struct {
		CurrentPrice          Currency    `json:"current_price"`
		TotalValueLocked      interface{} `json:"total_value_locked"`
		McapToTvlRatio        interface{} `json:"mcap_to_tvl_ratio"`
		FdvToTvlRatio         interface{} `json:"fdv_to_tvl_ratio"`
		Roi                   interface{} `json:"roi"`
		Ath                   Currency    `json:"ath"`
		AthChangePercentage   Currency    `json:"ath_change_percentage"`
		AthDate               Currency    `json:"ath_date"`
		Atl                   Currency    `json:"atl"`
		AtlChangePercentage   Currency    `json:"atl_date"`
		MarketCap             Currency    `json:"market_cap"`
		MarketCapRank         int         `json:"market_cap_rank"`
		FullyDilutedValuation struct {
		} `json:"fully_diluted_valuation"`
		TotalVolume                            Currency    `json:"total_volume"`
		High24H                                Currency    `json:"high_24h"`
		Low24H                                 Currency    `json:"low_24h"`
		PriceChange24H                         float64     `json:"price_change_24h"`
		PriceChangePercentage24H               float64     `json:"price_change_percentage_24h"`
		PriceChangePercentage7D                float64     `json:"price_change_percentage_7d"`
		PriceChangePercentage14D               float64     `json:"price_change_percentage_14d"`
		PriceChangePercentage30D               float64     `json:"price_change_percentage_30d"`
		PriceChangePercentage60D               float64     `json:"price_change_percentage_60d"`
		PriceChangePercentage200D              float64     `json:"price_change_percentage_200d"`
		PriceChangePercentage1Y                float64     `json:"price_change_percentage_1y"`
		MarketCapChange24H                     int         `json:"market_cap_change_24h"`
		MarketCapChangePercentage24H           float64     `json:"market_cap_change_percentage_24h"`
		PriceChange24HInCurrency               Currency    `json:"price_change_24h_in_currency"`
		PriceChangePercentage1HInCurrency      Currency    `json:"price_change_percentage_1h_in_currency"`
		PriceChangePercentage24HInCurrency     Currency    `json:"price_change_percentage_24h_in_currency"`
		PriceChangePercentage7DInCurrency      Currency    `json:"price_change_percentage_7d_in_currency"`
		PriceChangePercentage14DInCurrency     Currency    `json:"price_change_percentage_14d_in_currency"`
		PriceChangePercentage30DInCurrency     Currency    `json:"price_change_percentage_30d_in_currency"`
		PriceChangePercentage60DInCurrency     Currency    `json:"price_change_percentage_60d_in_currency"`
		PriceChangePercentage200DInCurrency    Currency    `json:"price_change_percentage_200d_in_currency"`
		PriceChangePercentage1YInCurrency      Currency    `json:"price_change_percentage_1y_in_currency"`
		MarketCapChange24HInCurrency           Currency    `json:"market_cap_change_24h_in_currency"`
		MarketCapChangePercentage24HInCurrency Currency    `json:"market_cap_change_percentage_24h_in_currency"`
		TotalSupply                            interface{} `json:"total_supply"`
		MaxSupply                              interface{} `json:"max_supply"`
		CirculatingSupply                      float64     `json:"circulating_supply"`
		LastUpdated                            time.Time   `json:"last_updated"`
	} `json:"market_data"`
	CommunityData struct {
		FacebookLikes            interface{} `json:"facebook_likes"`
		TwitterFollowers         int         `json:"twitter_followers"`
		RedditAveragePosts48H    float64     `json:"reddit_average_posts_48h"`
		RedditAverageComments48H float64     `json:"reddit_average_comments_48h"`
		RedditSubscribers        int         `json:"reddit_subscribers"`
		RedditAccountsActive48H  int         `json:"reddit_accounts_active_48h"`
		TelegramChannelUserCount interface{} `json:"telegram_channel_user_count"`
	} `json:"community_data"`
	DeveloperData struct {
		Forks                        int `json:"forks"`
		Stars                        int `json:"stars"`
		Subscribers                  int `json:"subscribers"`
		TotalIssues                  int `json:"total_issues"`
		ClosedIssues                 int `json:"closed_issues"`
		PullRequestsMerged           int `json:"pull_requests_merged"`
		PullRequestContributors      int `json:"pull_request_contributors"`
		CodeAdditionsDeletions4Weeks struct {
			Additions int `json:"additions"`
			Deletions int `json:"deletions"`
		} `json:"code_additions_deletions_4_weeks"`
		CommitCount4Weeks              int   `json:"commit_count_4_weeks"`
		Last4WeeksCommitActivitySeries []int `json:"last_4_weeks_commit_activity_series"`
	} `json:"developer_data"`
	PublicInterestStats struct {
		AlexaRank   int         `json:"alexa_rank"`
		BingMatches interface{} `json:"bing_matches"`
	} `json:"public_interest_stats"`
	StatusUpdates []interface{} `json:"status_updates"`
	LastUpdated   time.Time     `json:"last_updated"`
	Tickers       []struct {
		Base   string `json:"base"`
		Target string `json:"target"`
		Market struct {
			Name                string `json:"name"`
			Identifier          string `json:"identifier"`
			HasTradingIncentive bool   `json:"has_trading_incentive"`
		} `json:"market"`
		Last          float64 `json:"last"`
		Volume        float64 `json:"volume"`
		ConvertedLast struct {
			Btc float64 `json:"btc"`
			Eth float64 `json:"eth"`
			Usd float64 `json:"usd"`
		} `json:"converted_last"`
		ConvertedVolume struct {
			Btc int `json:"btc"`
			Eth int `json:"eth"`
			Usd int `json:"usd"`
		} `json:"converted_volume"`
		TrustScore             string      `json:"trust_score"`
		BidAskSpreadPercentage float64     `json:"bid_ask_spread_percentage"`
		Timestamp              time.Time   `json:"timestamp"`
		LastTradedAt           time.Time   `json:"last_traded_at"`
		LastFetchAt            time.Time   `json:"last_fetch_at"`
		IsAnomaly              bool        `json:"is_anomaly"`
		IsStale                bool        `json:"is_stale"`
		TradeURL               string      `json:"trade_url"`
		TokenInfoURL           interface{} `json:"token_info_url"`
		CoinID                 string      `json:"coin_id"`
		TargetCoinID           string      `json:"target_coin_id,omitempty"`
	} `json:"tickers"`
}
