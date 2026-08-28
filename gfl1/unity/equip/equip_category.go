package equip

import "encoding/json/v2"

type EquipCategory struct {
	Name string `json:"name"`
}

type EquipCategoryList []EquipCategory

func (s *EquipCategory) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
