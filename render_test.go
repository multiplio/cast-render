package main

import (
	"testing"
)

type mockRoutingContext struct{}

func (m mockRoutingContext) Param(key string) string {
	return "hash-from-url"
}
func (m mockRoutingContext) Write(str []byte) (int, error) {
	return 0, nil
}

func TestHandleTwitter(t *testing.T) {
	render := renderContext{shell: nil, renderer: nil}
	routing := mockRoutingContext{}

	err := render.handleTwitter(&routing)
	if err != nil {
		t.Error("handleTwitter returned error :", err)
	}
}
