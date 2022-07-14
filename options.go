package websitecontacts

import (
	"net/url"
	"strconv"
	"strings"
)

// Option adds parameters to the query.
type Option func(v url.Values)

var _ = []Option{
	OptionOutputFormat("JSON"),
	OptionHardRefresh(0),
}

// OptionOutputFormat sets Response output format JSON | XML. Default: JSON.
func OptionOutputFormat(outputFormat string) Option {
	return func(v url.Values) {
		v.Set("outputFormat", strings.ToUpper(outputFormat))
	}
}

// OptionHardRefresh sets the parameter for demanding the website contacts information from scratch.
func OptionHardRefresh(value int) Option {
	return func(v url.Values) {
		v.Set("hardRefresh", strconv.Itoa(value))
	}
}
