package fairy

import "encoding/json/v2"

type FairyTalentType struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type FairyTalentTypeList []FairyTalentType

func (s *FairyTalentType) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
