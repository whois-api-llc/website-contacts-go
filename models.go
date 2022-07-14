package websitecontacts

import (
	"fmt"
)

// Meta is the metadata defined in the title and description meta tag on the main page.
type Meta struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Emails is the list of email addresses with description.
type Emails struct {
	Email       string `json:"email"`
	Description string `json:"description"`
}

// Phones is the list of phone numbers with description and call hours.
type Phones struct {
	PhoneNumber string `json:"phoneNumber"`
	Description string `json:"description"`
	CallHours   string `json:"callHours"`
}

// SocialLinks is the Facebook/Instagram/LinkedIn/Twitter URLs found on the main website's page.
type SocialLinks struct {
	Facebook  string `json:"facebook"`
	Instagram string `json:"instagram"`
	LinkedIn  string `json:"linkedIn"`
	Twitter   string `json:"twitter"`
}

// WContactsResponse is a response of Website Contacts API.
type WContactsResponse struct {
	// DomainName is a domain name.
	DomainName string `json:"domainName"`

	// WebsiteResponded Determines if the website was active during the crawling.
	WebsiteResponded bool `json:"websiteResponded"`

	// Meta is the metadata defined in the title and description meta tag on the main page.
	Meta Meta `json:"meta"`

	// CountryCode is the Website country (ISO-3166).
	CountryCode string `json:"countryCode"`

	// CompanyNames is the list of possible company names.
	CompanyNames []string `json:"companyNames"`

	// Emails is the list of email addresses with description.
	Emails []Emails `json:"emails"`

	// Phones is the list of phone numbers with description and call hours.
	Phones []Phones `json:"phones"`

	// PostalAddresses is the list of postal addresses. Every address is presented in the format of:
	// street, city, state, zip code, country. Some positions may be missing.
	PostalAddresses []string `json:"postalAddresses"`

	// SocialLinks is the Facebook/Instagram/LinkedIn/Twitter URLs found on the main website's page.
	SocialLinks SocialLinks `json:"socialLinks"`
}

// ErrorMessage is an error message.
type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"messages"`
}

// Error returns error message as a string.
func (e ErrorMessage) Error() string {
	return fmt.Sprintf("API error: [%d] %s", e.Code, e.Message)
}
