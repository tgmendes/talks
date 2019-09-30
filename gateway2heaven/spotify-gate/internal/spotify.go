package spotify

type Image struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type Artist struct {
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Popularity   int               `json:"popularity"`
	ExternalURLs map[string]string `json:"external_urls"`
	Genres       []string          `json:"genres"`
	Images       []Image           `json:"images"`
}

type ArtistsPagingObject struct {
	Items []Artist `json:"items"`
	Total int      `json:"total"`
}
type ArtistResults struct {
	Artists ArtistsPagingObject `json:"artists"`
}

type Copyright struct {
	Text string `json:"text"`
	Type string `json:"type"`
}
type Album struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Tracks      Tracks      `json:"tracks"`
	Popularity  int         `json:"popularity"`
	AlbumType   string      `json:"album_type"`
	Copyrights  []Copyright `json:"copyrights"`
	Genres      []string    `json:"genres"`
	Images      []Image     `json:"images"`
	ReleaseDate string      `json:"release_date"`
}

type AlbumResults struct {
	Items []Album `json:"items"`
	Total int     `json:"total"`
}

type Track struct {
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	DurationMs   int               `json:"duration_ms"`
	Explicit     bool              `json:"explicit"`
	ExternalUrls map[string]string `json:"external_urls"`
	TrackNumber  int               `json:"track_number"`
}

type Tracks struct {
	Tracks []Track `json:"tracks"`
}
