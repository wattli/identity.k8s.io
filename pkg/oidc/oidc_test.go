package oidc

import (
	"encoding/json"
	"testing"
)

func TestFormat(t *testing.T) {
	out, err := json.MarshalIndent(Provider("https://accounts.google.com"), "", "  ")
	t.Errorf("err:%v\nout:\n%s", err, out)
}
