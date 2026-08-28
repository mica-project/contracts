package equip

import "encoding/json/v2"

type EquipType struct {
	Des  string `json:"des"`
	Name string `json:"name"`
}

type EquipTypeList []EquipType

func (s *EquipType) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
