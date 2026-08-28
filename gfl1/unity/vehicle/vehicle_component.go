package vehicle

import "encoding/json/v2"

type VehicleComponent struct {
	AtkSpeedRatio  int64    `json:"atk_speed_ratio"`
	AttRoll        int64    `json:"att_roll"`
	AttackType     int64    `json:"attack_type"`
	BasicEffect    string   `json:"basic_effect"`
	BonusType      []string `json:"bonus_type"`
	Code           string   `json:"code"`
	CustomType     string   `json:"custom_type"`
	Description    string   `json:"description"`
	EffectRange    string   `json:"effect_range"`
	ExpProvide     int64    `json:"exp_provide"`
	Group          int64    `json:"group"`
	HWeaponRatio   int64    `json:"h_weapon_ratio"`
	HitRatio       int64    `json:"hit_ratio"`
	Id             int64    `json:"id"`
	Introduction   string   `json:"introduction"`
	Name           string   `json:"name"`
	PrecisionRatio int64    `json:"precision_ratio"`
	Rank           int64    `json:"rank"`
	ReloadRatio    int64    `json:"reload_ratio"`
	SearchRange    int64    `json:"search_range"`
	Skill          string   `json:"skill"`
	Type           int64    `json:"type"`
	UnlockLv       []int64  `json:"unlock_lv"`
}

type VehicleComponentList []VehicleComponent

func (s *VehicleComponent) UnmarshalJSON(b []byte) error {
	var tmp map[string]any
	return json.Unmarshal(b, &tmp)
}
