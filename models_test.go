package websitecontacts

import (
	"encoding/json"
	"testing"
)

// TestBool tests JSON encoding/parsing functions for the bool values.
func TestBool(t *testing.T) {
	tests := []struct {
		name   string
		decErr string
		encErr string
	}{
		{
			name:   `true`,
			decErr: "",
			encErr: "",
		},
		{
			name:   `false`,
			decErr: "",
			encErr: "",
		},
		{
			name:   `""`,
			decErr: "json: cannot unmarshal string into Go value of type bool",
			encErr: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var v bool

			err := json.Unmarshal([]byte(tt.name), &v)
			checkErr(t, err, tt.decErr)
			if tt.decErr != "" {
				return
			}

			bb, err := json.Marshal(v)
			checkErr(t, err, tt.encErr)
			if tt.encErr != "" {
				return
			}

			if string(bb) != tt.name {
				t.Errorf("got = %v, want %v", string(bb), tt.name)
			}
		})
	}
}

// checkErr checks for an error.
func checkErr(t *testing.T, err error, want string) {
	if (err != nil || want != "") && (err == nil || err.Error() != want) {
		t.Errorf("error = %v, wantErr %v", err, want)
	}
}
