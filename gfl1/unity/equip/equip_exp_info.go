package equip

import "encoding/json/v2"

type EquipExpInfo struct {
	Exp   string `json:"exp"`
	Level string `json:"level"`
}

type EquipExpInfoList []EquipExpInfo

func (s *EquipExpInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
