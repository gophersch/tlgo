# tlgo

A Go package to query TL rest API.

## Installation

```shell
go get github.com/gophersch/tlgo
```

## Usage

```go
package main

import (
    "github.com/gophersch/tlgo"
    "fmt"
)

func main() {

    // Get a Client instance
    client := tlgo.NewClient()
	lines, err := client.ListLines()
	if err != nil {
		fmt.Printf("Can not get line lists: %s\n", err)
	}

	_, err = client.ListRoutes(lines[0])
	if err != nil {
		fmt.Printf("Can not get routes list: %s\n", err)
	}

	details, err := client.ListStopDeparturesFromIDs("1970329131941987", "11822125115506799", time.Now(), false)
	if err != nil {
		fmt.Printf("Can not get details list: %s\n", err)
	}

	fmt.Printf("Details: %+v\n", details)
```