package gateway

import (
	"strings"
	"testing"
)

// An image is part of what a request costs: the count rules and counting
// clients read takes each one for imageTokens. Its base64 is not its cost,
// or a single screenshot would be hundreds of thousands of tokens (#4239).
func TestEstimateCountsImages(t *testing.T) {
	text := &Request{Messages: []Message{{Role: "user", Parts: []Part{
		{Kind: Text, Text: "abcd"},
	}}}}
	one := &Request{Messages: []Message{{Role: "user", Parts: []Part{
		{Kind: Text, Text: "abcd"},
		{Kind: Image, MediaType: "image/png", Data: "aGVsbG8="},
	}}}}
	big := &Request{Messages: []Message{{Role: "user", Parts: []Part{
		{Kind: Text, Text: "abcd"},
		{Kind: Image, MediaType: "image/png", Data: strings.Repeat("aGVsbG8=", 50000)},
	}}}}

	if got, want := estimate(one)-estimate(text), imageTokens; got != want {
		t.Errorf("an image added %d tokens to estimate(), want imageTokens (%d)", got, want)
	}
	if got, want := estimate(big), estimate(one); got != want {
		t.Errorf("a larger image's base64 changed the count: %d, want %d", got, want)
	}
}
