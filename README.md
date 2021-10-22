# go-opensea
Go (golang) Wrapper for the OpenSea API https://docs.opensea.io/reference/api-overview

## Installation
If using Go modules (Go version >= 11.1) simply import as needed.
```sh
go mod init github.com/yourusername/yourprojectname
```

### Older Go versions
```sh
go get github.com/rbbrown1/go-opensea
```

## Getting started
```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	goopensea "github.com/rbbrown1/go-opensea"
)

// get NFTs for a wallet address
func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(time.Second*10))
	defer cancel()

	// create opensea client object
	client := goopensea.NewClient(ctx)

    // create parameter object
	params := goopensea.GetAssetsParams{
		Owner: "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045",
		Limit: 1}

    // execute the api call
	assets, err := client.GetAssets(params)
	if err != nil {
		log.Fatalf("Error getting multiple assets: %v", err)
	}

	fmt.Printf("%+v", assets)
}
```

## Testing
```sh
go test
```

## Documentation
Documentation coming soon :)