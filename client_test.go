package websitecontacts

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

const (
	pathWebsiteContactsResponseOK         = "/WebsiteContacts/ok"
	pathWebsiteContactsResponseError      = "/WebsiteContacts/error"
	pathWebsiteContactsResponse500        = "/WebsiteContacts/500"
	pathWebsiteContactsResponsePartial1   = "/WebsiteContacts/partial"
	pathWebsiteContactsResponsePartial2   = "/WebsiteContacts/partial2"
	pathWebsiteContactsResponseUnparsable = "/WebsiteContacts/unparsable"
)

const apiKey = "at_LoremIpsumDolorSitAmetConsect"

// dummyServer is the sample of the Website Contacts API server for testing.
func dummyServer(resp, respUnparsable string, respErr string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var response string

		response = resp

		switch req.URL.Path {
		case pathWebsiteContactsResponseOK:
		case pathWebsiteContactsResponseError:
			w.WriteHeader(499)
			response = respErr
		case pathWebsiteContactsResponse500:
			w.WriteHeader(500)
			response = respUnparsable
		case pathWebsiteContactsResponsePartial1:
			response = response[:len(response)-10]
		case pathWebsiteContactsResponsePartial2:
			w.Header().Set("Content-Length", strconv.Itoa(len(response)))
			response = response[:len(response)-10]
		case pathWebsiteContactsResponseUnparsable:
			response = respUnparsable
		default:
			panic(req.URL.Path)
		}
		_, err := w.Write([]byte(response))
		if err != nil {
			panic(err)
		}
	}))

	return server
}

// newAPI returns new Website Contacts API client for testing.
func newAPI(apiServer *httptest.Server, link string) *Client {
	apiURL, err := url.Parse(apiServer.URL)
	if err != nil {
		panic(err)
	}

	apiURL.Path = link

	params := ClientParams{
		HTTPClient:       apiServer.Client(),
		WContactsBaseURL: apiURL,
	}

	return NewClient(apiKey, params)
}

// TestWebsiteContactsGet tests the Get function.
func TestWebsiteContactsGet(t *testing.T) {
	checkResultRec := func(res *WContactsResponse) bool {
		return res != nil
	}

	ctx := context.Background()

	const resp = `{"companyNames":["WHOIS API Inc","University of Maryland","Domain Research Tools"],"countryCode":"US",
				"domainName":"whoisxmlapi.com","emails":[{"description":"","email":"sales@whoisxmlapi.com"}],
				"phones":[],"postalAddresses":[],"websiteResponded":true}`

	const respUnparsable = `<?xml version="1.0" encoding="utf-8"?><>`

	const errResp = `{"code":499,"messages":"test error message"}`

	server := dummyServer(resp, respUnparsable, errResp)
	defer server.Close()

	type args struct {
		ctx     context.Context
		options string
	}

	tests := []struct {
		name    string
		path    string
		args    args
		want    bool
		wantErr string
	}{
		{
			name: "successful request",
			path: pathWebsiteContactsResponseOK,
			args: args{
				ctx:     ctx,
				options: "whoisxmlapi.com",
			},
			want:    true,
			wantErr: "",
		},
		{
			name: "non 200 status code",
			path: pathWebsiteContactsResponse500,
			args: args{
				ctx:     ctx,
				options: "whoisxmlapi.com",
			},
			want:    false,
			wantErr: "cannot parse response: invalid character '<' looking for beginning of value",
		},
		{
			name: "partial response 1",
			path: pathWebsiteContactsResponsePartial1,
			args: args{
				ctx:     ctx,
				options: "whoisxmlapi.com",
			},
			want:    false,
			wantErr: "cannot parse response: unexpected EOF",
		},
		{
			name: "partial response 2",
			path: pathWebsiteContactsResponsePartial2,
			args: args{
				ctx:     ctx,
				options: "whoisxmlapi.com",
			},
			want:    false,
			wantErr: "cannot read response: unexpected EOF",
		},
		{
			name: "could not process request",
			path: pathWebsiteContactsResponseError,
			args: args{
				ctx:     ctx,
				options: "whoisxmlapi.com",
			},
			want:    false,
			wantErr: "API error: [499] test error message",
		},
		{
			name: "unparsable response",
			path: pathWebsiteContactsResponseUnparsable,
			args: args{
				ctx:     ctx,
				options: "whoisxmlapi.com",
			},
			want:    false,
			wantErr: "cannot parse response: invalid character '<' looking for beginning of value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := newAPI(server, tt.path)

			gotRec, _, err := api.Get(tt.args.ctx, tt.args.options)
			if (err != nil || tt.wantErr != "") && (err == nil || err.Error() != tt.wantErr) {
				t.Errorf("WebsiteContacts.Get() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if tt.want {
				if !checkResultRec(gotRec) {
					t.Errorf("WebsiteContacts.Get() got = %v, expected something else", gotRec)
				}
			} else {
				if gotRec != nil {
					t.Errorf("WebsiteContacts.Get() got = %v, expected nil", gotRec)
				}
			}
		})
	}
}

// TestWebsiteContactsAPIGetRaw tests the GetRaw function.
func TestWebsiteContactsGetRaw(t *testing.T) {
	checkResultRaw := func(res []byte) bool {
		return len(res) != 0
	}

	ctx := context.Background()

	const resp = `{"companyNames":["WHOIS API Inc","University of Maryland","Domain Research Tools"],"countryCode":"US",
				"domainName":"whoisxmlapi.com","emails":[{"description":"","email":"sales@whoisxmlapi.com"}],
				"phones":[],"postalAddresses":[],"websiteResponded":true}`

	const respUnparsable = `<?xml version="1.0" encoding="utf-8"?><>`

	const errResp = `{"code":499,"messages":"test error message"}`

	server := dummyServer(resp, respUnparsable, errResp)
	defer server.Close()

	type args struct {
		ctx     context.Context
		options string
	}

	tests := []struct {
		name    string
		path    string
		args    args
		wantErr string
	}{
		{
			name: "successful request",
			path: pathWebsiteContactsResponseOK,
			args: args{
				ctx:     ctx,
				options: "whoisxmlapi.com",
			},
			wantErr: "",
		},
		{
			name: "non 200 status code",
			path: pathWebsiteContactsResponse500,
			args: args{
				ctx:     ctx,
				options: "whoisxmlapi.com",
			},
			wantErr: "API failed with status code: 500",
		},
		{
			name: "partial response 1",
			path: pathWebsiteContactsResponsePartial1,
			args: args{
				ctx:     ctx,
				options: "whoisxmlapi.com",
			},
			wantErr: "",
		},
		{
			name: "partial response 2",
			path: pathWebsiteContactsResponsePartial2,
			args: args{
				ctx:     ctx,
				options: "whoisxmlapi.com",
			},
			wantErr: "cannot read response: unexpected EOF",
		},
		{
			name: "unparsable response",
			path: pathWebsiteContactsResponseUnparsable,
			args: args{
				ctx:     ctx,
				options: "whoisxmlapi.com",
			},
			wantErr: "",
		},
		{
			name: "could not process request",
			path: pathWebsiteContactsResponseError,
			args: args{
				ctx:     ctx,
				options: "whoisxmlapi.com",
			},
			wantErr: "API failed with status code: 499",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := newAPI(server, tt.path)

			resp, err := api.GetRaw(tt.args.ctx, tt.args.options)
			if (err != nil || tt.wantErr != "") && (err == nil || err.Error() != tt.wantErr) {
				t.Errorf("WebsiteContacts.Get() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !checkResultRaw(resp.Body) {
				t.Errorf("WebsiteContacts.Get() got = %v, expected something else", string(resp.Body))
			}
		})
	}
}
