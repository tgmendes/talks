package album

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	spotify "github.com/tgmendes/spotify-gate/internal"
)

func Retrieve(ID string) *spotify.Album {

	resp, err := http.Get("http://localhost:8080/album/{" + ID + "}/")
	if err != nil {
		return nil
	}

	meh, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	a := &spotify.Album{}
	err = json.Unmarshal(meh, a)
	if err != nil {
		return nil
	}

	return a
}
