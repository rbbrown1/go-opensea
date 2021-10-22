package goopensea

import (
	"context"
	"testing"
	"time"
)

var eth_addr = "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045" // vitalik buterin
var eth_cont = "0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb" // cryptopunks

func TestGetAssets(t *testing.T) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(time.Second*10))
	defer cancel()
	client := NewClient(ctx)

	//* retreive assets
	params := GetAssetsParams{
		Owner: eth_addr}
	_, err := client.GetAssets(params)
	if err != nil {
		t.Errorf("Error getting multiple assets: %v", err)
	}
}

func TestBadParams(t *testing.T) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(time.Second*10))
	defer cancel()
	client := NewClient(ctx)

	params := GetAssetsParams{
		Owner: "1"}
	_, err := client.GetAssets(params)
	if err == nil {
		t.Error("Expected Error, got none")
	}
}

func TestGetAsset(t *testing.T) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(time.Second*10))
	defer cancel()
	client := NewClient(ctx)

	params := GetAssetParams{
		AssetContractAddress: eth_cont,
		TokenId:              "3100",
	}
	_, err := client.GetAsset(params)
	if err != nil {
		t.Errorf("Error getting a single asset: %v", err)
	}
}

func TestGetEvents(t *testing.T) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(time.Second*10))
	defer cancel()
	client := NewClient(ctx)

	params := GetEventsParams{
		AccountAddress: eth_addr}
	_, err := client.GetEvents(params)
	if err != nil {
		t.Errorf("Error getting events: %v", err)
	}
}

func TestGetContract(t *testing.T) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(time.Second*10))
	defer cancel()
	client := NewClient(ctx)

	_, err := client.GetContract(eth_cont)
	if err != nil {
		t.Errorf("Error getting a single contract: %v", err)
	}
}

func TestGetCollection(t *testing.T) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(time.Second*10))
	defer cancel()
	client := NewClient(ctx)

	params := GetCollectionParams{
		AssetOwner: eth_addr}
	_, err := client.GetCollection(params)
	if err != nil {
		t.Errorf("Error getting a collection: %v", err)
	}
}

func TestGetBundles(t *testing.T) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(time.Second*10))
	defer cancel()
	client := NewClient(ctx)

	params := GetBundlesParams{Limit: 1}
	_, err := client.GetBundles(params)
	if err != nil {
		t.Errorf("Error getting a collection: %v", err)
	}
}

func TestGetOrders(t *testing.T) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(time.Second*10))
	defer cancel()
	client := NewClient(ctx)

	params := GetOrdersParams{Owner: eth_addr}
	_, err := client.GetOrders(params)
	if err != nil {
		t.Errorf("Error getting orders: %v", err)
	}
}
