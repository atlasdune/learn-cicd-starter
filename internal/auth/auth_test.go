package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestAuth(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{

		{
			input: "ApiKey YWxhZGRpbjpvcGVuc2VzYW1l", want: "YWxhZGRpbjpvcGVuc2VzYW1l"},
		{input: "", want: ""},
	}
	for _, tc := range tests {
		header := http.Header{}
		header.Set("Authorization", tc.input)
		got, _ := GetAPIKey(header)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("Want | %s, got | %s", tc.want, got)
		}
	}
}
