package misc

import "encoding/json/v2"

type Item struct {
	Arg                string   `json:"arg"`
	Code               string   `json:"code"`
	DetailIntroduction string   `json:"detail_introduction"`
	Id                 int64    `json:"id"`
	Introduction       []string `json:"introduction"`
	ItemAccess         []int64  `json:"item_access"`
	ItemName           string   `json:"item_name"`
	Rank               int64    `json:"rank"`
	Sort               int64    `json:"sort"`
	Type               int64    `json:"type"`
	TypeSort           int64    `json:"type_sort"`
}

type ItemList []Item

func (s *Item) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
