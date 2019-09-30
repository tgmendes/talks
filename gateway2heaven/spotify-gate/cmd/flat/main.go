package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	spotify "github.com/tgmendes/spotify-gate/internal"
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

	mux := http.NewServeMux()
	mux.HandleFunc("/search/", func(w http.ResponseWriter, r *http.Request) {
		a := searchArtist(tkn.AccessToken)

		b, err := json.Marshal(a)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(b)
	})

	mux.HandleFunc("/artist/", func(w http.ResponseWriter, r *http.Request) {
		a := getArtist(tkn.AccessToken)

		b, err := json.Marshal(a)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(b)
	})

	mux.HandleFunc("/artist/albums/", func(w http.ResponseWriter, r *http.Request) {
		a := getArtistAlbums(tkn.AccessToken)
		b, err := json.Marshal(a)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(b)
	})

	mux.HandleFunc("/artist/tracks/top/", func(w http.ResponseWriter, r *http.Request) {
		a := getArtistTopTracks(tkn.AccessToken)

		b, err := json.Marshal(a)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(b)
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func searchArtist(accToken string) *spotify.ArtistResults {
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

	a := spotify.ArtistResults{}

	if err = json.NewDecoder(resp.Body).Decode(&a); err != nil {
		log.Fatal(err)
	}

	return &a
}

func getArtist(accToken string) spotify.Artist {
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

	a := spotify.Artist{}

	if err = json.NewDecoder(resp.Body).Decode(&a); err != nil {
		log.Fatal(err)
	}

	return a
}

func getArtistTopTracks(accToken string) spotify.Tracks {
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

	t := spotify.Tracks{}

	if err = json.NewDecoder(resp.Body).Decode(&t); err != nil {
		log.Fatal(err)
	}

	return t
}

func getArtistAlbums(accToken string) spotify.AlbumResults {
	r, err := http.NewRequest(http.MethodGet, "https://api.spotify.com/v1/artists/36QJpDe2go2KgaRleHCDTp/albums", nil)
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

	a := spotify.AlbumResults{}

	if err = json.NewDecoder(resp.Body).Decode(&a); err != nil {
		log.Fatal(err)
	}

	return a

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
