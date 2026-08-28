package fairy

import "encoding/json/v2"

type FairyTalent struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	TypeId int64  `json:"type_id"`
}

type FairyTalentList []FairyTalent

func (s *FairyTalent) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
