package handler

import (
	"gospsgamemod/Updates"
	"testing"
)

func TestHandleInline(t *testing.T) {
	handleInlineQuery(Updates.InlineQuery{ID: "55778800"})
}
