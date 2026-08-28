package misc

import "encoding/json/v2"

type Medal struct {
	AppearanceId string `json:"appearance_id"`
	EvoStep      string `json:"evo_step"`
	Id           int64  `json:"id"`
	ItemId       int64  `json:"item_id"`
	MedalColor   string `json:"medal_color"`
	MedalObtain  string `json:"medal_obtain"`
	Name         string `json:"name"`
	Order        int64  `json:"order"`
}

type MedalList []Medal

func (s *Medal) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
