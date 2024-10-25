package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	tests := map[string]struct {
		input http.Header
		want string
		wantErr bool
	}{
		"no auth handler": { input: http.Header{}, want: "", wantErr: true},
		"split is less than 2": {input: http.Header{"Authorization": {"TESTFAIL"}}, want: "", wantErr: true},
		"split first word is not ApiKey": {input:http.Header{"Authorization": {"TESTFAIL Bearer"}}, want: "", wantErr: true},
		"return ApiKey": {input: http.Header{"Authorization": {"ApiKey TestKeyGotTested"}}, want: "TestKeyGotTested", wantErr: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t * testing.T) {
			got, err := GetAPIKey(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if got != tc.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tc.want)
			}
		})
	}
}