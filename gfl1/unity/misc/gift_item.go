package misc

import "encoding/json/v2"

type GiftItem struct {
	BonusDescription  string   `json:"bonus_description"`
	Code              string   `json:"code"`
	Description       []string `json:"description"`
	Favor             int64    `json:"favor"`
	FitGun            int64    `json:"fit_gun"`
	Id                int64    `json:"id"`
	Name              string   `json:"name"`
	NormalDescription string   `json:"normal_description"`
	Poster            int64    `json:"poster"`
	Rank              int64    `json:"rank"`
	Skin              int64    `json:"skin"`
	Sort              int64    `json:"sort"`
	Type              int64    `json:"type"`
	TypeSort          int64    `json:"type_sort"`
}

type GiftItemList []GiftItem

func (s *GiftItem) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
