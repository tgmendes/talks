package auth_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tgmendes/spotify-gate/internal/auth"
)

func TestTokenExpiry(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		desc string
		tok  *auth.Token
		want bool
	}{
		{desc: "12 seconds", tok: &auth.Token{AccessToken: "foo", Expiry: now.Add(12 * time.Second)}, want: true},
		{desc: "10 seconds", tok: &auth.Token{AccessToken: "foo", Expiry: now.Add(10 * time.Second)}, want: true},
		{desc: "10 seconds-1ns", tok: &auth.Token{AccessToken: "foo", Expiry: now.Add(10*time.Second - 1*time.Nanosecond)}, want: false},
		{desc: "-1 hour", tok: &auth.Token{AccessToken: "foo", Expiry: now.Add(-1 * time.Hour)}, want: false},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			assert.Equal(t, tC.want, tC.tok.Valid(now), "validity does not match")
		})
	}
}
