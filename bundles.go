package goopensea

type Bundles struct {
	Bundles []Bundle `json:"bundles"`
}

type GetBundlesParams struct {
	OnSale                 bool   `url:"on_sale,omitempty"`
	Owner                  string `url:"owner,omitempty"`
	AssetContractAddress   string `url:"asset_contract_address,omitempty"`
	AssetContractAddresses string `url:"asset_contract_addresses,omitempty"`
	TokenIds               string `url:"token_ids,omitempty"`
	Offset                 int    `url:"offset,omitempty"`
	Limit                  int    `url:"limit,omitempty"`
}

func (c *Client) GetBundles(params GetBundlesParams) (Bundles, error) {
	var bundles Bundles

	_, err := c.Request("GET", "/api/v1/bundles", params, &bundles)

	return bundles, err
}
