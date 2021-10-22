package goopensea

type Orders struct {
	Count  int     `json:"count"`
	Orders []Order `json:"orders"`
}

type GetOrdersParams struct {
	AssetContractAddress string   `url:"asset_contract_address,omitempty"`
	PaymentTokenAddress  string   `url:"payment_token_address,omitempty"`
	Maker                string   `url:"maker,omitempty"`
	Taker                string   `url:"taker,omitempty"`
	Owner                string   `url:"owner,omitempty"`
	IsEnglish            bool     `url:"is_english,omitempty"`
	Bundled              bool     `url:"bundled,omitempty"`
	IncludeBundled       bool     `url:"include_bundled,omitempty"`
	IncludeInvalid       bool     `url:"include_invalid,omitempty"`
	ListedAfter          string   `url:"listed_after,omitempty"`
	ListedBefore         string   `url:"listed_before,omitempty"`
	TokenId              string   `url:"token_id,omitempty"`
	TokenIds             []string `url:"token_ids,omitempty"`
	Side                 int      `url:"side,omitempty"`
	SaleKind             int      `url:"sale_kind,omitempty"`
	Limit                int      `url:"limit,omitempty"`
	Offset               int      `url:"offset,omitempty"`
	OrderBy              string   `url:"order_by,omitempty"`
	OrderDirection       string   `url:"order_direction,omitempty"`
}

func (c *Client) GetOrders(params GetOrdersParams) (Orders, error) {
	var orders Orders
	_, err := c.Request("GET", "/wyvern/v1/orders", params, &orders)

	return orders, err
}
