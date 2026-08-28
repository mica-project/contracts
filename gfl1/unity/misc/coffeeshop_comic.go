package misc

import "encoding/json/v2"

type CoffeeshopComic struct {
	Cost        string `json:"cost"`
	Description string `json:"description"`
	Id          string `json:"id"`
	PrizeId     string `json:"prize_id"`
	Title       string `json:"title"`
}

type CoffeeshopComicList []CoffeeshopComic

func (s *CoffeeshopComic) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
