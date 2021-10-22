package goopensea

type Assets struct {
	Assets []Asset `json:"assets"`
}

type GetAssetsParams struct {
	Owner                  string   `url:"owner,omitempty"`
	TokenIds               []string `url:"token_ids,omitempty"`
	AssetContractAddress   string   `url:"asset_contract_address,omitempty"`
	AssetContractAddresses []string `url:"asset_contract_addresses,omitempty"`
	OrderBy                string   `url:"order_by,omitempty"`
	OrderDirection         string   `url:"order_direction,omitempty"`
	Offset                 int      `url:"offset,omitempty"`
	Limit                  int      `url:"limit,omitempty"`
	Collection             string   `url:"collection,omitempty"`
}

type GetAssetParams struct {
	TokenId              string `url:"token_id,omitempty" binding:"required"`
	AssetContractAddress string `url:"asset_contract_address,omitempty" binding:"required"`
}

// retrieve multiple assets, such as all NFTs owned by a wallet
func (c *Client) GetAssets(params GetAssetsParams) (Assets, error) {
	var assets Assets

	_, err := c.Request("GET", "/api/v1/assets", params, &assets)

	return assets, err
}

// retrieve a single asset
func (c *Client) GetAsset(params GetAssetParams) (Asset, error) {
	var asset Asset

	path := "/api/v1/asset/" + params.AssetContractAddress + "/" + params.TokenId + "/"
	_, err := c.Request("GET", path, nil, &asset)

	return asset, err
}
