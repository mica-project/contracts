package vehicle

import "encoding/json/v2"

type VehicleComponentRoll struct {
	HArmorPiercing    int64    `json:"H_armor_piercing"`
	AttrNum           int64    `json:"attr_num"`
	DefBreak          int64    `json:"def_break"`
	Description       string   `json:"description"`
	HeavyDamage       int64    `json:"heavy_damage"`
	Id                int64    `json:"id"`
	Name              string   `json:"name"`
	Precision         int64    `json:"precision"`
	Reload            int64    `json:"reload"`
	SkillGroup1       []string `json:"skill_group1"`
	SkillGroup1Weight int64    `json:"skill_group1_weight"`
}

type VehicleComponentRollList []VehicleComponentRoll

func (s *VehicleComponentRoll) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
