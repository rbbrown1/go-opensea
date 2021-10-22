package goopensea

type Collections struct {
	Collections []Collection `json:"collections"`
}

type GetCollectionParams struct {
	AssetOwner string `url:"asset_owner,omitempty"`
	Offset     int    `url:"offset,omitempty"`
	Limit      int    `url:"limit,omitempty"`
}

func (c *Client) GetCollection(params GetCollectionParams) (Collections, error) {
	var collections Collections

	_, err := c.Request("GET", "/api/v1/collections", nil, &collections)

	return collections, err
}
