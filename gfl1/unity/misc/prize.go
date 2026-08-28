package misc

import "encoding/json/v2"

type Prize struct {
	Coins   []string `json:"coins"`
	Id      int64    `json:"id"`
	ItemIds string   `json:"item_ids"`
	Name    string   `json:"name"`
}

type PrizeList []Prize

func (s *Prize) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
