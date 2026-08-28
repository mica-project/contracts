package battle

import "encoding/json/v2"

type FightType struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type FightTypeList []FightType

func (s *FightType) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
