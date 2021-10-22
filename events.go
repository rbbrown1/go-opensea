package goopensea

type Events struct {
	AssetEvents []AssetEvent `json:"asset_events"`
}

type GetEventsParams struct {
	AssetContractAddress string `url:"asset_contract_address,omitempty"`
	CollectionSlug       string `url:"collection_slug,omitempty"`
	TokenId              int    `url:"token_id,omitempty"`
	AccountAddress       string `url:"account_address,omitempty"`
	EventType            string `url:"event_type,omitempty"`
	OnlyOpensea          bool   `url:"only_opensea,omitempty"`
	AuctionType          string `url:"auction_type,omitempty"`
	Offset               int    `url:"offset,omitempty"`
	Limit                int    `url:"limit,omitempty"`
	OccuredBefore        string `url:"occurred_before,omitempty"`
	OccuredAfter         string `url:"occurred_after,omitempty"`
}

func (c *Client) GetEvents(params GetEventsParams) (Events, error) {
	var events Events

	_, err := c.Request("GET", "/api/v1/events", params, &events)

	return events, err
}
