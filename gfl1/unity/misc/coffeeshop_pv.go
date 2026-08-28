package misc

import "encoding/json/v2"

type CoffeeshopPv struct {
	Cost        string `json:"cost"`
	Description string `json:"description"`
	Id          string `json:"id"`
	PrizeId     string `json:"prize_id"`
	Title       string `json:"title"`
}

type CoffeeshopPvList []CoffeeshopPv

func (s *CoffeeshopPv) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
