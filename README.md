[![website-contacts-go license](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![website-contacts-go made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://pkg.go.dev/github.com/whois-api-llc/website-contacts-go)
[![website-contacts-go test](https://github.com/whois-api-llc/website-contacts-go/workflows/Test/badge.svg)](https://github.com/whois-api-llc/website-contacts-go/actions/)

# Overview

The client library for
[Website Contacts API](https://website-contacts.whoisxmlapi.com)
in Go language.

The minimum go version is 1.17.

# Installation

The library is distributed as a Go module

```bash
go get github.com/whois-api-llc/website-contacts-go
```

# Examples

Full API documentation available [here](https://website-contacts.whoisxmlapi.com/api/documentation/making-requests)

You can find all examples in `example` directory.

## Create a new client

To start making requests you need the API Key. 
You can find it on your profile page on [whoisxmlapi.com](https://whoisxmlapi.com/).
Using the API Key you can create Client.

Most users will be fine with `NewBasicClient` function. 
```go
client := websitecontacts.NewBasicClient(apiKey)
```

If you want to set custom `http.Client` to use proxy then you can use `NewClient` function.
```go
transport := &http.Transport{Proxy: http.ProxyURL(proxyUrl)}

client := websitecontacts.NewClient(apiKey, simplegeoip.ClientParams{
    HTTPClient: &http.Client{
        Transport: transport,
        Timeout:   20 * time.Second,
    },
})
```

## Make basic requests

Website Contacts API lets you get well-structured domain owner contact information, including company name and key contacts with direct-dial phone numbers, email addresses, and social media links.

```go

// Make request to get parsed Website Contacts API response for the domain name
wContactsResp, resp, err := client.Get(ctx, "whoisxmlapi.com")
if err != nil {
    log.Fatal(err)
}

log.Println(wContactsResp.DomainName)
log.Println(wContactsResp.Emails)

// Make request to get raw Website Contacts API data
resp, err := client.WContactsService.GetRaw(context.Background(), "whoisxmlapi.com")
if err != nil {
    log.Fatal(err)
}

log.Println(string(resp.Body))


```
