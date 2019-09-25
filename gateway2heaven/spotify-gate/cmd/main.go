package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type SpotifyCredentials struct {
	ClientID     string
	ClientSecret string
}

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

func main() {
	tkn, err := generateToken()
	if err != nil {
		log.Fatal(err)
	}
	// getCategories(tkn.AccessToken)
	// searchSpotify(tkn.AccessToken)
	// getArtist(tkn.AccessToken)
	getArtistTopTracks(tkn.AccessToken)
}

func generateToken() (*Token, error) {
	creds := SpotifyCredentials{
		ClientID:     "",
		ClientSecret: "",
	}

	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", creds.ClientID, creds.ClientSecret)))

	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	r, err := http.NewRequest(http.MethodPost, "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	r.Header.Add("Authorization", fmt.Sprintf("Basic %s", auth))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	tkn := Token{}
	if err = json.NewDecoder(resp.Body).Decode(&tkn); err != nil {
		return nil, err
	}

	return &tkn, nil
}

func getCategories(accToken string) {
	r, err := http.NewRequest(http.MethodGet, "https://api.spotify.com/v1/browse/categories", nil)
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accToken))
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(b))
}

func searchSpotify(accToken string) {
	r, err := http.NewRequest(http.MethodGet, "https://api.spotify.com/v1/search", nil)
	q := r.URL.Query()
	q.Add("q", "Led Zeppelin")
	q.Add("type", "artist")
	r.URL.RawQuery = q.Encode()

	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accToken))
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(b))
}

func getArtist(accToken string) {
	r, err := http.NewRequest(http.MethodGet, "https://api.spotify.com/v1/artists/36QJpDe2go2KgaRleHCDTp", nil)

	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accToken))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(b))
}

func getArtistTopTracks(accToken string) {
	r, err := http.NewRequest(http.MethodGet, "https://api.spotify.com/v1/artists/36QJpDe2go2KgaRleHCDTp/top-tracks", nil)
	q := r.URL.Query()
	q.Add("country", "PT")
	r.URL.RawQuery = q.Encode()

	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accToken))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(b))
}
