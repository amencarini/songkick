package songkick

type Artist struct {
	Ontouruntil interface{} `json:"onTourUntil"`
	Identifier  []struct {
		Setlistshref string `json:"setlistsHref"`
		Href         string `json:"href"`
		Mbid         string `json:"mbid"`
		Eventshref   string `json:"eventsHref"`
	} `json:"identifier"`
	URI  string `json:"uri"`
	ID   uint   `json:"id"`
	Name string `json:"displayName"`
}

type Concert struct {
	Name string `json:"displayName"`
}

func (c *Client) SearchArtist(query string) ([]Artist, error) {
	var artists []Artist
	params := map[string]string{
		"query": query,
	}

	result, err := c.doSearch("search/artists.json", params)

	if err != nil {
		return artists, err
	}

	artists = result.Resultspage.Results.Artists
	return artists, nil
}
