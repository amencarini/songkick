package songkick

import "strconv"

type Event struct {
	Name           string      `json:"displayName"`
	Type           string      `json:"type"`
	Popularity     float64     `json:"popularity"`
	Status         string      `json:"status"`
	Agerestriction interface{} `json:"ageRestriction"`
	Start          struct {
		Time     string `json:"time"`
		Datetime string `json:"datetime"`
		Date     string `json:"date"`
	} `json:"start"`
	Performances []struct {
		Billingindex int    `json:"billingIndex"`
		Name         string `json:"displayName"`
		Artist       Artist `json:"artist"`
		Billing      string `json:"billing"`
		ID           uint   `json:"id"`
	} `json:"performance"`
	Location struct {
		City string  `json:"city"`
		Lat  float64 `json:"lat"`
		Lng  float64 `json:"lng"`
	} `json:"location"`
	Venue struct {
		Displayname string `json:"displayName"`
		Metroarea   struct {
			Displayname string `json:"displayName"`
			URI         string `json:"uri"`
			ID          int    `json:"id"`
			Country     struct {
				Displayname string `json:"displayName"`
			} `json:"country"`
		} `json:"metroArea"`
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
		URI string  `json:"uri"`
		ID  int     `json:"id"`
	} `json:"venue"`
	URI string `json:"uri"`
	ID  uint   `json:"id"`
}

func (c *Client) GetEvents(minDate, maxDate string, locationId int) ([]Event, error) {
	params := map[string]string{
		"min_date": minDate,
		"max_date": maxDate,
		"location": "sk:" + strconv.Itoa(locationId),
		"page":     "1",
	}

	var events []Event
	page := 1

	for {
		result, err := c.doSearch("events.json", params)
		if err != nil {
			return events, err
		}

		events = append(events, result.Resultspage.Results.Events...)

		if len(events) == result.Resultspage.Totalentries {
			break
		}

		page++
		params["page"] = strconv.Itoa(page)
	}

	return events, nil
}
