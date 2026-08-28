package equip

import "encoding/json/v2"

type EquipGroup struct {
	Code       string   `json:"code"`
	Des        []string `json:"des"`
	EquipUnit  []int64  `json:"equip_unit"`
	FitGun     string   `json:"fit_gun"`
	GroupSkill []string `json:"group_skill"`
	Id         int64    `json:"id"`
	IsAddition int64    `json:"is_addition"`
	Name       string   `json:"name"`
}

type EquipGroupList []EquipGroup

func (s *EquipGroup) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
