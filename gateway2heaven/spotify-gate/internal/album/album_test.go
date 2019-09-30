package album_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	spotify "github.com/tgmendes/spotify-gate/internal"
	"github.com/tgmendes/spotify-gate/internal/album"
)

func Test(t *testing.T) {
	img1 := spotify.Image{"https://i.scdn.co/image/776b35448d2877049145bf9f2a32f7b938c3dd6e", 640, 628}
	img2 := spotify.Image{"https://i.scdn.co/image/82991fe6814b025d4cc0408cf3942a5cbf7795f2", 300, 295}
	img3 := spotify.Image{"https://i.scdn.co/image/beb358c39c411b52964d29cd23f7ea9238638acb", 64, 63}

	imgs := []spotify.Image{img1, img2, img3}

	b := spotify.Album{
		ID:          "6VH2op0GKIl3WNTbZmmcmI",
		Name:        "The Complete BBC Sessions",
		Tracks:      spotify.Tracks{},
		Popularity:  0,
		AlbumType:   "album",
		Copyrights:  nil,
		Genres:      nil,
		Images:      imgs,
		ReleaseDate: "2016-09-16",
	}

	a := album.Retrieve("6VH2op0GKIl3WNTbZmmcmI")

	assert.Equal(t, &a, b)
}
