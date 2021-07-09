# sumologic-go-sdk

`sumolog-go-sdk` is the Sumo Logic SDK for the Go programming language.

The SDK requires a minimum version of `Go 1.16`.

## Region Endpoints

When using the `service/cip` SDK you will need to specify the region to send the API request to. These regional endpoints can 
be found on the [Sumo Logic Endpoints and Firewall Security](https://help.sumologic.com/APIs/General-API-Information/Sumo-Logic-Endpoints-and-Firewall-Security) page.

## Getting started

To get started working with the SDK setup your project for Go modules, and retrieve the SDK dependencies
with `go get`. This example shows how you can use the SDK to make an API request to list the personal folder of the authenticated user.

###### Add SDK Dependencies

You only need to add the dependencies that your project requires. For example if you only need to to interact with Sumo Logic 
Core Intelligence Platform (CIP) resources then you only need to use `service/cip`.

```
go get github.com/wizedkyle/sumologic-go-sdk/service/cip
```

###### Write Code

In your preferred code editor add the following content to `main.go`

```go
package main

import (
	"fmt"
	"github.com/wizedkyle/sumologic-go-sdk/service/cip"
)

func main() {
	client := cip.APIClient{
		Cfg: &cip.Configuration{
			Authentication: cip.BasicAuth{
				AccessId:  "<accessId>",
				AccessKey: "<accessKey>",
			},
			BasePath:   "https://api.<regioncode>.sumologic.com/api"
			HTTPClient: &http.Client{},
		},
	}
	apiResponse, httpResponse, errorResponse := client.GetPersonalFolder()
	if errorResponse != nil {
		fmt.Println(errorResponse.Error())
	} else {
		if httpResponse.StatusCode == 200 {
			fmt.Println(apiResponse.Name)
        }
    }
}
```
## Compatibility

You can find a list of API endpoints and sources that the sumologic-go-sdk supports in the [supported capabilities](COMPATIBILITY.md) documentation.

## Contributing

The Sumo Logic Go SDK uses GitHub [Issues](https://github.com/wizedkyle/sumologic-go-sdk/issues) to report and track 
issues with the SDK. If you have found a bug, identified an area of improvement or want a new feature add please create or 
upvote an existing issue.

You can open pull requests for the Sumo Logic Go SDK for fixes or additions, these pull requests will be reviewed by the maintainers
of `sumologic-go-sdk` before being merged.
