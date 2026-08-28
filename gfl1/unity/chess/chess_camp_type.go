package chess

import "encoding/json/v2"

type ChessCampType struct {
	CampSkill    string   `json:"camp_skill"`
	Code         string   `json:"code"`
	Color        string   `json:"color"`
	Des          []string `json:"des"`
	Id           int64    `json:"id"`
	Name         string   `json:"name"`
	PassiveSkill []int64  `json:"passive_skill"`
}

type ChessCampTypeList []ChessCampType

func (s *ChessCampType) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
