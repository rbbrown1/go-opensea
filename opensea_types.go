package goopensea

import "encoding/json"

type Asset struct {
	Id                      json.Number   `json:"id"`
	TokenId                 string        `json:"token_id"`
	NumSales                json.Number   `json:"num_sales"`
	ImageURL                string        `json:"image_url"`
	ImagePreviewUrl         string        `json:"image_preview_url"`
	ImageThumbnailUrl       string        `json:"image_thumbnail_url"`
	AnimationUrl            string        `json:"animation_url"`
	BackgroundColor         string        `json:"background_color"`
	Name                    string        `json:"name"`
	Description             string        `json:"description"`
	ExternalLink            string        `json:"external_link"`
	AssetContract           AssetContract `json:"asset_contract"`
	PermaLink               string        `json:"permalink"`
	Collection              Collection    `json:"collection"`
	Decimals                json.Number   `json:"decimals"`
	TokenMetaData           string        `json:"token_meta_data"`
	Owner                   Account       `json:"owner"`
	Creator                 Account       `json:"user"`
	Traits                  []Trait       `json:"traits"`
	LastSale                Sale          `json:"last_sale"`
	ListingDate             string        `json:"listing_date"`
	IsPresale               bool          `json:"is_presale"`
	TransferFeePaymentToken PaymentToken  `json:"transfer_fee_payment_token"`
	TransferFee             string        `json:"transfer_fee"`
	TopBid                  string        `json:"top_bid"`
	SupportsWyvern          bool          `json:"supports_wyvern"`
	TopOwnerships           []Ownership   `json:"top_ownerships"`
	Ownership               Ownership     `json:"ownership"`
	// RelatedAssets
	// Orders
	// Auctions
	// HighestBuyerCommitment
	// SellOrders
}

type Sale struct {
	AssetBundle    string       `json:"asset_bundle"`
	EventType      string       `json:"event_type"`
	EventTimestamp string       `json:"event_timestamp"`
	AuctionType    string       `json:"auction_type"`
	TotalPrice     string       `json:"total_price"`
	PaymentToken   PaymentToken `json:"payment_token"`
	Transaction    Transaction  `json:"transaction"`
	CreatedDate    string       `json:"created_date"`
	Quantity       string       `json:"quantity"`
}

type PaymentToken struct {
	Id       json.Number `json:"id"`
	Symbol   string      `json:"symbol"`
	Address  string      `json:"address"`
	ImageUrl string      `json:"image_url"`
	Name     string      `json:"name"`
	Decimals json.Number `json:"decimals"`
	EthPrice json.Number `json:"eth_price"`
	UsdPrice json.Number `json:"usd_price"`
}

type Trait struct {
	TraitType   string      `json:"trait_type"`
	Value       interface{} `json:"value"`
	DisplayType string      `json:"display_type"`
	MaxValue    json.Number `json:"max_value"`
	TraitCount  json.Number `json:"trait_count"`
	Order       interface{} `json:"order"`
}

type Ownership struct {
	Owner    Account     `json:"owner"`
	Quantity json.Number `json:"quantity"`
}

type Account struct {
	Address       string      `json:"address"`
	ProfileImgUrl string      `json:"profile_img_url"`
	User          interface{} `json:"user"`
	Config        string      `json:"config"`
}

type Bundle struct {
	Maker         Account       `json:"maker"`
	Slug          string        `json:"slug"`
	Assets        []Asset       `json:"assets"`
	Schemas       []string      `json:"schemas"`
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	ExternalLink  string        `json:"external_link"`
	AssetContract AssetContract `json:"asset_contract"`
	PermaLink     string        `json:"perma_link"`
	SellOrders    []Order       `json:"sell_orders"`
}

type Collection struct {
	PrimaryAssetContracts       []AssetContract `json:"primary_asset_contracts"`
	Stats                       Stats           `json:"stats"`
	BannerImageUrl              string          `json:"banner_image_url"`
	ChatUrl                     string          `json:"chat_url"`
	CreatedDate                 string          `json:"created_date"`
	DefaultToFiat               bool            `json:"default_to_fiat"`
	Description                 string          `json:"description"`
	DevBuyerFeeBasisPoints      string          `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     string          `json:"dev_seller_fee_basis_points"`
	DiscordUrl                  string          `json:"discord_url"`
	DisplayData                 DisplayData     `json:"display_data"`
	ExternalUrl                 string          `json:"external_url"`
	Featured                    bool            `json:"featured"`
	FeaturedImageUrl            string          `json:"featured_image_url"`
	Hidden                      bool            `json:"hidden"`
	SafelistRequestStatus       string          `json:"safelist_request_status"`
	ImageURL                    string          `json:"image_url"`
	IsSubjectToWhitelist        bool            `json:"is_subject_to_whitelist"`
	LargeImageUrl               string          `json:"large_image_url"`
	MediumUsername              string          `json:"medium_username"`
	Name                        string          `json:"name"`
	OnlyProxiedTransfers        bool            `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  string          `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints string          `json:"opensea_seller_fee_basis_points"`
	PayoutAddress               string          `json:"payout_address"`
	RequireEmail                bool            `json:"require_email"`
	ShortDescription            string          `json:"short_description"`
	Slug                        string          `json:"slug"`
	TelegramUrl                 string          `json:"telegram_url"`
	TwitterUsername             string          `json:"twitter_username"`
	InstagramUsername           string          `json:"instagram_username"`
	WikiUrl                     string          `json:"wiki_url"`
	OwnedAssetsCount            json.Number     `json:"owned_asset_count"`
	PaymentTokens               []PaymentToken  `json:"payment_tokens"`
	Traits                      interface{}     `json:"traits"`
}

type Stats struct {
	OneDayVolume          float64     `json:"one_day_volume"`
	OneDayChange          float64     `json:"one_day_change"`
	OneDaySales           json.Number `json:"one_day_sales"`
	OneDayAveragePrice    float64     `json:"one_day_average_price"`
	SevenDayVolume        float64     `json:"seven_day_volume"`
	SevenDayChange        float64     `json:"seven_day_change"`
	SevenDaySales         json.Number `json:"seven_day_sales"`
	SevenDayAverage       float64     `json:"seven_day_average_price"`
	ThirtyDayVolume       float64     `json:"thirty_day_volume"`
	ThirtyDayChange       float64     `json:"thirty_day_change"`
	ThirtyDaySales        json.Number `json:"thirty_day_sales"`
	ThirtyDayAveragePrice float64     `json:"thirty_day_average_price"`
	TotalVolume           float64     `json:"total_volume"`
	TotalSales            json.Number `json:"total_sales"`
	TotalSupplyCount      json.Number `json:"total_supply"`
	Count                 json.Number `json:"count"`
	NumOwners             json.Number `json:"num_owners"`
	AveragePrice          float64     `json:"average_price"`
	NumReports            json.Number `json:"num_reports"`
	MarketCap             float64     `json:"market_cap"`
	FloorPrice            float64     `json:"floor_price"`
}

type DisplayData struct {
	CardDisplayStyle string `json:"card_display_style"`
}

type AssetContract struct {
	Collection                  Collection `json:"collection"`
	Address                     string     `json:"address"`
	AddressContractType         string     `json:"asset_contract_type"`
	CreatedDate                 string     `json:"created_date"`
	Name                        string     `json:"name"`
	NftVersion                  string     `json:"nft_versiom"`
	OpenseaVersion              string     `json:"opensea_version"`
	Owner                       int        `json:"owner"`
	SchemaName                  string     `json:"schema_name"`
	Symbol                      string     `json:"symbol"`
	TotalSupply                 string     `json:"total_supply"`
	ImageURL                    string     `json:"image_url"`
	Description                 string     `json:"description"`
	ExternalLink                string     `json:"external_link"`
	DefaultToFiat               bool       `json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int        `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int        `json:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool       `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int        `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int        `json:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int        `json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int        `json:"seller_fee_basis_points"`
	PayoutAddress               string     `json:"payout_address"`
}

type AssetEvent struct {
	ApprovedAccount    Account      `json:"approved_account"`
	Asset              Asset        `json:"asset"`
	AssetBundle        string       `json:"asset_bundle"`
	AuctionType        string       `json:"auction_type"`
	BidAmount          string       `json:"bid_amount"`
	CollectionSlug     string       `json:"collection_slug"`
	ContractAddress    string       `json:"contract_address"`
	CreatedDate        string       `json:"created_date"`
	CustomEventName    string       `json:"custom_event_name"`
	DevFeePaymentEvent string       `json:"dev_fee_payment_event"`
	Duration           string       `json:"duration"`
	EndingPrice        string       `json:"ending_price"`
	EventType          string       `json:"event_type"`
	FromAccount        Account      `json:"from_account"`
	Id                 json.Number  `json:"id"`
	IsPrivate          bool         `json:"is_private"`
	OwnerAccount       Account      `json:"owner_account"`
	PaymentToken       PaymentToken `json:"payment_token"`
	Quantity           string       `json:"quantity"`
	Seller             Account      `json:"seller"`
	StartingPrice      string       `json:"starting_price"`
	ToAccount          Account      `json:"to_account"`
	TotalPrice         string       `json:"total_price"`
	Transaction        Transaction  `json:"transaction"`
	WinnerAccount      Account      `json:"winner_account"`
}

type Transaction struct {
	BlockHash        string      `json:"block_hash"`
	BlockNumber      string      `json:"block_number"`
	FromAccount      Account     `json:"from_account"`
	Id               json.Number `json:"id"`
	Timestamp        string      `json:"timestamp"`
	ToAccount        Account     `json:"to_account"`
	TransactionHash  string      `json:"transaction_hash"`
	TransactionIndex string      `json:"transaction_index"`
}

type Order struct {
	Asset                Asset        `json:"asset"`
	AssetBundle          Bundle       `json:"asset_bundle"`
	CreatedDate          string       `json:"created_date"`
	ClosingDate          string       `json:"closing_date"`
	ClosingExtendable    bool         `json:"closing_extendable"`
	ExpirationTime       int          `json:"expiration_time"`
	ListingTime          int          `json:"listing_time"`
	OrderHash            string       `json:"order_hash"`
	Metadata             interface{}  `json:"metadata"`
	Exchange             string       `json:"exchange"`
	Maker                Account      `json:"maker"`
	Taker                Account      `json:"taker"`
	CurrentPrice         string       `json:"current_price"`
	CurrentBounty        string       `json:"current_bounty"`
	BountyMultiple       string       `json:"bounty_multiple"`
	MakerRelayerFee      string       `json:"maker_relayer_fee"`
	TakerRelayerFee      string       `json:"taker_relayer_fee"`
	MakerProtocolFee     string       `json:"maker_protocol_fee"`
	TakerProtocolFee     string       `json:"taker_protocol_fee"`
	MakerReferrerFee     string       `json:"maker_referrer_fee"`
	FeeRecipient         Account      `json:"fee_recipient"`
	FeeMethod            int          `json:"fee_method"`
	Side                 int          `json:"side"`
	SaleKind             int          `json:"sale_kind"`
	Target               string       `json:"target"`
	HowToCall            int          `json:"how_to_call"`
	CallData             string       `json:"calldata"`
	ReplacementPattern   string       `json:"replacement_pattern"`
	StaticTarget         string       `json:"static_target"`
	StaticExtradata      string       `json:"static_extradata"`
	PaymentToken         string       `json:"payment_token"`
	PaymentTokenContract PaymentToken `json:"payment_token_contract"`
	BasePrice            string       `json:"base_price"`
	Extra                string       `json:"extra"`
	Quantity             string       `json:"quantity"`
	Salt                 string       `json:"salt"`
	V                    int          `json:"v"`
	R                    string       `json:"r"`
	S                    string       `json:"s"`
	ApprovedOnChain      bool         `json:"approved_on_chain"`
	Cancelled            bool         `json:"cancelled"`
	Finalized            bool         `json:"finalized"`
	MarkedInvalid        bool         `json:"marked_invalid"`
	PrefixedHash         string       `json:"prefixed_hash"`
}
