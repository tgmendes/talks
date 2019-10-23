package auth

import "time"

const expiryDelta = 10 * time.Second

var timeNow = time.Now

// Token represents a token response from an API.
type Token struct {
	AccessToken string    `json:"access_token"`
	ExpiresIn   int       `json:"expires_in"`
	Scope       string    `json:"scope"`
	Expiry      time.Time `json:"expiry,omitempty"`
}

// Valid reports whether t is non-nil, has an AccessToken, and is not expired.
func (t *Token) Valid(now time.Time) bool {
	return t != nil && t.AccessToken != "" && !t.expired(now)
}

// expired reports whether the token is expired.
// t must be non-nil.
func (t *Token) expired(now time.Time) bool {
	if t.Expiry.IsZero() {
		return false
	}
	return t.Expiry.Round(0).Add(-expiryDelta).Before(now)
}
