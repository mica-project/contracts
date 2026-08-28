package misc

import "encoding/json/v2"

type MedalInfo struct {
	AppearanceId string `json:"appearance_id"`
	EvoStep      string `json:"evo_step"`
	Id           string `json:"id"`
	IfEliteMedal string `json:"if_elite_medal"`
	ItemId       string `json:"item_id"`
	MedalColor   string `json:"medal_color"`
	MedalObtain  string `json:"medal_obtain"`
	Name         string `json:"name"`
	Order        string `json:"order"`
}

type MedalInfoList []MedalInfo

func (s *MedalInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
