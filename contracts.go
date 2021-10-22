package goopensea

func (c *Client) GetContract(asset_contract_address string) (AssetContract, error) {
	var contract AssetContract

	_, err := c.Request("GET", "/api/v1/asset_contract/"+asset_contract_address, nil, &contract)

	return contract, err
}
