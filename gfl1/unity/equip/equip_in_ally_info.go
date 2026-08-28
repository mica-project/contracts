package equip

import "encoding/json/v2"

type EquipInAllyInfo struct {
	EquipId string `json:"equip_id"`
	Id      string `json:"id"`
}

type EquipInAllyInfoList []EquipInAllyInfo

func (s *EquipInAllyInfo) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
