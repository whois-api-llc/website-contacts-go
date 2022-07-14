package example

import (
	"context"
	"errors"
	websitecontacts "github.com/whois-api-llc/website-contacts-go"
	"log"
	"strconv"
)

func GetData(apikey string) {
	client := websitecontacts.NewBasicClient(apikey)

	// Get parsed Website Contacts API response as a model instance
	wContactsResp, resp, err := client.Get(context.Background(),
		"whoisxmlapi.com",
		// this option is ignored, as the inner parser works with JSON only
		websitecontacts.OptionOutputFormat("XML"))

	if err != nil {
		// Handle error message returned by server
		var apiErr *websitecontacts.ErrorMessage
		if errors.As(err, &apiErr) {
			log.Println(apiErr.Code)
			log.Println(apiErr.Message)
		}
		log.Fatal(err)
	}

	// Then print some values from the parsed response
	log.Printf("DomainName: %s, Responded: %s, Title: %s, Emails: %s, Facebook URL: %s\n",
		wContactsResp.DomainName,
		strconv.FormatBool(wContactsResp.WebsiteResponded),
		wContactsResp.Meta.Title,
		wContactsResp.Emails,
		wContactsResp.SocialLinks.Facebook)

	log.Println("raw response is always in JSON format. Most likely you don't need it.")
	log.Printf("raw response: %s\n", string(resp.Body))
}

func GetRawData(apikey string) {
	client := websitecontacts.NewBasicClient(apikey)

	// Get raw API response
	resp, err := client.GetRaw(context.Background(),
		"whoisxmlapi.com",
		websitecontacts.OptionOutputFormat("JSON"),
		// this option causes the website contacts information to be demanded from scratch
		websitecontacts.OptionHardRefresh(1))

	if err != nil {
		// Handle error message returned by server
		log.Fatal(err)
	}

	log.Println(string(resp.Body))
}
