package vehicle

import "encoding/json/v2"

type VehicleComponentAttValue struct {
	HArmorPiercing int64  `json:"H_armor_piercing"`
	LArmorPiercing int64  `json:"L_armor_piercing"`
	Armor          int64  `json:"armor"`
	AtkSpeed       int64  `json:"atk_speed"`
	CritDamage     int64  `json:"crit_damage"`
	CritRate       int64  `json:"crit_rate"`
	DefBreak       int64  `json:"def_break"`
	Dodge          int64  `json:"dodge"`
	Energy         int64  `json:"energy"`
	EnergyBak      int64  `json:"energy_bak"`
	HeavyDamage    int64  `json:"heavy_damage"`
	Hit            int64  `json:"hit"`
	Id             int64  `json:"id"`
	LightDamage    int64  `json:"light_damage"`
	Name           string `json:"name"`
	Precision      int64  `json:"precision"`
	Reload         int64  `json:"reload"`
	SkillGroup1Lv  int64  `json:"skill_group1_lv"`
	SkillGroup2Lv  int64  `json:"skill_group2_lv"`
	SkillGroup3Lv  int64  `json:"skill_group3_lv"`
}

type VehicleComponentAttValueList []VehicleComponentAttValue

func (s *VehicleComponentAttValue) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
